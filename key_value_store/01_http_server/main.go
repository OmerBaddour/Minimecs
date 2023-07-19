package main

import (
	"fmt"
	"net/http"
	"sync"
)

type ConcurrentMapHandler struct {
	concurrentMap *sync.Map
}

func (handler *ConcurrentMapHandler) httpPut(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	if len(key) > 0 && len(value) > 0 {
		handler.concurrentMap.Store(key, value)
		fmt.Fprintf(w, "201 Created", key, value)
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func (handler *ConcurrentMapHandler) httpGet(w http.ResponseWriter, req *http.Request) {
	value, found := handler.concurrentMap.Load(req.URL.Query().Get("key"))
	if found {
		fmt.Fprintf(w, "200 Found", value)
	} else {
		fmt.Fprintf(w, "404 Not Found")
	}
}

func (handler *ConcurrentMapHandler) httpDelete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if len(key) > 0 {
		_, found := handler.concurrentMap.Load(key)
		if found {
			handler.concurrentMap.Delete(key)
			fmt.Fprintf(w, "204 Deleted Key")	
		} else {
			fmt.Fprintf(w, "404 Not Found")
		}
	} else {
		fmt.Fprintf(w, "400 Bad Request")
	}
}

func main() {

	concurrentMap := sync.Map{}
	handler := &ConcurrentMapHandler{
		concurrentMap: &concurrentMap,
	}

	http.HandleFunc("/put", handler.httpPut)
	http.HandleFunc("/get", handler.httpGet)
	http.HandleFunc("/delete", handler.httpDelete)

	http.ListenAndServe(":3000", nil)
}