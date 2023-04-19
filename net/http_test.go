package net

import (
	"bytes"
	"fmt"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)


func TestHttp(t *testing.T) {
	// Create request payload
	payload := []byte(`{"jsonRpc": {"method":"getBlockLimit","params":[]}}`)

	responseBody,err := HttpPost(common.UrlQuery,payload)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("Response Body:", string(responseBody))
}

func TestHttpBase(t *testing.T) {
	// Create request payload
	payload := []byte(`{"jsonRpc": {"method":"getBlockNumber","params":[]}}`)

	fmt.Println("url:" + common.UrlQuery)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", common.UrlQuery, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set content type header
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "utf-8")
	req.Header.Set("merchantNo","1020000189025321")
	req.Header.Set("sign","RgdvfNBvfkccRgnOC2y5Zcfyh5myAYZnTGS6mhBCv0t0rIQfhLBqLVge7S931/1QavD9X0G4E0wTB1Dl2iMcWGZABC+HrtgZtndigUwNaIJhxkX+drkxPeFZROjwk02xvfx3deMvAHqMYx3EZVmTB6x3Og/7Z7HEKt3mGWNuDh/usS+SV8CG+iRrFDSlnCTJc4N6/6B+bFT2O2U5a0wCwgaIg5q3mrNgSMQ7W3WxGue8/DxonYNUUupqQltwUH56SX6WKj0sy87ahAJe1fdIuNH/DN9R/guOqqLYARMYyC5bStqj+lfKm9EkIUahOoy3QWG5jVNG7g6wqxbSYXjHSg==")
	req.Header.Set("timestamp","1681178471210")

	// Send HTTP request
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	defer resp.Body.Close()

	// Print HTTP response status code and body
	fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
}
