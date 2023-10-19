package main

type server interface {
	handleRequest(url, method string) (int, string)
}
