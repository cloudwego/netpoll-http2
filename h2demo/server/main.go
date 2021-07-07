// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/netpoll"
	"github.com/cloudwego/netpoll-http2"
)

func main() {
	network, address := "tcp", "127.0.0.1:8888"

	// 创建 listener
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic("create netpoll listener fail")
	}

	server := netpoll_http2.Server{Handler: &serverHandler{}, IdleTimeout: time.Minute}
	// handle: 连接读数据和处理逻辑
	var onRequest netpoll.OnRequest = server.ServeConn

	// options: EventLoop 初始化自定义配置项
	opts := []netpoll.Option{
		// netpoll.WithReadTimeout(3 * time.Minute), // netpoll bug, wait fix
		netpoll.WithIdleTimeout(10 * time.Minute),
		netpoll.WithOnPrepare(func(conn netpoll.Connection) context.Context {
			conn.SetReadTimeout(3 * time.Minute)
			return context.Background()
		}),
	}

	// 创建 EventLoop
	eventLoop, err := netpoll.NewEventLoop(onRequest, opts...)
	if err != nil {
		panic("create netpoll event-loop fail")
	}

	// 运行 Server
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("netpoll server exit")
	}
}

// http handler
type serverHandler struct{}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	w.Header().Set("server", "h2test")
	fmt.Fprintf(w, "Hello, %v, http: %v", req.URL.Path, req.TLS == nil)
}
