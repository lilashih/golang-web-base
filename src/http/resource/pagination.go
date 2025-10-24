package resource

type Pagination struct {
	Total       int `json:"total"`       // 資料總筆數
	PerPage     int `json:"perPage"`     // 每頁的資筆數
	CurrentPage int `json:"currentPage"` // 當前頁數
	LastPage    int `json:"lastPage"`    // 最後一頁的頁數
} //@name Pagination
