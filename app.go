package app

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/y16ra/cloud-functions-hello-go/hello"
)

const entryPoint = "HelloWorld"

func init() {
	functions.HTTP(entryPoint, hello.HelloHTTP)
}
