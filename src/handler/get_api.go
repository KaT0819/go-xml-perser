package handler

import (
	"go-xml-parser/src/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetTest(req operations.GetTestParams) middleware.Responder {

	response := "テストです"

	return operations.NewGetTestOK().WithPayload(response).WithAccessControlAllowHeaders("*").WithAccessControlAllowOrigin("*").WithAccessControlAllowMethods("GET, POST, PUT, DELETE, OPTIONS")
}
