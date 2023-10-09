package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	port      int = 3333
	counter   int
	instPorts = [2]int{8080, 8081}
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// perfect microservices balance strategy
func randomBalancer() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}

// This handle read gathered request, make new request
// based on them, send it to choosen service, get response from it
// and tranlsate response backwards to client
func handleProxy(w http.ResponseWriter, req *http.Request) {
	// Forward request get exactly same hader and body
	// with gathered request
	pReqHeader := http.Header{}
	copyHeader(pReqHeader, req.Header)
	pData, _ := io.ReadAll(req.Body)

	srvs := randomBalancer()

	fmt.Printf("proxy get request - %v, send to service number - %v\n", req.URL.Path, srvs)

	pReq, err := http.NewRequest(req.Method, fmt.Sprintf("http://localhost:%d%s", instPorts[srvs], req.URL.Path), bytes.NewReader(pData))
	if err != nil {
		log.Fatalf("handleProxy(): error new request - %v\n", err)
	}
	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(pReq)
	if err != nil {
		log.Fatalf("handleProxy(): error sending request - %v\n", err)
	}

	// translate response to client
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handleProxy)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	fmt.Printf("Proxy listening at port - %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running proxy: %s\n", err)
		}
	}

}
