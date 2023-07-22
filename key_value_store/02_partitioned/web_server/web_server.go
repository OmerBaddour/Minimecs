package web_server

import (
	"fmt"
	"net/http"
	"crypto/sha256"
	"encoding/binary"
	
	"github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common"
)

/**************************\
  WebServer implementation 
\**************************/

type WebServer struct {
	RequestChannels []chan common.KeyValueWorkerRequest
	ResponseChannels []chan common.KeyValueWorkerResponse
}

func getPartitionIndex(key string) uint64 {
	hash := sha256.Sum256([]byte(key))
	hashInt := binary.BigEndian.Uint64(hash[:8])
	return hashInt % common.NUM_PARTITIONS
}

func (web_server *WebServer) handleTransaction(method string, key string, value string) (string, string) {
	fmt.Println("WebServer: Handling transaction")
	request := common.KeyValueWorkerRequest{
		Method: method,
		Key: key,
		Value: value,
	}
	index := getPartitionIndex(key)
	fmt.Printf("WebServer: Sending request to worker %d\n", index)
	web_server.RequestChannels[index] <- request
	response := <- web_server.ResponseChannels[index]
	fmt.Printf("WebServer: Received response from worker %d\n", index)
	return response.Status, response.Value
}

func (web_server *WebServer) HttpPut(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	if len(key) > 0 && len(value) > 0 {
		status, _ := web_server.handleTransaction("put", key, value)
		fmt.Fprintf(w, status)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func (web_server *WebServer) HttpGet(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if len(key) > 0 {
		status, value := web_server.handleTransaction("get", key, "")
		fmt.Fprintf(w, status, value)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func (web_server *WebServer) HttpDelete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if len(key) > 0 {
		status, _ := web_server.handleTransaction("delete", key, "")
		fmt.Fprintf(w, status)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}