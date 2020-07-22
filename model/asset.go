package model

import (
	"github.com/rs/rest-layer/schema"
	"github.com/rs/rest-layer/resource"

	utils "github.com/circuit-platform/models-utils"

	"github.com/apuigsech/rest-layer-ttl"
	"github.com/apuigsech/rest-layer-sql"
)

var (
	AssetTypesList = func() []string {
    	var types []string 
    	for k := range AssetTypeValidators {
        	types = append(types, k)
        }
   		return types
	}

	AssetTypeField = schema.Field{
		Required: true,
		Filterable: true,
		Validator: &schema.String{
			MaxLen: 64,
			Allowed: AssetTypesList(),
		},
	}

	Asset = schema.Schema{
		Fields: schema.Fields{
			"id": schema.IDField,
			"created": utils.CreatedField,
			"updated": utils.UpdatedField,
			"namespace": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					Regexp: "^[0-9a-v]{20}$",
				},
			},
			"type": AssetTypeField,
			"value": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 512,
				},
			},
			"env_cvss": schema.Field{
				Required: false,
				Default:  "CVSS:3.0/CR:X/IR:X/AR:X/MAV:X/MAC:X/MPR:X/MUI:X/MS:X/MC:X/MI:X/MA:X",
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 128,
					Regexp: "^CVSS:3.0/CR:([XHML])/IR:([XHML])/AR:([XHML])/MAV:([XNALP])/MAC:([XLH])/MPR:([XNLH])/MUI:([XNR])/MS:([XUC])/MC:([XHLN])/MI:([XHLN])/MA:([XHLN])$",
				},
			},
            "ttl": ttl.TTLField,
            "deleteat": ttl.DeleteAtField,
            "active": ttl.ActiveField,
		},
	}

	AssetMetadata = schema.Schema{
		Fields: schema.Fields{
			"id": schema.IDField,
			"created": utils.CreatedField,
			"updated": utils.UpdatedField,
			"asset": {
				Required:   true,
				Filterable: true,
				ReadOnly:   true,
				Validator: &schema.Reference{
					Path: "assets",
				},
			},
			"key": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
			"value": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}

	AssetQueryTemplates = map[string]string{
		"insert": "%s ON CONFLICT ON CONSTRAINT asset_in_namespace DO UPDATE SET updated = EXCLUDED.updated, env_cvss = EXCLUDED.env_cvss, ttl = EXCLUDED.ttl, deleteat = EXCLUDED.deleteat, active = EXCLUDED.active",
	}
)


func CreateAssetsIndex(databaseSource string, databaseSchema string) resource.Index {
	index := resource.NewIndex()

	if databaseSchema != "" {
		databaseSchema = databaseSchema + "."
	}

	assets := index.Bind("assets", Asset, utils.CreateStorerOrDie(databaseSource, databaseSchema + "assets",
		&sqlStorage.Config{
			//QueryTemplates: AssetQueryTemplates,
			VerboseLevel: sqlStorage.DEBUG,
		},
	), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	assets.Bind("metadata", "asset", AssetMetadata, utils.CreateStorerOrDie(databaseSource, databaseSchema + "assets_metadata",
		&sqlStorage.Config{},
	), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	return index
}