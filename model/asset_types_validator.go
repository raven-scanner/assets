package model

import (
    "context"
 
    "github.com/rs/rest-layer/resource"
)

type AssetValidator struct {
    AssetTypeValidators map[string]AssetTypeValidator
}
 
func (av AssetValidator) OnInsert(ctx context.Context, items []*resource.Item) error {
    for _, i := range items {
        asset_value, err := validate_type(i.Payload["type"].(string), i.Payload["value"].(string))
        if err != nil {
            return err
        }
        i.Payload["value"] = asset_value
    }
    return nil
}
 
func (agv AssetValidator) OnUpdate(ctx context.Context, item *resource.Item, original *resource.Item) error {
    asset_value, err := validate_type(item.Payload["type"].(string), item.Payload["value"].(string))
    if err != nil {
        return err
    }
    item.Payload["value"] = asset_value
    return nil
}