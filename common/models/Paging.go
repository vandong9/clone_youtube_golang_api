package models

type Paging struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
