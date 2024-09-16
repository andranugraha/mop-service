package order

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/empnefsi/mop-service/internal/module/additionalfee"
	"github.com/empnefsi/mop-service/internal/module/invoice"
	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/order"
	"github.com/empnefsi/mop-service/internal/module/paymentgateway"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"github.com/empnefsi/mop-service/internal/module/tableorder"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
	dto "github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"github.com/empnefsi/mop-service/internal/module/orderitem"
	"google.golang.org/protobuf/proto"
)

func (m *impl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	// Validate payment method
	if req.PaymentMethod != paymenttype.PaymentTypeQR {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderPaymentMethodNotSupported.Error())
		return nil, constant.ErrOrderPaymentMethodNotSupported
	}

	// Fetch merchant and table info
	merchant, err := m.merchantModule.GetMerchantByID(ctx, req.MerchantID)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to get merchant: %v", err)
		return nil, err
	}

	table, err := m.tableModule.GetTableByID(ctx, req.TableID)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to get table: %v", err)
		return nil, err
	}

	// Validate table and merchant association
	if table.GetMerchantId() != req.MerchantID {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderTableAndMerchantMismatch.Error())
		return nil, constant.ErrOrderTableAndMerchantMismatch
	}

	// Extract item, variant, and option IDs
	itemIds, variantIds, variantOptionIds := extractIds(req.Items)

	// Fetch data for items, variants, and options concurrently
	itemMap, itemVariantMap, variantOptionMap, err := m.fetchData(ctx, itemIds, variantIds, variantOptionIds)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to fetch item data: %v", err)
		return nil, err
	}

	// Validate fetched data
	if len(itemMap) != len(itemIds) || len(itemVariantMap) != len(variantIds) || len(variantOptionMap) != len(variantOptionIds) {
		logger.Error(ctx, "CreateOrder", constant.ErrItemNotFound.Error())
		return nil, constant.ErrItemNotFound
	}

	// Construct order items and calculate base price
	orderItems, basePrice, err := m.constructOrderItems(req.Items, itemMap, itemVariantMap, variantOptionMap)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to create order items: %v", err)
		return nil, err
	}

	// Calculate additional fees and grand total
	extraCharge, additionalFees, err := m.calculateAdditionalFees(merchant.GetAdditionalFees(), basePrice)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to calculate additional fees: %v", err)
		return nil, err
	}

	grandTotal := basePrice + extraCharge
	if grandTotal != req.TotalPrice {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderTotalPriceMismatch.Error())
		return nil, constant.ErrOrderTotalPriceMismatch
	}

	// Get payment ID
	paymentID, err := m.getPaymentID(merchant.GetPaymentTypes(), req.PaymentMethod)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to get payment ID: %v", err)
		return nil, err
	}

	// Marshal additional fees into JSON
	additionalFeeJSON, err := json.Marshal(additionalFees)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to marshal additional fees: %v", err)
		return nil, constant.ErrInternalServer
	}

	// Generate invoice code and create invoice data
	todayLatestInvoice, _ := m.invoiceModule.GetTodayLatestInvoice(ctx, req.MerchantID)
	invoiceData := m.buildInvoiceData(merchant, todayLatestInvoice, paymentID, grandTotal, additionalFeeJSON)

	// Create order data
	orderData := m.buildOrderData(req, basePrice, orderItems, invoiceData)
	logger.Data(ctx, "CreateOrder", "order", orderData)

	// Charge payment via gateway
	chargePayment, err := m.chargePayment(ctx, invoiceData.GetCode(), grandTotal)
	if err != nil {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderGeneratePayment.Error())
		return nil, constant.ErrOrderGeneratePayment
	}

	// Create order in the system
	err = m.orderModule.CreateOrder(ctx, orderData)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to create order: %v", err)
		return nil, constant.ErrInternalServer
	}

	// Return response with payment info
	return &dto.CreateOrderResponse{
		OrderID:   orderData.GetId(),
		OrderCode: invoiceData.GetCode(),
		Total:     grandTotal,
		PaymentQR: chargePayment.Actions[0].URL,
		DueTime:   uint64(time.Now().Add(15 * time.Minute).Unix()),
	}, nil
}

