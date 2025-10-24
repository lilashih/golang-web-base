package resource

import "gbase/src/model"

var _ IResource = (*MenuResource)(nil)

type Menus struct {
	Menus []model.Menu `json:"menus"`
}

type MenuResource struct {
} //@name MenusResource

// 沒用到
func (r MenuResource) Collection(pagination Pagination, data interface{}) interface{} {
	return data
}

func (r MenuResource) Single(data interface{}) interface{} {
	if items, ok := data.([]model.Menu); ok {
		return Menus{Menus: items}
	}
	return data
}
