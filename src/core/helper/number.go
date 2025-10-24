package helper

import (
	"encoding/json"
	"strings"
)

// 把數字的一維陣列轉換為資料庫相容的 JSON 格式
func IntToJson(v []int) interface{} {
	if len(v) == 0 {
		return nil // return nil, for storing NULL in SQLite
	}

	bytes, _ := json.Marshal(v)

	return string(bytes)
}

// 把字串 '[1,23,3,2]' 轉成用 separator 分隔的字串 '1, 23, 3, 2'
func Implode(jsonStr string, separator string) string {
	var rawNumbers []json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &rawNumbers); err != nil {
		return ""
	}

	if separator == "" {
		separator = ", "
	}

	numberStr := make([]string, len(rawNumbers))
	for i, raw := range rawNumbers {
		var s string
		// try to parse as a string
		if err := json.Unmarshal(raw, &s); err == nil {
			numberStr[i] = s
			continue
		}
		// if parsing  as a string fails, treat it as a number
		var n json.Number
		if err := json.Unmarshal(raw, &n); err == nil {
			numberStr[i] = n.String()
		}
	}

	return strings.Join(numberStr, separator)
}
