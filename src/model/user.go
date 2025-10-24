package model

var _ IModel = (*User)(nil)

type UserInput struct {
	Name string `json:"name" validate:"required,max=300" example:"Tom"` // User name
} //@name UserInput

type User struct {
	SoftDeleteModel
	UserInput
} //@name User

func (User) NewModel() IModel {
	return &User{}
}

func (User) NewModels() interface{} {
	return []User{}
}
