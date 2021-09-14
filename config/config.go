package config

import (
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"os"
	"strings"
)

var Module *sonic.Module

func init() {
	var name = "Supreme"

	if podName := os.Getenv("POD_NAME"); podName != "" {
		name = strings.Split(podName, "-")[0]
	}

	categoryKey := "LOOKUP_category"

	Module = &sonic.Module{
		Name:     name,
		Accounts: false,
		Fields: []*sonic.ModuleField{
			{
				Type:  sonic.FieldTypeKeywords,
				Label: "Keywords",
			},
			{
				Type:     sonic.FieldTypeDropDown,
				Label:    "Category",
				FieldKey: &categoryKey,
				DropdownValues: []string{"Shirts", "Pants", "Accessories", "Tops/Sweaters", "Sweatshirts", "Hats", "Jackets", "Skate", "Bags", "Shoes", "New",},
			},
		},
	}
}
