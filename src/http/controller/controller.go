package controller

import (
	"errors"
	"fmt"
	"gbase/src/core/http/request"
	"gbase/src/core/http/response"
	"gbase/src/def"
	"gbase/src/http/resource"
	"gbase/src/model"
	"gbase/src/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Controller struct {
	repo repository.IRepository
	res  resource.IResource
}

func NewBaseController(repo repository.IRepository, res resource.IResource) *Controller {
	return &Controller{repo: repo, res: res}
}

func (ctrl *Controller) Index(c *gin.Context) {
	search := map[string]string{}
	for key, val := range c.Request.URL.Query() {
		if len(val) > 0 {
			search[key] = val[0]
		}
	}

	list, pagination, err := ctrl.repo.FindAll(search)
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Collection(pagination, list))
}

func (ctrl *Controller) Show(c *gin.Context) {
	id := request.GetParamId(c)

	result, err := ctrl.repo.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, def.ErrRecordIsTrashed) {
			response.ErrorNotFound(c)
			return
		}
		response.Error(c, nil, fmt.Sprintf("查詢失敗: %v", err))
		return
	}

	response.Success(c, ctrl.res.Single(result))
}

func (ctrl *Controller) Store(c *gin.Context) {
	model, err := ctrl.bindModel(c)
	if err != nil {
		return
	}

	// 新增
	model, err = ctrl.repo.Create(model)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.CREATE_FAILED, err))
		return
	}

	response.Success(c, ctrl.res.Single(model), def.CREATE_SUCCESS)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id := request.GetParamId(c)

	model, err := ctrl.bindModel(c)
	if err != nil {
		return
	}

	// 編輯
	model, err = ctrl.repo.Update(model, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, def.ErrRecordIsTrashed) {
			response.ErrorNotFound(c)
			return
		}
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.UPDATE_FAILED, err))
		return
	}

	response.Success(c, ctrl.res.Single(model), def.UPDATE_SUCCESS)
}

func (ctrl *Controller) Destroy(c *gin.Context) {
	id := request.GetParamId(c)

	if err := ctrl.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, def.ErrRecordIsTrashed) {
			response.ErrorNotFound(c)
			return
		}
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.DELETE_FAILED, err))
		return
	}

	response.Success(c, nil, def.DELETE_SUCCESS)
}

func (ctrl *Controller) UpdateOrder(c *gin.Context) {
	var input model.OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, "參數解析錯誤")
		return
	}

	// 驗證
	if err := request.Validate.Struct(input); err != nil {
		response.ErrorValidation(c, err)
		return
	}

	// 交換順序
	if err := ctrl.repo.UpdateOrder(input.Id1, input.Id2); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.UPDATE_FAILED, err))
		return
	}

	response.Success(c, nil, def.UPDATE_SUCCESS)
}

// 把新增/編輯的值綁到model上，從請求的body中解析出資料，並進行驗證
func (ctrl *Controller) bindModel(c *gin.Context) (model.IModel, error) {
	body := ctrl.repo.NewInput()

	// get body data
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
		return nil, err
	}

	// 驗證
	if err := request.Validate.Struct(body); err != nil {
		response.ErrorValidation(c, err)
		return nil, err
	}

	// 複製到model
	model := ctrl.repo.NewModel()
	if err := copier.Copy(model, body); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_MODEL_FAILED, err))
		return nil, err
	}

	return model, nil
}
