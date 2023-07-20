package main

import (
	"fmt"
	"net/http"
	"crypto/sha256"
	"encoding/binary"
)

const NUM_PARTITIONS = 4

/******************\
  Worker functions 
\******************/

type KeyValueWorkerRequest struct {
	Method string
	Key string
	Value string
}

type KeyValueWorkerResponse struct {
	Status string
	Value string
}

func keyValueWorker(requestChannel chan KeyValueWorkerRequest, responseChannel chan KeyValueWorkerResponse) {
	partition := make(map[string]string)
	fmt.Println("keyValueWorker: Activated partition")
	
	for {
		request := <- requestChannel
		fmt.Println("keyValueWorker: Received request", request.Method, request.Key, request.Value)

		response := KeyValueWorkerResponse{}
		switch request.Method {
		case "put":
			partition[request.Key] = request.Value
			response.Status = "201 Created"
		case "get":
			value, ok := partition[request.Key]
			if ok {
				response.Status = "200 Found"
				response.Value = value
			} else {
				response.Status = "404 Not Found"
			}
		case "delete":
			delete(partition, request.Key)
			response.Status = "204 Deleted Key"
		}
		fmt.Println("keyValueWorker: Sending response", response.Status, response.Value)
		responseChannel <- response
	}
}

/****************************\
  HttpHandler implementation 
\****************************/

type HttpHandler struct {
	RequestChannels []chan KeyValueWorkerRequest
	ResponseChannels []chan KeyValueWorkerResponse
}

func getPartitionIndex(key string) uint64 {
	hash := sha256.Sum256([]byte(key))
	hashInt := binary.BigEndian.Uint64(hash[:8])
	return hashInt % NUM_PARTITIONS
}

func (handler *HttpHandler) handleTransaction(method string, key string, value string) (string, string) {
	fmt.Println("HttpHandler: Handling transaction")
	request := KeyValueWorkerRequest{
		Method: method,
		Key: key,
		Value: value,
	}
	index := getPartitionIndex(key)
	fmt.Println("HttpHandler: Sending request to request channel", request)
	handler.RequestChannels[index] <- request
	response := <- handler.ResponseChannels[index]
	fmt.Println("HttpHandler: Received response from response channel", response)
	return response.Status, response.Value
}

func (handler *HttpHandler) httpPut(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	if len(key) > 0 && len(value) > 0 {
		status, _ := handler.handleTransaction("put", key, value)
		fmt.Fprintf(w, status)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func (handler *HttpHandler) httpGet(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if len(key) > 0 {
		status, value := handler.handleTransaction("get", key, "")
		fmt.Fprintf(w, status, value)
	} else {
		fmt.Fprintf(w, "404 Not Found")
	}
}

func (handler *HttpHandler) httpDelete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if len(key) > 0 {
		status, _ := handler.handleTransaction("delete", key, "")
		fmt.Fprintf(w, status)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func main() {

	requestChannels := make([]chan KeyValueWorkerRequest, NUM_PARTITIONS)
	responseChannels := make([]chan KeyValueWorkerResponse, NUM_PARTITIONS)
	handler := &HttpHandler{
		RequestChannels: requestChannels,
		ResponseChannels: responseChannels,
	}

	for i := 0; i < NUM_PARTITIONS; i++ {
		handler.RequestChannels[i] = make(chan KeyValueWorkerRequest)
		handler.ResponseChannels[i] = make(chan KeyValueWorkerResponse)
		go keyValueWorker(handler.RequestChannels[i], handler.ResponseChannels[i])
	}

	http.HandleFunc("/put", handler.httpPut)
	http.HandleFunc("/get", handler.httpGet)
	http.HandleFunc("/delete", handler.httpDelete)

	http.ListenAndServe(":3000", nil)
}