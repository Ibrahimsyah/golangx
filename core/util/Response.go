package util

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}