func extractIds(items []dto.Item) (itemIds, variantIds, variantOptionIds []uint64) {
	itemIdsMap := make(map[uint64]struct{})
	variantIdsMap := make(map[uint64]struct{})
	variantOptionIdsMap := make(map[uint64]struct{})

	for _, reqItem := range items {
		// Collect unique item IDs
		if _, exists := itemIdsMap[reqItem.ItemID]; !exists {
			itemIdsMap[reqItem.ItemID] = struct{}{}
			itemIds = append(itemIds, reqItem.ItemID)
		}

		// Collect unique variant IDs and variant option IDs
		for _, variant := range reqItem.Variants {
			if _, exists := variantIdsMap[variant.VariantID]; !exists {
				variantIdsMap[variant.VariantID] = struct{}{}
				variantIds = append(variantIds, variant.VariantID)
			}

			for _, optionID := range variant.OptionIDs {
				if _, exists := variantOptionIdsMap[optionID]; !exists {
					variantOptionIdsMap[optionID] = struct{}{}
					variantOptionIds = append(variantOptionIds, optionID)
				}
			}
		}
	}
	return itemIds, variantIds, variantOptionIds
}

// calculateAdditionalFees calculates extra charges based on the merchant's fee policies
func (m *impl) calculateAdditionalFees(fees []*additionalfee.AdditionalFee, basePrice uint64) (extraCharge uint64, additionalFees []*invoice.AdditionalFee, err error) {
	for _, fee := range fees {
		data := &invoice.AdditionalFee{
			Id:   fee.Id,
			Type: fee.Type,
			Name: fee.Name,
			Fee:  fee.Fee,
		}
		switch fee.GetType() {
		case additionalfee.AdditionalFeeTypeFixed:
			data.Amount = proto.Uint64(fee.GetFee())
			extraCharge += fee.GetFee()
		case additionalfee.AdditionalFeeTypePercentage:
			charge := basePrice * fee.GetFee() / 100
			data.Amount = proto.Uint64(charge)
			extraCharge += charge
		}
		additionalFees = append(additionalFees, data)
	}
	return extraCharge, additionalFees, nil
}

// getPaymentID retrieves the payment ID for the specified payment method
func (m *impl) getPaymentID(paymentTypes []*paymenttype.PaymentType, paymentMethod uint32) (uint64, error) {
	for _, paymentType := range paymentTypes {
		if paymentType.GetType() == paymentMethod {
			return paymentType.GetId(), nil
		}
	}
	return 0, constant.ErrOrderPaymentMethodNotSupported
}

// buildInvoiceData creates and returns invoice data
func (m *impl) buildInvoiceData(merchant *merchant.Merchant, latestInvoice *invoice.Invoice, paymentID, grandTotal uint64, additionalFees []byte) *invoice.Invoice {
	return &invoice.Invoice{
		Code:           proto.String(invoice.GenerateInvoiceCode(merchant.GetCode(), latestInvoice)),
		MerchantId:     proto.Uint64(merchant.GetId()),
		PaymentTypeId:  proto.Uint64(paymentID),
		TotalPayment:   proto.Uint64(grandTotal),
		AdditionalFees: additionalFees,
		Status:         proto.Uint32(invoice.StatusPending),
	}
}

// buildOrderData creates and returns order data
func (m *impl) buildOrderData(req *dto.CreateOrderRequest, basePrice uint64, orderItems []*orderitem.OrderItem, invoiceData *invoice.Invoice) *order.Order {
	return &order.Order{
		MerchantId: proto.Uint64(req.MerchantID),
		TotalSpend: proto.Uint64(basePrice),
		Status:     proto.Uint32(order.StatusPending),
		OrderItems: orderItems,
		Tables:     []*tableorder.TableOrder{{TableId: proto.Uint64(req.TableID)}},
		Invoice:    invoiceData,
	}
}

// chargePayment handles payment charging via payment gateway
func (m *impl) chargePayment(ctx context.Context, orderID string, grandTotal uint64) (*paymentgateway.PaymentResponse, error) {
	now := time.Now()
	return paymentgateway.GetModule().ChargePayment(ctx, &paymentgateway.PaymentRequest{
		PaymentType: paymentgateway.PaymentTypeQRIS,
		TransactionDetails: paymentgateway.TransactionDetails{
			OrderID:     orderID,
			GrossAmount: int(grandTotal),
		},
		CustomExpiry: &paymentgateway.CustomExpiry{
			ExpiryDuration: 15,
			Unit:           paymentgateway.UnitMinute,
			OrderTime:      now.Format("2006-01-02 15:04:05 -0700"),
		},
	})
}

