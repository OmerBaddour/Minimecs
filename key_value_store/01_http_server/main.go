package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ConcurrentMapHandler struct {
	concurrentMap *sync.Map
}

func (handler *ConcurrentMapHandler) httpGet(w http.ResponseWriter, req *http.Request) {
	value, found := handler.concurrentMap.Load(req.URL.Query().Get("key"))
	if found {
		fmt.Fprintf(w, "200 Found", value)
	} else {
		fmt.Fprintf(w, "204 No Content")
	}
}

func main() {

	concurrentMap := sync.Map{}

	// TODO: delete after implementing /put
	concurrentMap.Store("test_key", "test_value")
	fmt.Println("stored test_key")

	handler := &ConcurrentMapHandler{
		concurrentMap: &concurrentMap,
	}

	http.HandleFunc("/get", handler.httpGet)
	// TODO: implement http.HandleFunc("/put", put)

	http.ListenAndServe(":3000", nil)
}