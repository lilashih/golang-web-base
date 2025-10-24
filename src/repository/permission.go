package repository

import (
	"gbase/src/http/resource"
	"gbase/src/model"
)

var _ IRepository = (*PermissionRepository)(nil)

type PermissionRepository struct {
	model *model.Permission
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{model: &model.Permission{}}
}

func (r PermissionRepository) NewInput() interface{} {
	return nil
}

func (r PermissionRepository) NewModel() model.IModel {
	return r.model.NewModel()
}

func (r PermissionRepository) FindMenu() ([]model.Menu, error) {
	db := r.model.DB()
	var items []model.Permission

	if err := db.
		Where("isActive = ? AND system = ? AND group1 = ?", 1, "app", "菜單").
		Order("groupOrder").
		Order("`order`").
		Find(&items).Error; err != nil {
		return nil, err
	}

	var groups []model.Menu
	var curGroup *model.Menu // 指向 groups 最後一筆，方便直接操作
	var lastGroup2 string    // 記錄上一筆的 Group2

	for _, item := range items {
		if item.Group2 != lastGroup2 { // 遇到新群組
			groups = append(groups, model.Menu{
				Group2:     item.Group2,
				GroupOrder: item.GroupOrder,
				Children:   []model.Permission{},
			})
			curGroup = &groups[len(groups)-1] // 更新指標避免影響排序
			lastGroup2 = item.Group2
		}
		curGroup.Children = append(curGroup.Children, item)
	}

	return groups, nil
}

func (r PermissionRepository) Find(key any) (interface{}, error) {
	result := r.NewModel()
	err := FindHelper(r.model, key, &result)
	return result, err
}

func (r PermissionRepository) FindAll(search map[string]string) (interface{}, resource.Pagination, error) {
	model := r.NewModel()
	result, pagination, err := FindAllHelper(model, search)
	return result, pagination, err
}

func (r PermissionRepository) Create(data model.IModel) (model.IModel, error) {
	return CreateHelper(data)
}

func (r PermissionRepository) Update(data model.IModel, key any) (model.IModel, error) {
	data.SetKey(key)
	return UpdateHelper(data)
}

func (r PermissionRepository) Delete(key any) error {
	model := r.NewModel()
	return DeleteHelper(model, key)
}

func (r PermissionRepository) UpdateOrder(id1, id2 any) error {
	return nil
}
