package basenode

import (
  "io"
  "net/http"
  urlPkg "net/url"
  
  "mock-network-golang/network"
)

type Basenode struct {
  HostUrl string
  Network *network.Network
  RestApi map[string]map[string]func(w http.ResponseWriter, r *http.Request)(network.Response)
}

func New(hostUrl string, n *network.Network) *Basenode {
  restApi := make(map[string]map[string]func(w http.ResponseWriter, r *http.Request)(network.Response))
  baseNode := &Basenode{hostUrl, n, restApi}
  n.RegisterNode(baseNode.HostUrl, baseNode)
  return baseNode
}

// ReceiveRequest - receive http request
func (baseNode *Basenode) ReceiveRequest(url string, httpMethod string, body io.ReadCloser, headers map[string]string) network.Response {
  u, err := urlPkg.Parse(url) 
  if err != nil {
    response := network.Response{
      Status: "not found",
      StatusCode: 404,
    }
    return response
  }

  path := u.Path

  // check path is supported
  if _, ok := baseNode.RestApi[path]; !ok {
    response := network.Response{
      Status: "path not found",
      StatusCode: 404,
    }
    return response
  }

  // check http method is supported by path
  if _, ok := baseNode.RestApi[path][httpMethod]; !ok {
    response := network.Response{
      Status: "path not found",
      StatusCode: 404,
    }
    return response
  }

  requestHeaders := http.Header{}
  for k, v := range headers {
    requestHeaders.Add(k, v)
  }

  var w http.ResponseWriter
  var r http.Request
  r.Method = httpMethod
  r.URL = u
  r.Header = requestHeaders
  r.Body = body

  return baseNode.RestApi[path][httpMethod](w, &r)
}

// SendRequest - send an http request
func (baseNode *Basenode) SendRequest(url string, httpMethod string, body io.ReadCloser, headers map[string]string) (network.Response) {
  return baseNode.Network.NetworkCall(url, httpMethod, body, headers)
}

func (baseNode *Basenode) RegisterHandlerFunc(path string, httpVerb string, handler func(w http.ResponseWriter, r *http.Request)(network.Response)) {
  // path already exists
  if pathMap, ok := baseNode.RestApi[path]; ok {
    pathMap[httpVerb] = handler
    return
  }

  // create path map and then save it
  pathMap := make(map[string]func(w http.ResponseWriter, r *http.Request)(network.Response))
  pathMap[httpVerb] = handler
  baseNode.RestApi[path] = pathMap
}
