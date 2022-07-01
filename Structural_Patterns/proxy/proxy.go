package main

import (
	"fmt"
	"net/http"
)


// ########## NGINX ##########

type nginx struct {
	application *application
	rateLimiter map[string]int
	maxRequest  int
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	if cnt := n.rateLimiter[url]; cnt >= n.maxRequest {
		return 403, "NOT ALLOWED"
	}

	n.rateLimiter[url]++
	return n.application.handleRequest(url, method)
}

func newNginx() *nginx {
	return &nginx{
		application: &application{},
		rateLimiter: map[string]int{},
		maxRequest:  3,
	}
}

// ########## SERVER ##########


type server interface {
	handleRequest(string, string) (int, string)
}

// ########## APP ##########

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}
	if url == "/app/user" && method == "POST" {
		return 201, "CREATED"
	}

	return 404, "NOT FOUND"
}



func main() {
	n := newNginx()
	appStatusURL := "/app/status"
	createUserURL := "/app/user"

	statusCode, respBody := n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(appStatusURL, http.MethodGet)
	fmt.Println(appStatusURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(createUserURL, http.MethodGet)
	fmt.Println(createUserURL, http.MethodGet, statusCode, respBody)

	statusCode, respBody = n.handleRequest(createUserURL, http.MethodPost)
	fmt.Println(createUserURL, http.MethodPost, statusCode, respBody)
}