func (m *impl) fetchData(ctx context.Context, itemIds, variantIds, variantOptionIds []uint64) (map[uint64]*item.Item, map[uint64][]*itemvariant.ItemVariant, map[uint64][]*itemvariantoption.ItemVariantOption, error) {
	var (
		wg               sync.WaitGroup
		wgErr            error
		itemMap          = make(map[uint64]*item.Item, len(itemIds))
		itemVariantMap   = make(map[uint64][]*itemvariant.ItemVariant, len(variantIds))
		variantOptionMap = make(map[uint64][]*itemvariantoption.ItemVariantOption, len(variantOptionIds))
	)
	wg.Add(3)

	go func() {
		defer wg.Done()
		items, err := m.itemModule.GetActiveItemsByIDs(ctx, itemIds)
		if err != nil {
			logger.Error(ctx, "CreateOrder", "failed to get items: %v", err.Error())
			wgErr = err
			return
		}
		for _, item := range items {
			itemMap[item.GetId()] = item
		}
	}()

	go func() {
		defer wg.Done()
		variants, err := m.itemVariantModule.GetActiveItemVariantsByIDs(ctx, variantIds)
		if err != nil {
			logger.Error(ctx, "CreateOrder", "failed to get item variants: %v", err.Error())
			wgErr = err
			return
		}
		for _, variant := range variants {
			itemVariantMap[variant.GetItemId()] = append(itemVariantMap[variant.GetItemId()], variant)
		}
	}()

	go func() {
		defer wg.Done()
		options, err := m.itemVariantOptionModule.GetActiveItemVariantOptionsByIDs(ctx, variantOptionIds)
		if err != nil {
			logger.Error(ctx, "CreateOrder", "failed to get item variant options: %v", err.Error())
			wgErr = err
			return
		}
		for _, option := range options {
			variantOptionMap[option.GetItemVariantId()] = append(variantOptionMap[option.GetItemVariantId()], option)
		}
	}()

	wg.Wait()
	if wgErr != nil {
		return nil, nil, nil, wgErr
	}
	return itemMap, itemVariantMap, variantOptionMap, nil
}

func (m *impl) constructOrderItems(reqItems []dto.Item, itemMap map[uint64]*item.Item, itemVariantMap map[uint64][]*itemvariant.ItemVariant, variantOptionMap map[uint64][]*itemvariantoption.ItemVariantOption) ([]*orderitem.OrderItem, uint64, error) {
	var (
		orderItems = make([]*orderitem.OrderItem, 0, len(reqItems))
		basePrice  uint64
	)
	for _, reqItem := range reqItems {
		item := itemMap[reqItem.ItemID]
		itemPrice := item.GetPrice()
		orderItem := &orderitem.OrderItem{
			ItemId:       item.Id,
			ItemName:     item.Name,
			PricePerItem: item.Price,
			Amount:       proto.Uint64(reqItem.Amount),
			Note:         reqItem.Note,
		}
		itemOptions := make([]*orderitem.Variant, 0, len(reqItem.Variants))
		for _, variant := range itemVariantMap[item.GetId()] {
			variantOptions := make([]orderitem.Options, 0, len(variantOptionMap[variant.GetId()]))
			for _, option := range variantOptionMap[variant.GetId()] {
				variantOptions = append(variantOptions, orderitem.Options{
					Id:         option.Id,
					OptionName: option.Name,
					Price:      option.Price,
				})
				itemPrice += option.GetPrice()
			}
			itemOption := &orderitem.Variant{
				Id:          variant.Id,
				VariantName: variant.Name,
				Options:     variantOptions,
			}
			itemOptions = append(itemOptions, itemOption)
		}

		jsonItemOptions, err := json.Marshal(itemOptions)
		if err != nil {
			return nil, 0, err
		}
		orderItem.ItemOptions = jsonItemOptions
		orderItem.TotalPrice = proto.Uint64(itemPrice * reqItem.Amount)

		orderItems = append(orderItems, orderItem)
		basePrice += orderItem.GetTotalPrice()
	}
	return orderItems, basePrice, nil
}
