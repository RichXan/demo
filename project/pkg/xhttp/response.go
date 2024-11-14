package xhttp

import (
	"encoding/json"

	"xproject/pkg/xerror"
)

/*
list page data

	{
		"status": true,
		"code": 0,
		"message": "成功",
		"total": 25,
		"current": 1,
		"per_page": 100,
		"size": 25
		"data": [...]
	}
*/
type ListData struct {
	PageData *PageData
	Size     int         `json:"size"`              // 返回payload数据条数
	Payload  interface{} `json:"payload,omitempty"` // 成功时返回的数据
}

type PageData struct {
	Total   int `json:"total,omitempty"`    // 总条数
	Current int `json:"current,omitempty"`  // 当前页
	PerPage int `json:"per_page,omitempty"` // 每页条数
}

type ApiResponse struct {
	ID      string      `json:"id,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
	Status  bool        `json:"status"`
	Code    int         `json:"code"`              // 业务编码
	Message string      `json:"message,omitempty"` // 错误描述
	Data    interface{} `json:"data,omitempty"`    // 成功时返回的数据
	Size    *int        `json:"size,omitempty"`    // 成功时返回的数据
	PageData
}

func NewResponse(err *xerror.Error) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}

	res := &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Message: err.Msg(),
	}
	return res
}

func ResponseSuccess() *ApiResponse {
	res := &ApiResponse{
		Status: true,
		Code:   0,
	}
	return res
}

func NewResponseData(err *xerror.Error, data interface{}) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}
	resp := &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Data:    data,
		Message: err.Msg(),
	}

	if d, ok := data.(*ListData); ok {
		if d != nil {
			if d.PageData != nil {
				resp.PageData = *d.PageData
			}
			resp.Data = d.Payload
			resp.Size = &d.Size
		}
	}

	return resp
}

func NewResponseMessage(err *xerror.Error, message string) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}

	resp := &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Message: err.Msg(),
	}
	if len(message) > 0 {
		resp.Message = message
	}
	return resp
}

func (res *ApiResponse) SetData(data interface{}) *ApiResponse {
	res.Data = data
	return res
}

func (res *ApiResponse) SetID(id string) *ApiResponse {
	res.ID = id
	return res
}

func (res *ApiResponse) SetMessage(msg string) *ApiResponse {
	res.Message = msg
	return res
}

// ToString 返回 JSON 格式的错误详情
func (res *ApiResponse) ToString() string {
	err := &ApiResponse{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
		ID:      res.ID,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
