package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
)

func main() {

	//Create a web service
	webService := new(restful.WebService)

	//Create a route and attach it to handler in the service
	webService.Route(webService.GET("/ping").To(pingTime))

	//Add the service to application
	restful.Add(webService)

	http.ListenAndServe(":8000", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {

	//Write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
