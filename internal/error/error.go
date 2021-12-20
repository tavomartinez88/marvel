package error

import "encoding/json"

type ClientError struct {
	HttpStatus int `json:"httpStatus"`
	Message string `json:"msg"`
}

func (age ClientError) Error() string {
	js, _ :=  json.Marshal(age)
	return string(js)
}