// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
}

type Response struct {
	Status  int8        `json:"status"`
	Message string      `json:"message"`
	Values  interface{} `json:"values"`
}