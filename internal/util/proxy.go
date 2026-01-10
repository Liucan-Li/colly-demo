package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Proxy struct {
	Proxy string
}

func GetProxyList() ([]string, error) {
	resp, err := http.Get("http://127.0.0.1:5010/all?type=sock5")
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var proxyList []Proxy
	// 使用json解码器直接解析响应体
	err = json.NewDecoder(resp.Body).Decode(&proxyList)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return nil, err
	}
	log.Printf("proxy list: %v\n", proxyList)

	list := make([]string, 0)
	for _, proxy := range proxyList {
		list = append(list, "sock5://"+proxy.Proxy)
	}
	return list, nil
}
