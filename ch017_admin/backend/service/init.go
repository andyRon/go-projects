package service

import "github.com/andyron/mini-admin/define"

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		Keyword: "",
	}
}
