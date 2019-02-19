package main

import (
  "flag"
  "google.golang.org/grpc"
  "github.com/golang/protobuf/proto"
  "net"
  pb "./genproto/demo"
)

var (
  tls = flag.Bool("tls", false," Tls configuration options")
  certFile = flag.String("cert", "", "Cert file")
  keyFile = flag.String("Key", "", "Key file")
)
