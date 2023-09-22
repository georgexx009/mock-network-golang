package basenode

import "mock-network-golang/network"
import "io"

type Basenode struct {
  HostUrl string
  Network *network.Network
}

func (baseNode *Basenode) ReceiveRequest(url string, httpMethod string, body io.Reader, headers map[string]string) (*network.Response, error) {
  return &network.Response{}, nil
}

func (baseNode *Basenode) SendRequest(url string, httpMethod string, body io.Reader, headers map[string]string) (*network.Response, error) {
  return baseNode.Network.NetworkCall(url, httpMethod, body, headers)
}

func New(hostUrl string, n *network.Network) *Basenode {
  baseNode := &Basenode{hostUrl, n}
  n.RegisterNode(baseNode.HostUrl, baseNode)
  return baseNode
}

