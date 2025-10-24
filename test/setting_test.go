package test

import (
	"fmt"
	"gbase/src/http/resource"
	"gbase/src/model"
	"gbase/src/route"
	"testing"
	"time"
)

func TestSettingApi(t *testing.T) {
	r := route.SetupRouter()

	// 測試列表
	List(t, r, "settings", "settings", false, resource.Settings{}, model.Setting{})

	// 測試編輯
	code := time.Now().Unix()
	input := fmt.Sprintf(`[{"id":"ip", "value":"測試_%d"}]`, code)
	Update(t, r, "settings", "", "system", input, nil, nil)
}
