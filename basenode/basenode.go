package basenode

import (
	"io"
  "log/slog"
	"net/http"
	urlPkg "net/url"
  "time"

	"mock-network-golang/network"
)

type Response struct {
  Status         string
  StatusCode int
	Body       io.ReadCloser
}

// MARK
type Request struct {
  Url string
  HttpMethod string
  Body io.ReadCloser
  Headers http.Header
}

type Basenode struct {
  HostUrl string
  Network *network.Network
  RestApi map[string]map[string]func(req Request) Response
  logger *slog.Logger
}

func New(hostUrl string, n *network.Network) *Basenode {
  restApi := make(map[string]map[string]func(req Request) Response)
  baseNode := &Basenode{hostUrl, n, restApi, slog.With("pkg","basenode")}
  n.RegisterNode(baseNode.HostUrl, baseNode)
  return baseNode
}

// ReceiveRequest - receive http request
func (baseNode *Basenode) ReceiveRequest(url string, httpMethod string, body io.ReadCloser, headers http.Header) network.Response {
  logger := baseNode.logger.With("method", "ReceiveRequest")

  u, err := urlPkg.Parse(url) 
  if err != nil {
    logger.Error("error parsing the url", "url", url)
    response := network.Response{
      Status: "not found",
      StatusCode: 404,
    }
    return response
  }

  path := u.Path

  // check path is supported
  if _, ok := baseNode.RestApi[path]; !ok {
    logger.Error("path not found")
    response := network.Response{
      Status: "path not found",
      StatusCode: 404,
    }
    return response
  }

  // check http method is supported by path
  if _, ok := baseNode.RestApi[path][httpMethod]; !ok {
    logger.Error("http method not supported", "httpMethod", httpMethod, "path", path)
    response := network.Response{
      Status: "path not found",
      StatusCode: 404,
    }
    return response
  }

  req := Request{
    Body: body,
    HttpMethod: httpMethod,
    Url: u.String(),
    Headers: headers,
  }
  logger.Info("request received", "req", req)

  simulateWorkingTime(3)
  baseNodeResponse := baseNode.RestApi[path][httpMethod](req)
  logger.Info("response", "res", baseNodeResponse)

  return network.Response{
    Status: baseNodeResponse.Status,
    StatusCode: baseNodeResponse.StatusCode,
    Body: baseNodeResponse.Body,
  }
}

// SendRequest - send an http request
func (baseNode *Basenode) SendRequest(req *Request) Response {
  logger := baseNode.logger.With("method", "SendRequest")
  u, err := urlPkg.Parse(req.Url)
  if err != nil {
    return Response{
      Status: "bad url",
      StatusCode: 400,
    }
  }
  logger.Info("sending request", "from", baseNode.HostUrl, "to", req.Url)

  res := baseNode.Network.NetworkCall(u.Hostname(), req.Url, req.HttpMethod, req.Body, req.Headers)
  return Response{Status: res.Status, StatusCode: res.StatusCode, Body: res.Body}
}

func (baseNode *Basenode) RegisterHandlerFunc(path string, httpVerb string, handler func(req Request) Response) {
  // path already exists
  if pathMap, ok := baseNode.RestApi[path]; ok {
    pathMap[httpVerb] = handler
    return
  }

  // create path map and then save it
  pathMap := make(map[string]func(req Request) Response)
  pathMap[httpVerb] = handler
  baseNode.RestApi[path] = pathMap
}

func simulateWorkingTime(secs int) {
  time.Sleep(time.Duration(secs) * time.Second)
}
