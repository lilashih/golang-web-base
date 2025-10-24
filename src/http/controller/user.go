package controller

import (
	"gbase/src/http/resource"
	"gbase/src/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

func NewUserController() *UserController {
	repo := repository.NewUserRepository()
	res := resource.UserResource{}
	base := NewBaseController(repo, res)

	return &UserController{Controller: base}
}

func (ctrl *UserController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/users", ctrl.UserIndex)
	r.GET("/users/:id", ctrl.UserShow)
	r.POST("/users", ctrl.UserStore)
	r.PUT("/users/:id", ctrl.UserUpdate)
	r.DELETE("/users/:id", ctrl.UserDestroy)
	r.POST("/users/order", ctrl.UserUpdateOrder)
}

// @Tags User 使用者
// @Router /users [get]
// @Summary 獲取使用者列表
// @Accept  json
// @Produce  json
// @Param   page query int false "資料頁數" default(1)
// @Param   perPage query int false "資料限制，最大值為100" default(100)
// @Param   name query string false "搜尋名稱"
// @Success 200 {object} response.Response{data=resource.Users} "code:200"
func (ctrl *UserController) UserIndex(c *gin.Context) {
	ctrl.Index(c)
}

// @Tags User 使用者
// @Router /users/{id} [get]
// @Summary 獲取使用者資訊
// @Accept  json
// @Produce  json
// @Param   id path int true "Id"
// @Success 200 {object} response.Response{data=resource.User} "code:200"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
func (ctrl *UserController) UserShow(c *gin.Context) {
	ctrl.Show(c)
}

// @Tags User 使用者
// @Router /users [post]
// @Summary 新增使用者
// @Accept  json
// @Produce  json
// @Param  data body model.UserInput true "使用者資訊"
// @Success 200 {object} response.Response{data=resource.User} "code:200, message:新增成功"
// @Failure 400 {object} response.Response{} "code:400, message:新增失敗"
// @Failure 422 {object} response.Response{data=resource.ErrorValidation} "code:422 (表單資料有誤)"
func (ctrl *UserController) UserStore(c *gin.Context) {
	ctrl.Store(c)
}

// @Tags User 使用者
// @Router /users/{id} [put]
// @Summary 編輯使用者
// @Accept  json
// @Produce  json
// @Param   id path int true "Id"
// @Param  data body model.UserInput true "使用者資訊"
// @Success 200 {object} response.Response{data=resource.User} "code:200, message:編輯成功"
// @Failure 400 {object} response.Response{} "code:400, message:編輯失敗"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
// @Failure 422 {object} response.Response{data=resource.ErrorValidation} "code:422 (表單資料有誤)"
func (ctrl *UserController) UserUpdate(c *gin.Context) {
	ctrl.Update(c)
}

// @Tags User 使用者
// @Router /users/{id} [delete]
// @Summary 刪除使用者
// @Accept  json
// @Produce  json
// @Param   id path int true "Id"
// @Success 200 {object} response.Response{} "code:200, message:刪除成功"
// @Failure 400 {object} response.Response{} "code:400, message:刪除失敗"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
func (ctrl *UserController) UserDestroy(c *gin.Context) {
	ctrl.Destroy(c)
}

// @Tags User 使用者
// @Router /users/order [post]
// @Summary 修改使用者排序
// @Accept  json
// @Produce  json
// @Param  data body  model.OrderInput true "要交換順序的兩筆資料id"
// @Success 200 {object} response.Response{} "code:200, message:編輯成功"
// @Failure 400 {object} response.Response{} "code:400, message:編輯失敗"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
func (ctrl *UserController) UserUpdateOrder(c *gin.Context) {
	ctrl.UpdateOrder(c)
}
