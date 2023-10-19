package main

import "fmt"

const (
	appStatusURL  = "/app/status"
	createUserURl = "/create/user"
)

var (
	nginxServer *Nginx
)

func Init() {
	nginxServer = newNginxServer()
}

func main() {

	// 初始化
	Init()

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("url:%s HttpCode:%d Body:%s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("url:%s HttpCode:%d Body:%s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("url:%s HttpCode:%d Body:%s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURl, "POST")
	fmt.Printf("url:%s HttpCode:%d Body:%s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURl, "GET")
	fmt.Printf("url:%s HttpCode:%d Body:%s\n", appStatusURL, httpCode, body)

}
