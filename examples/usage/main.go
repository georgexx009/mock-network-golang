package main

import (
	"fmt"
	"mock-network-golang/basenode"
	"mock-network-golang/network"
)

func main() {
  n := network.New()
  b1 := basenode.New("url-1.com", n)
  b2 := basenode.New("url-2.com", n)

  b1.RegisterHandlerFunc("/example", "GET", func(basenode.Request) basenode.Response {
    return basenode.Response{
      Status: "ok",
      StatusCode: 200,
    }
  })

  req := basenode.Request{
    Url: "http://url-1.com/example",
    HttpMethod: "GET",
  }
  res := b2.SendRequest(&req)
  fmt.Printf("response b2: %+v", res)
}
