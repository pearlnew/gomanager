package utils

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io"
)

type HttpResponse struct {
	Code int `json:"code"`
	ErrMsg string `json:"msg"`
	Data interface{} `json:"data"`
}

func RespondStruct(response *restful.Response, statusCode int, data interface{}, err error) {
	var errString string
	if err != nil {
		response.WriteHeader(500)
		errString = err.Error()
	}
	ret := HttpResponse{Code: statusCode, Data: data, ErrMsg: errString}
	//response.AddHeader("content-type", "application/json")
	jsonByte, err := json.Marshal(ret)
	if err != nil {
		response.WriteHeader(500)
		io.WriteString(response, `{"code": 500, "msg":"internal server error"}`)
	} else {
		response.WriteHeader(statusCode)
		io.WriteString(response, string(jsonByte))
	}
}
