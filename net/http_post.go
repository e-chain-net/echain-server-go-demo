package net

import (
	"bytes"
	"fmt"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)


func HttpPost(url string,payload []byte) ([]byte,error){
	timeStamp := strconv.FormatInt(common.GetTimeMilli(),10);
	rsaPrivate,err := common.LoadRsaPrivateKey(common.RsaPrivate)
	if err != nil{
		return nil,err
	}
	signature,err := common.RsaSign(rsaPrivate,[]byte(common.MerchantNo + "-" + timeStamp))
	if err != nil{
		return nil,err
	}
	// Create request payload
	//payload := []byte(`{"jsonRpc": {"method":"getBlockLimit","params":[]}}`)

	//fmt.Println("url:" + url)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		//fmt.Println("Error creating HTTP request:", err)
		return nil,err
	}

	// Set content type header
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "utf-8")
	req.Header.Set("merchantNo",common.MerchantNo)
	req.Header.Set("sign",signature)
	req.Header.Set("timestamp",timeStamp)

	// Send HTTP request
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil,err
	}

	defer resp.Body.Close()

	//Print HTTP response status code and body
	//fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return nil,err
	}
	return body,nil
}
