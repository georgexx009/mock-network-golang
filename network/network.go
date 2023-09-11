package network

import (
	"errors"
	"fmt"
)

func Init() {
	fmt.Print("init network v0.0.1 local 2")
}

func New() *Network {
	return &Network{registeredNodes: make(map[string]Node)}
}

type Node interface {
	receiveRequest(url string, httpMethod string, body string, headers string) (*Response, error)
}

type Network struct {
	registeredNodes map[string]Node
}

func (network *Network) RegisterNode(url string, n Node) {
	network.registeredNodes[url] = n
}

type Response struct {
	ok         bool
	body       string
	statusCode int
}

func (network *Network) Send(url string, httpMethod string, body string, headers string) (*Response, error) {
	if _, ok := network.registeredNodes[url]; !ok {
		return nil, errors.New("url not found")
	}

	return network.registeredNodes[url].receiveRequest(url, httpMethod, body, headers)
}
