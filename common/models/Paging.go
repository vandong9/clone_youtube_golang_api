package models

type Paging struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type PagingQuery struct {
	PageIndex int `form:"page_index"`
	PageSize  int `form:"page_size"`
}
