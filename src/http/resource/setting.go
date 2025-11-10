package resource

import "gbase/src/model"

var _ IResource = (*SettingResource)(nil)

type Settings struct {
	Settings []model.Setting `json:"settings"`
} //@name SettingsResource

type Setting struct {
	Setting model.Setting `json:"setting"`
} //@name SettingResource

type SettingResource struct {
}

func (r SettingResource) Collection(pagination Pagination, data interface{}) interface{} {
	if items, ok := data.([]model.Setting); ok {
		return Settings{Settings: items}
	}
	return data
}

func (r SettingResource) Single(data interface{}) interface{} {
	switch v := data.(type) {
	case model.Setting:
		return Setting{Setting: v}
	case *model.Setting:
		return Setting{Setting: *v}
	default:
		return data
	}
}
