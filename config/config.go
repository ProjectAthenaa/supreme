package config

import (
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
)

var Module *sonic.Module

func init() {

	categoryKey := "LOOKUP_category"

	Module = &sonic.Module{
		Name:     string(product.SiteSupreme),
		Accounts: false,
		Fields: []*sonic.ModuleField{
			{
				Type:  sonic.FieldTypeKeywords,
				Label: "Keywords",
			},
			{
				Type:           sonic.FieldTypeDropDown,
				Label:          "Category",
				FieldKey:       &categoryKey,
				DropdownValues: []string{"Shirts", "Pants", "Accessories", "Tops/Sweaters", "Sweatshirts", "Hats", "Jackets", "Skate", "Bags", "Shoes", "New"},
			},
		},
	}
}
