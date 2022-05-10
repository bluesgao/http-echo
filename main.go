package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var port string
var ip string

func getLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		//log.Println("ip:", ipAddr.IP.String())
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

func main() {
	// index 为向 url发送请求时，调用的函数
	http.HandleFunc("/", index)

	//port
	port = os.Getenv("HTTP_ECHO_PORT")
	if port == "" {
		port = "8888"
	}
	//ip
	ip, _ = getLocalIP()

	addr := fmt.Sprintf(":%s", port)
	log.Println("listen and serve ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func mockResponseData(r *http.Request) interface{}{
	//req info
	req := make(map[string]interface{})
	req["path"] = r.URL.Path
	req["method"] = r.Method
	req["host"] = r.Host
	req["referer"] = r.Referer()
	req["header"] = fmt.Sprintf("%v", r.Header)

	//server info
	server := make(map[string]interface{})
	server["ip"] = ip
	server["port"] = port

	//result info
	result := make(map[string]interface{})
	result["time"] = time.Now().Format("2006/1/2 15:04:05.000")
	result["code"] = 0
	result["msg"] = "success"

	data := make(map[string]interface{})
	data["request"] = req
	data["server"] = server

	result["data"] = data

	return result
}

func index(w http.ResponseWriter, r *http.Request) {
	//json
	jsonData, _ := json.Marshal(mockResponseData(r))
	fmt.Fprintf(w, "%s", jsonData)
}
