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
	"fmt"
	"io/ioutil"
	"net/http"

	http2 "github.com/cloudwego/nhttp2"
)

func main() {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8888/foobar", nil)

	tr := &http2.Transport{
		AllowHTTP: true,
	}
	res, err := tr.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.ProtoMajor != 2 {
		fmt.Println("proto not h2c")
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if got := res.Header.Get("server"); got != "h2test" {
		fmt.Errorf("header server got %v, want %v", got, "h2test")
		return
	} else {
		fmt.Printf("recv headers: %+v\n", res.Header)
	}
	if got, want := string(body), "Hello, /foobar, http: true"; got != want {
		fmt.Errorf("response got %v, want %v", got, want)
	} else {
		fmt.Printf("recv body: %+v\n", string(body))
	}
}
