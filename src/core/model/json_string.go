package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonString struct {
	Raw any
}

// 寫入資料庫時轉成 JSON 或 string
func (v JsonString) Value() (driver.Value, error) {
	if v.Raw == nil {
		return "", nil // 預設值是空字串，不是null
	}

	switch val := v.Raw.(type) {
	case string:
		return val, nil
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
}

// 從資料庫讀出時自動解析
func (v *JsonString) Scan(value any) error {
	if value == nil {
		v.Raw = nil
		return nil
	}
	switch data := value.(type) {
	case []byte:
		s := string(data)
		var obj any
		if json.Unmarshal(data, &obj) == nil {
			v.Raw = obj
		} else {
			v.Raw = s
		}
	case string:
		var obj any
		if json.Unmarshal([]byte(data), &obj) == nil {
			v.Raw = obj
		} else {
			v.Raw = data
		}
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

// 讓 JsonString 可被直接序列化回 JSON
func (v JsonString) MarshalJSON() ([]byte, error) {
	if v.Raw == nil {
		return json.Marshal("")
	}
	return json.Marshal(v.Raw)
}

// 讓 JsonString 可被反序列化
func (v *JsonString) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		// 若不是合法 JSON，就當作字串處理
		v.Raw = string(data)
		return nil
	}
	v.Raw = val
	return nil
}
