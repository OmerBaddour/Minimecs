package key_value_worker

import (
	"fmt"
	
	"github.com/OmerBaddour/Minimecs/key_value_store/02_partitioned/common"
)

/****************************\
  Key value worker functions 
\****************************/

func KeyValueWorker(requestChannel chan common.KeyValueWorkerRequest, responseChannel chan common.KeyValueWorkerResponse, partitionNumber int) {
	keyValueStore := make(map[string]string)
	fmt.Printf("KeyValueWorker %d: Activated partition\n", partitionNumber)
	
	for {
		request := <- requestChannel
		fmt.Printf("KeyValueWorker %d: Received request: method: %s, key: %s, value: %s\n", partitionNumber, request.Method, request.Key, request.Value)

		response := common.KeyValueWorkerResponse{}
		switch request.Method {
		case "put":
			keyValueStore[request.Key] = request.Value
			response.Status = "201 Created"
		case "get":
			value, ok := keyValueStore[request.Key]
			if ok {
				response.Status = "200 Found"
				response.Value = value
			} else {
				response.Status = "404 Not Found"
			}
		case "delete":
			delete(keyValueStore, request.Key)
			response.Status = "204 Deleted Key"
		}
		fmt.Printf("KeyValueWorker %d: Sending response: status: %s, value: %s\n", partitionNumber, response.Status, response.Value)
		responseChannel <- response
	}
}