package network

import (
	"errors"
	"fmt"
	"io"
)

func Init() {
	fmt.Print("init network v0.0.1 local 2")
}

func New() *Network {
	return &Network{registeredNodes: make(map[string]Node)}
}

type Node interface {
	ReceiveRequest(url string, httpMethod string, body io.Reader, headers map[string]string) (*Response, error)
}

type Network struct {
	registeredNodes map[string]Node
}

func (network *Network) RegisterNode(url string, n Node) {
	network.registeredNodes[url] = n
}

type Response struct {
	Status         string
  StatusCode int
	Body       io.ReadCloser
	statusCode int
}

func (network *Network) NetworkCall(url string, httpMethod string, body io.Reader, headers map[string]string) (*Response, error) {
  // query parameters are included already in the url
	if _, ok := network.registeredNodes[url]; !ok {
		return nil, errors.New("url not found")
	}

	return network.registeredNodes[url].ReceiveRequest(url, httpMethod, body, headers)
}
