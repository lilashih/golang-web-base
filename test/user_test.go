package test

import (
	"fmt"
	"gbase/src/http/resource"
	"gbase/src/model"
	"gbase/src/route"
	"testing"
	"time"
)

func TestUserAPI(t *testing.T) {
	m := model.User{}
	r := route.SetupRouter()

	// 測試新增第一筆資料
	code1 := time.Now().Unix()
	input1 := fmt.Sprintf(`{"name":"name_%d"}`, code1)
	id1 := Create(t, r, "users", "user", input1, m.GetKeyName(), resource.User{}, model.UserInput{}, model.User{})

	// 測試新增第二筆資料
	code2 := time.Now().Unix() + 1
	input2 := fmt.Sprintf(`{"name":"name_%d"}`, code2)
	id2 := Create(t, r, "users", "user", input2, m.GetKeyName(), resource.User{}, model.UserInput{}, model.User{})

	// 測試列表
	List(t, r, "users", "users", true, resource.Users{}, model.UserInput{}, model.User{})

	// 測試編輯，用新增拿到的id
	code1 = time.Now().Unix()
	input1 = fmt.Sprintf(`{"name":"name_%d"}`, code1)
	Update(t, r, "users", "user", id1, input1, resource.User{}, model.UserInput{}, model.User{})

	// 測試更新排序
	UpdateOrder(t, r, "users", id1, id2)

	// 測試刪除
	Delete(t, r, "users", id1)
	Delete(t, r, "users", id2)
}
