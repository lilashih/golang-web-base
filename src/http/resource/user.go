package resource

import "gbase/src/model"

var _ IResource = (*UserResource)(nil)

type Users struct {
	Pagination Pagination   `json:"pagination"`
	Users      []model.User `json:"users"`
} //@name UsersResource

type User struct {
	User model.User `json:"user"`
} //@name UserResource

type UserResource struct {
}

func (r UserResource) Collection(pagination Pagination, data interface{}) interface{} {
	if items, ok := data.([]model.User); ok {
		return Users{Pagination: pagination, Users: items}
	}
	return data
}

func (r UserResource) Single(data interface{}) interface{} {
	switch v := data.(type) {
	case model.User:
		return User{User: v}
	case *model.User:
		return User{User: *v}
	default:
		return data
	}
}
