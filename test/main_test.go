package test

import (
	"encoding/json"
	"fmt"
	"gbase/src/core/db"
	"gbase/src/core/helper"
	"gbase/src/def"
	"gbase/src/http/resource"
	"gbase/src/migrate"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {

	// 執行 migration
	migrate.Run()

	// 關閉 SQL Debug
	db.DisableDebug()

	code := m.Run()

	os.Exit(code)
}

func List(
	t *testing.T,
	r *gin.Engine,
	uri string,
	key string,
	hasPagination bool,
	resourceStruct any,
	itemsStruct ...any,
) {
	url := fmt.Sprintf("/api/%s", uri)
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("列表失敗，HTTP code 應為 %d，實際收到 %d，回傳: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	CheckResponseFields(t, w, "", key, resourceStruct, itemsStruct...)

	if hasPagination {
		CheckResponseFields(t, w, "", "pagination", resourceStruct, resource.Pagination{})
	}
}

func Create(
	t *testing.T,
	r *gin.Engine,
	uri string,
	key string,
	input string,
	pkName string,
	resourceStruct any,
	itemsStruct ...any,
) any {

	url := fmt.Sprintf("/api/%s", uri)
	req, _ := http.NewRequest("POST", url, strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("新增失敗，HTTP code 應為 %d，實際收到 %d，response: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	item := CheckResponseFields(t, w, def.CREATE_SUCCESS, key, resourceStruct, itemsStruct...)

	id, ok := item[pkName]
	if !ok {
		t.Fatalf("新增回傳未取得 %s", pkName)
	}

	return id
}

func Update(
	t *testing.T,
	r *gin.Engine,
	uri string,
	key string,
	id any,
	input string,
	resourceStruct any,
	itemsStruct ...any,
) {
	url := fmt.Sprintf("/api/%s/%v", uri, id)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("新增失敗，HTTP code 應為 %d，實際收到 %d，response: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	CheckResponseFields(t, w, def.UPDATE_SUCCESS, key, resourceStruct, itemsStruct...)
}

func Delete(
	t *testing.T,
	r *gin.Engine,
	uri string,
	id any,
) {
	url := fmt.Sprintf("/api/%s/%v", uri, id)
	req, _ := http.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("刪除失敗，HTTP code 應為 %d，實際收到 %d，data: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	CheckResponseFields(t, w, def.DELETE_SUCCESS, "", nil, nil)
}

func UpdateOrder(
	t *testing.T,
	r *gin.Engine,
	uri string,
	id1 any,
	id2 any,
) {
	url := fmt.Sprintf("/api/%s/order", uri)
	input := fmt.Sprintf(`{"id1":%v, "id2":%v}`, id1, id2)
	req, _ := http.NewRequest("POST", url, strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("更新排序失敗，HTTP code 應為 %d，實際收到 %d，response: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	CheckResponseFields(t, w, def.UPDATE_SUCCESS, "", nil, nil)
}

func CheckResponseFields(
	t *testing.T,
	w *httptest.ResponseRecorder,
	assertMessage string,
	key string,
	resourceStruct any,
	itemsStruct ...any,
) map[string]interface{} {

	var responseBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("無法解析回傳的 body: %v", err)
	}

	if _, exists := responseBody["data"]; !exists {
		t.Fatalf("回傳欄位錯誤: data 欄位不存在")
	}

	if message, ok := responseBody["message"].(string); !ok || message != assertMessage {
		t.Fatalf("回傳 message 欄位錯誤: %v", responseBody["message"])
	}

	// 不比對 data 只比對 message
	if resourceStruct == nil && key == "" {
		return nil
	}

	data, ok := responseBody["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("回傳 data 欄位型別錯誤: %+v", responseBody["data"])
	}

	// data 結構應該包含的欄位
	dataFields := helper.GeStructJsonFields(resourceStruct)
	for _, field := range dataFields {
		if _, exists := data[field]; !exists {
			t.Errorf("data 裡缺少欄位: %s", field)
		}
	}

	if items, ok := data[key].([]interface{}); ok {
		// 1. data 結構的子項目資料應該包含的欄位 (列表)

		itemFields := helper.GeStructJsonFields(itemsStruct...)
		for _, item := range items {
			d, ok := item.(map[string]interface{})
			if !ok {
				t.Errorf("%s 欄位型別錯誤: %T", key, item)
				continue
			}
			for _, field := range itemFields {
				if _, exists := d[field]; !exists {
					t.Errorf("%s 缺少欄位: %s", key, field)
				}
			}
		}
		return nil
	} else if d, ok := data[key].(map[string]interface{}); ok {
		// 2. data 結構的子項目資料應該包含的欄位 (新增、編輯)

		fields := helper.GeStructJsonFields(itemsStruct...)
		// 檢查每個欄位是否存在
		for _, field := range fields {
			if _, exists := d[field]; !exists {
				t.Errorf("%s 缺少欄位: %s", key, field)
			}
		}
		return d
	}

	t.Errorf("%s 欄位型別錯誤: %T", key, data[key])
	return nil
}
