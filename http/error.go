package http

import (
	"fmt"
	"github.com/kataras/iris"
)

type HttpError struct {
	Message string            `json:"message"`
	Status  int               `json:"status"`
	Data    map[string]string `json:"data"`
}

func (h HttpError) Error() string {
	return fmt.Sprintf("http error (%s): %d", h.Message, h.Status)
}

func (h HttpError) ToIrisMap() iris.Map {
	return iris.Map{"status": h.Status, "message": h.Message, "data": h.Data}
}
