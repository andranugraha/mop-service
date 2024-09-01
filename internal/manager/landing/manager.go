package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemcategory"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"github.com/empnefsi/mop-service/internal/module/merchant"
)

type Manager interface {
	Landing(ctx context.Context, code string) (*landing.LandingResponseData, error)
}

type impl struct {
	merchantModule          merchant.Module
	itemcategoryModule      itemcategory.Module
	itemModule              item.Module
	itemvariantModule       itemvariant.Module
	itemvariantoptionModule itemvariantoption.Module
}

func NewManager() Manager {
	return &impl{
		merchantModule:     merchant.GetModule(),
		itemcategoryModule: itemcategory.GetModule(),
		itemModule:         item.GetModule(),
	}
}
