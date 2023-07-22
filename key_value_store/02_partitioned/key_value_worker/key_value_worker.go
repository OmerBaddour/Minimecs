package key_value_worker

import (
	"fmt"

	"github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common"
)

/*********************************\
  Key value worker implementation
\*********************************/

type KeyValueWorker struct {
	KeyValueStore map[string]string
	RequestChannel chan common.KeyValueWorkerRequest
	ResponseChannel chan common.KeyValueWorkerResponse
	PartitionNumber int
}

func (worker *KeyValueWorker) DoWork() {
	fmt.Printf("KeyValueWorker %d: Activated partition\n", worker.PartitionNumber)

	for {
		request := <- worker.RequestChannel
		fmt.Printf("KeyValueWorker %d: Received request: method: %s, key: %s, value: %s\n", worker.PartitionNumber, request.Method, request.Key, request.Value)

		response := common.KeyValueWorkerResponse{}
		switch request.Method {
		case "put":
			worker.KeyValueStore[request.Key] = request.Value
			response.Status = "201 Created"
		case "get":
			value, ok := worker.KeyValueStore[request.Key]
			if ok {
				response.Status = "200 Found"
				response.Value = value
			} else {
				response.Status = "404 Not Found"
			}
		case "delete":
			delete(worker.KeyValueStore, request.Key)
			response.Status = "204 Deleted Key"
		}
		fmt.Printf("KeyValueWorker %d: Sending response: status: %s, value: %s\n", worker.PartitionNumber, response.Status, response.Value)
		worker.ResponseChannel <- response
	}
}