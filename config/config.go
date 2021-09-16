package config

import (
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
)

var Module *sonic.Module

func init() {

	categoryKey := "LOOKUP_category"
	colorKey := "color"
	sizeKey := "size"

	Module = &sonic.Module{
		Name:     string(product.SiteSupreme),
		Accounts: false,
		Fields: []*sonic.ModuleField{
			{
				Type:           sonic.FieldTypeDropDown,
				Label:          "Category",
				FieldKey:       &categoryKey,
				DropdownValues: []string{"Shirts", "Pants", "Accessories", "Tops/Sweaters", "Sweatshirts", "Hats", "Jackets", "Skate", "Bags", "Shoes", "New"},
			},
			{
				Type:  sonic.FieldTypeKeywords,
				Label: "Keywords",
			},
			{
				Type:     sonic.FieldTypeText,
				Label:    "Color",
				FieldKey: &colorKey,
			},
			{
				Type:     sonic.FieldTypeDropDown,
				Label:    "Size",
				FieldKey: &sizeKey,
				DropdownValues: []string{
					"Small",
					"Medium",
					"Large",
					"XLarge",
					"XXL",
					"30",
					"32",
					"34",
					"36",
					"N/A",
					"129",
					"139",
					"149",
					"US 6 / UK 5",
					"US 6.5 / UK 5.5",
					"US 7 / UK 6",
					"US 7.5 / UK 6.5",
					"US 8 / UK 7",
					"US 8.5 / UK 7.5",
					"US 9 / UK 8",
					"US 9.5 / UK 8.5",
					"US 10 / UK 9",
					"US 10.5 / UK 9.5",
					"US 11 / UK 10",
					"US 11.5 / UK 10.5",
					"US 12 / UK 11",
					"US 13 / UK 12",
				},
			},
		},
	}
}
