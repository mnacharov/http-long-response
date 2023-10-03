package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "{\"timeout\": null,\"status\": \"FAIL\"}")
		if err != nil {
			log.Printf("Failed to write http response: %v\n", err)
		}
		return
	}
	time.Sleep(time.Duration(timeout) * time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"timeout\": %s,\"status\": \"OK\"}")
	if err != nil {
		log.Printf("Failed to write http response: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
