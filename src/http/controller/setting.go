package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"gbase/src/core/http/request"
	"gbase/src/core/http/response"
	"gbase/src/def"
	"gbase/src/http/resource"
	m "gbase/src/model"
	"gbase/src/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SettingController struct {
	repo *repository.SettingRepository
	res  resource.SettingResource
}

func NewSettingController() *SettingController {
	repo := repository.NewSettingRepository()
	res := resource.SettingResource{}

	return &SettingController{repo: repo, res: res}
}

func (ctrl *SettingController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/settings", ctrl.SettingIndex)
	r.GET("/settings/:id", ctrl.SettingShow)

	r.GET("/settings/groups/:group", ctrl.SettingGroupShow)
	r.PUT("/settings/groups/:group", ctrl.SettingUpdate)
}

// @Tags Setting 設定
// @Router /settings [get]
// @Summary 獲取設定
// @Description 獲取系統設定，依權限回傳可存取的設定
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=resource.Settings} "code:200"
func (ctrl *SettingController) SettingIndex(c *gin.Context) {
	list, pagination, err := ctrl.repo.FindAll(nil)
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Collection(pagination, list))
}

// @Tags Setting 設定
// @Router /settings/{id} [get]
// @Summary 獲取設定資訊
// @Accept  json
// @Produce  json
// @Param   id path string true "設定ID"
// @Success 200 {object} response.Response{data=resource.Settings} "code:200"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
func (ctrl *SettingController) SettingShow(c *gin.Context) {
	id := c.Param("id")

	result, err := ctrl.repo.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorNotFound(c)
			return
		}
		response.Error(c, nil, fmt.Sprintf("查詢失敗: %v", err))
		return
	}

	response.Success(c, ctrl.res.Single(result))
}

// @Tags Setting 設定
// @Router /settings/groups/{group} [get]
// @Summary 獲取設定資訊
// @Accept  json
// @Produce  json
// @Param   group path string true "群組"
// @Success 200 {object} response.Response{data=resource.Settings} "code:200"
func (ctrl *SettingController) SettingGroupShow(c *gin.Context) {
	group := c.Param("group")

	list, pagination, err := ctrl.repo.FindAll(map[string]string{"group": group})
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Collection(pagination, list))
}

// @Tags Setting 設定
// @Router /settings/groups/{group} [put]
// @Summary 編輯設定
// @Accept  json
// @Produce  json
// @Param   group path string true "群組"
// @Param  data body model.SettingInputs true "設定資訊"
// @Success 200 {object} response.Response{data=resource.Setting} "code:200, message:編輯成功"
// @Failure 400 {object} response.Response{} "code:400, message:編輯失敗"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
// @Failure 422 {object} response.Response{data=resource.ErrorValidation} "code:422 (表單資料有誤)"
func (ctrl *SettingController) SettingUpdate(c *gin.Context) {
	group := c.Param("group")

	// get body data
	body := m.SettingInputs{}
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
		return
	}

	// 驗證
	for _, item := range body {
		if err := request.Validate.Struct(item); err != nil {
			response.ErrorValidation(c, err)
			return
		}
	}

	err := ctrl.repo.UpdateSettings(body, group)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.UPDATE_FAILED, err))
		return
	}

	response.Success(c, nil, def.UPDATE_SUCCESS)
}

// 驗證陣列object
func validateSettingArray[T any](c *gin.Context, item m.SettingInput) error {
	var bindInput []T // 驗證的model格式

	b, err := json.Marshal(item.Value)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
		return err
	}

	if err := json.Unmarshal(b, &bindInput); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
		return err
	}

	errorsMap := map[string][]request.ErrorBag{}
	for i, s := range bindInput {
		if err := request.Validate.Struct(s); err != nil {
			if ve, ok := err.(validator.ValidationErrors); ok {
				errBags, ferr := request.FormateErrorBag(ve)
				if ferr != nil {
					response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, ferr))
					return ferr
				}
				errorsMap[strconv.Itoa(i)] = errBags
			} else {
				response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
				return err
			}
		}
	}

	if len(errorsMap) > 0 {
		response.Error(c, resource.ErrorValidation{Errors: map[string]map[string][]request.ErrorBag{item.Id: errorsMap}}, "", http.StatusUnprocessableEntity)

		return errors.New(def.BIND_BODY_FAILED)
	}

	return nil
}
