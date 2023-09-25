package main

import (
  "mock-network-golang/basenode"
  "mock-network-golang/network"
)

func main() {
  n := network.New()
  b1 := basenode.New("url-1.com", n)
  b2 := basenode.New("url-2.com", n)

  b1.RegisterHandlerFunc("/example", "GET", func() {

  })
}
