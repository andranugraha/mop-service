package order

import (
	"context"
	"encoding/json"
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
	dto "github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/empnefsi/mop-service/internal/module/additionalfee"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"github.com/empnefsi/mop-service/internal/module/order"
	"github.com/empnefsi/mop-service/internal/module/orderitem"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"github.com/empnefsi/mop-service/internal/module/tableorder"
	"google.golang.org/protobuf/proto"
	"sync"
)

func (m *impl) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	merchant, err := m.merchantModule.GetMerchantByID(ctx, req.MerchantID)
	if err != nil {
		return nil, err
	}

	table, err := m.tableModule.GetTableByID(ctx, req.TableID)
	if err != nil {
		return nil, err
	}

	if table.GetMerchantId() != req.MerchantID {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderTableAndMerchantMismatch.Error())
		return nil, constant.ErrOrderTableAndMerchantMismatch
	}

	itemIds, variantIds, variantOptionIds := extractIds(req.Items)

	itemMap, itemVariantMap, variantOptionMap, err := m.fetchData(ctx, itemIds, variantIds, variantOptionIds)
	if err != nil {
		return nil, err
	}

	if len(itemMap) != len(itemIds) || len(itemVariantMap) != len(variantIds) || len(variantOptionMap) != len(variantOptionIds) {
		logger.Error(ctx, "CreateOrder", constant.ErrItemNotFound.Error())
		return nil, constant.ErrItemNotFound
	}

	orderItems, basePrice, err := m.constructOrderItems(req.Items, itemMap, itemVariantMap, variantOptionMap)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to create order items: %v", err.Error())
		return nil, constant.ErrInternalServer
	}

	var extraCharge uint64
	for _, additionalFee := range merchant.GetAdditionalFees() {
		switch additionalFee.GetType() {
		case additionalfee.AdditionalFeeTypeFixed:
			extraCharge += additionalFee.GetFee()
		case additionalfee.AdditionalFeeTypePercentage:
			extraCharge += basePrice * additionalFee.GetFee() / 100
		}
	}

	grandTotal := basePrice + extraCharge
	if grandTotal != req.TotalPrice {
		logger.Error(ctx, "CreateOrder", constant.ErrOrderTotalPriceMismatch.Error())
		return nil, constant.ErrOrderTotalPriceMismatch
	}

	orderData := &order.Order{
		Code:       proto.String(merchant.GetCode()),
		MerchantId: proto.Uint64(req.MerchantID),
		TotalSpend: proto.Uint64(basePrice),
		Status:     proto.Uint32(order.StatusPending),
		OrderItems: orderItems,
		Tables:     []*tableorder.TableOrder{{TableId: proto.Uint64(req.TableID)}},
	}
	logger.Data(ctx, "CreateOrder", "order", orderData)
	err = m.orderModule.CreateOrder(ctx, orderData)
	if err != nil {
		logger.Error(ctx, "CreateOrder", "failed to create order: %v", err.Error())
		return nil, constant.ErrInternalServer
	}

	var paymentQR string
	for _, paymentType := range merchant.GetPaymentTypes() {
		if paymentType.GetType() == paymenttype.PaymentTypeQR {
			paymentQR = paymentType.GetQRPaymentTypeExtraData().ImageURL
			break
		}
	}

	dueTime := orderData.GetCtime() + 300 // 5 minutes

	return &dto.CreateOrderResponse{
		OrderID:   orderData.GetId(),
		OrderCode: orderData.GetCode(),
		Total:     grandTotal,
		PaymentQR: paymentQR,
		DueTime:   dueTime,
	}, nil
}

func extractIds(items []dto.Item) (itemIds, variantIds, variantOptionIds []uint64) {
	itemIds = make([]uint64, 0, len(items))
	for _, reqItem := range items {
		itemIds = append(itemIds, reqItem.ItemID)
		for _, variant := range reqItem.Variants {
			variantIds = append(variantIds, variant.VariantID)
			variantOptionIds = append(variantOptionIds, variant.OptionIDs...)
		}
	}
	return
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
