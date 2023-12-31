package main

import (
	"net/http"

	"github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common"
	ws "github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/web_server"
	kvw "github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/key_value_worker"
)

func main() {

	requestChannels := make([]chan common.KeyValueWorkerRequest, common.NUM_PARTITIONS)
	responseChannels := make([]chan common.KeyValueWorkerResponse, common.NUM_PARTITIONS)
	handler := &ws.WebServer{
		RequestChannels: requestChannels,
		ResponseChannels: responseChannels,
	}

	for i := 0; i < common.NUM_PARTITIONS; i++ {
		handler.RequestChannels[i] = make(chan common.KeyValueWorkerRequest)
		handler.ResponseChannels[i] = make(chan common.KeyValueWorkerResponse)
		key_value_worker := &kvw.KeyValueWorker{
			KeyValueStore: make(map[string]string),
			RequestChannel: handler.RequestChannels[i],
			ResponseChannel: handler.ResponseChannels[i],
			PartitionNumber: i,
		}
		go key_value_worker.DoWork()
	}

	http.HandleFunc("/put", handler.HttpPut)
	http.HandleFunc("/get", handler.HttpGet)
	http.HandleFunc("/delete", handler.HttpDelete)

	http.ListenAndServe(":3000", nil)
}