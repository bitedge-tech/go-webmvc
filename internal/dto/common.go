package dto

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Pagination struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
