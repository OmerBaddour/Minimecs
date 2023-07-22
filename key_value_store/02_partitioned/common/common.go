package common

const NUM_PARTITIONS = 4

type KeyValueWorkerRequest struct {
	Method string
	Key string
	Value string
}

type KeyValueWorkerResponse struct {
	Status string
	Value string
}
