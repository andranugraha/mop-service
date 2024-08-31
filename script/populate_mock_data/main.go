package main

import (
	"errors"
	"fmt"
	"github.com/empnefsi/mop-service/internal/common/strings"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/empnefsi/mop-service/internal/module/additionalfee"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemcategory"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"github.com/empnefsi/mop-service/internal/module/table"
	"github.com/empnefsi/mop-service/internal/module/user"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Populating mock data...")
	db := getDBInstance()
	if db == nil {
		panic("Database connection is nil")
	}
	fmt.Println("Database connected successfully!")

	fmt.Println("Initializing merchant data...")
	merchantData := merchant.Merchant{
		Name: proto.String("Patua Kopi"),
	}

	fmt.Println("Inserting user data...")
	merchantData.Users = []*user.User{
		{
			Email:    proto.String("admin@gmail.com"),
			Password: proto.String(strings.HashPassword("admin")),
		},
	}

	fmt.Println("Inserting payment type data...")
	merchantData.PaymentTypes = []*paymenttype.PaymentType{
		{
			Name: proto.String("Cash"),
			Type: proto.Uint32(paymenttype.PaymentTypeCashier),
		},
		{
			Name:      proto.String("QR"),
			Type:      proto.Uint32(paymenttype.PaymentTypeQR),
			ExtraData: []byte(`{"image_url":"https://example.com/qr.png"}`),
		},
	}

	fmt.Println("Inserting table data...")
	merchantData.Tables = []*table.Table{
		{
			Code: proto.String("T1"),
		},
		{
			Code: proto.String("T2"),
		},
		{
			Code: proto.String("T3"),
		},
	}

	fmt.Println("Inserting additional fee data...")
	merchantData.AdditionalFees = []*additionalfee.AdditionalFee{
		{
			Name:        proto.String("Service Charge"),
			Description: proto.String("Service charge for every order"),
			Type:        proto.Uint32(additionalfee.AdditionalFeeTypePercentage),
			Fee:         proto.Uint64(10),
		},
		{
			Name:        proto.String("Tax"),
			Description: proto.String("Tax for every order"),
			Type:        proto.Uint32(additionalfee.AdditionalFeeTypePercentage),
			Fee:         proto.Uint64(10),
		},
	}

	fmt.Println("Inserting item category data...")
	merchantData.ItemCategories = []*itemcategory.ItemCategory{
		{
			Name: proto.String("Appetizer"),
			Items: []*item.Item{
				{
					Name:        proto.String("French Fries"),
					Description: proto.String("Crispy and delicious"),
					Price:       proto.Uint64(35000),
				},
				{
					Name:        proto.String("Onion Rings"),
					Description: proto.String("Crispy and delicious"),
					Price:       proto.Uint64(30000),
				},
			},
		},
		{
			Name: proto.String("Main Course"),
			Items: []*item.Item{
				{
					Name:        proto.String("Nasi Goreng"),
					Description: proto.String("Spicy and delicious"),
					Price:       proto.Uint64(45000),
					Variants: []*itemvariant.ItemVariant{
						{
							Name:      proto.String("Spicy Level"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("No Spicy"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Medium"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Spicy"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Extra Spicy"),
									Price: proto.Uint64(2000),
								},
							},
						},
					},
				},
				{
					Name:        proto.String("Mie Goreng"),
					Description: proto.String("Spicy and delicious"),
					Price:       proto.Uint64(45000),
					Variants: []*itemvariant.ItemVariant{
						{
							Name:      proto.String("Spicy Level"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("No Spicy"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Medium"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Spicy"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Extra Spicy"),
									Price: proto.Uint64(2000),
								},
							},
						},
					},
				},
				{
					Name:        proto.String("Seblak"),
					Description: proto.String("Spicy and delicious"),
					Price:       proto.Uint64(40000),
					Variants: []*itemvariant.ItemVariant{
						{
							Name:      proto.String("Add On"),
							MinSelect: proto.Uint32(0),
							MaxSelect: proto.Uint32(4),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("Telur"),
									Price: proto.Uint64(3000),
								},
								{
									Name:  proto.String("Sosis"),
									Price: proto.Uint64(3000),
								},
								{
									Name:  proto.String("Bakso"),
									Price: proto.Uint64(3000),
								},
								{
									Name:  proto.String("Ayam"),
									Price: proto.Uint64(3000),
								},
							},
						},
					},
				},
			},
		},
		{
			Name: proto.String("Drinks"),
			Items: []*item.Item{
				{
					Name:        proto.String("Teh"),
					Description: proto.String("Refreshing and delicious"),
					Price:       proto.Uint64(5000),
					Variants: []*itemvariant.ItemVariant{
						{
							Name:      proto.String("Type"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("Hot"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Iced"),
									Price: proto.Uint64(0),
								},
							},
						},
						{
							Name:      proto.String("Sugar Level"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("No Sugar"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Less Sugar"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Normal"),
									Price: proto.Uint64(0),
								},
							},
						},
					},
				},
				{
					Name:        proto.String("Americano"),
					Description: proto.String("Refreshing and delicious"),
					Price:       proto.Uint64(20000),
					Variants: []*itemvariant.ItemVariant{
						{
							Name:      proto.String("Type"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("Hot"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Iced"),
									Price: proto.Uint64(0),
								},
							},
						},
						{
							Name:      proto.String("Extra Shot"),
							MinSelect: proto.Uint32(0),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("1"),
									Price: proto.Uint64(5000),
								},
								{
									Name:  proto.String("2"),
									Price: proto.Uint64(10000),
								},
								{
									Name:  proto.String("3"),
									Price: proto.Uint64(15000),
								},
							},
						},
						{
							Name:      proto.String("Extra Sugar"),
							MinSelect: proto.Uint32(1),
							MaxSelect: proto.Uint32(1),
							Options: []*itemvariantoption.ItemVariantOption{
								{
									Name:  proto.String("No"),
									Price: proto.Uint64(0),
								},
								{
									Name:  proto.String("Yes"),
									Price: proto.Uint64(0),
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Inserting merchant data...")
	err := db.Create(&merchantData).Error
	if err != nil {
		panic(errors.New("failed to populate mock data: " + err.Error()))
	}
	fmt.Println("Mock data populated successfully!")
	return
}

func getDBInstance() *gorm.DB {
	dbUrl := config.GetDBURL()
	fmt.Println("Connecting to database with URL: " + dbUrl)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(errors.New("failed to connect to database: " + err.Error()))
	}

	return db
}
