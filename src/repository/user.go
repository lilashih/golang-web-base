package repository

import (
	"gbase/src/http/resource"
	"gbase/src/model"
)

var _ IRepository = (*UserRepository)(nil)

type UserRepository struct {
	model *model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{model: &model.User{}}
}

func (r UserRepository) NewInput() interface{} {
	return &model.UserInput{}
}

func (r UserRepository) NewModel() model.IModel {
	return r.model.NewModel()
}

func (r UserRepository) Find(key any) (interface{}, error) {
	result := r.NewModel()
	err := FindHelper(r.model, key, &result)
	return result, err
}

func (r UserRepository) FindAll(search map[string]string) (interface{}, resource.Pagination, error) {
	return FindAllHelper(r.NewModel(), search)
}

func (r UserRepository) Create(data model.IModel) (model.IModel, error) {
	return CreateHelper(data)
}

func (r UserRepository) Update(data model.IModel, key any) (model.IModel, error) {
	data.SetKey(key)
	return UpdateHelper(data)
}

func (r UserRepository) Delete(key any) error {
	return DeleteHelper(r.NewModel(), key)
}

func (r UserRepository) UpdateOrder(id1, id2 any) error {
	return UpdateOrderHelper(r.NewModel(), id1, id2)
}
