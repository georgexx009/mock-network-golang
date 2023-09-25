package network

import (
	"fmt"
	"io"
  "net/http"
  "net/url"
)

func Init() {
	fmt.Print("init network v0.0.1 local 2")
}

func New() *Network {
	return &Network{registeredNodes: make(map[string]Node)}
}

type Node interface {
	ReceiveRequest(url string, httpMethod string, body io.ReadCloser, headers map[string]string) Response
}

type Network struct {
	registeredNodes map[string]Node
}

func (network *Network) RegisterNode(url string, n Node) {
	network.registeredNodes[url] = n
}

type Request struct {
  Method string
  Url url.URL
  Headers http.Header
  Body io.ReadCloser
}

type Response struct {
	Status         string
  StatusCode int
	Body       io.ReadCloser
}

func (network *Network) NetworkCall(url string, httpMethod string, body io.ReadCloser, headers map[string]string) Response {
  // query parameters are included already in the url
	if _, ok := network.registeredNodes[url]; !ok {
    response := Response{
      Status: "url not found",
      StatusCode: 503,
    }
		return response
	}

	return network.registeredNodes[url].ReceiveRequest(url, httpMethod, body, headers)
}
