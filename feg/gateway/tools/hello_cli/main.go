/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/registry"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <message> [code (0:16)]\n", os.Args[0])
	}
	flag.Parse()

	msg := flag.Arg(0)
	if len(msg) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	code, _ := strconv.Atoi(flag.Arg(1))
	conn, err := registry.Get().GetCloudConnection(strings.ToLower(registry.FEG_HELLO))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cl := protos.NewHelloClient(conn)
	fmt.Printf("Sending  Greeting: '%s', Code: %d\n", msg, code)
	res, err := cl.SayHello(context.Background(), &protos.HelloRequest{Greeting: msg, GrpcErrCode: uint32(code)})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received Greeting: '%s'\n", res.GetGreeting())
}