package def

import "errors"

var (
	ErrRecordIsTrashed = errors.New("此資料已被刪除，禁止修改")
)
