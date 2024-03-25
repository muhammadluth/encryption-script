package model

import (
	"net/http"
	"time"
)

type (
	ServiceProperties struct {
		ServiceName               string         `json:"service_name" validate:"required"`
		ServicePort               int            `json:"service_port" validate:"numeric,required"`
		ServicePoolSizeConnection int            `json:"service_pool_size_connection" validate:"numeric,required"`
		ServiceTimezone           *time.Location `json:"service_timezone" validate:"required"`
		ServiceDebugMode          bool           `json:"service_debug_mode"`
	}
)

type (
	Message struct {
		Header Header `json:"header"`
		Body   []byte `json:"body"`
		Param  []byte `json:"param"`
		Query  []byte `json:"query"`
	}
	Header struct {
		URL    string      `json:"url"`
		Query  string      `json:"query"`
		Header http.Header `json:"header"`
	}
)

type Response struct {
	Status int         `json:"status"`
	Header http.Header `json:"header"`
	Body   interface{} `json:"body"`
}

type (
	ResponseDefault struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	ResponseData struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func FResponseDefault(status int, message string) Response {
	resBody := ResponseDefault{
		Status:  status,
		Message: message,
	}
	return Response{status, http.Header{}, resBody}
}

func FResponseData(status int, message string, data interface{}) Response {
	resBody := ResponseData{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return Response{status, http.Header{}, resBody}
}
