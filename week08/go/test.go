package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if resp != nil {
		defer func() {
			resp.Body.Close()
			fmt.Println("exit")
		}()
	}

	bytes, err := io.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(bytes)
	}
}
