package network

import (
	"fmt"
	"io"
  "net/http"
  "net/url"
  "time"
)

func Init() {
	fmt.Print("init network v0.0.1 local 2")
}

func New() *Network {
	return &Network{registeredNodes: make(map[string]Node)}
}

type Node interface {
	ReceiveRequest(url string, httpMethod string, body io.ReadCloser, headers http.Header) Response
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

func latency(seconds int) {
  time.Sleep(time.Duration(seconds) * time.Second)
}

func (network *Network) NetworkCall(host string, url string, httpMethod string, body io.ReadCloser, headers http.Header) Response {
  // query parameters are included already in the url
	if _, ok := network.registeredNodes[host]; !ok {
    fmt.Printf("[error] node not found with host %s: %+v", host, network.registeredNodes)
    response := Response{
      Status: "url not found",
      StatusCode: 503,
    }
		return response
	}

  latency(1)

	return network.registeredNodes[host].ReceiveRequest(url, httpMethod, body, headers)
}
