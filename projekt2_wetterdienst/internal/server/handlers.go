package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	cities = sync.Map{}
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func getCitiesName(w http.ResponseWriter, r *http.Request) {
	cityName := r.PathValue("name")

	measurement, ok := cities.Load(cityName)
	if !ok {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(measurement)
	if err != nil {
		fmt.Println("ERROR: failed to marshall measurement:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}

func postCitiesName(w http.ResponseWriter, r *http.Request) {
	cityName := r.PathValue("name")

	body, err := io.ReadAll(r.Body)
	fmt.Println("POST", cityName, string(body))
	if err != nil {
		fmt.Println("ERROR: failed to read request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Evaluate input.

	var msg TempMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println("ERROR: failed to unmarshal request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cities.Store(cityName, msg)
	Post(cityName, msg)

	w.WriteHeader(http.StatusOK)
}

func getCitiesNameStream(w http.ResponseWriter, r *http.Request) {
	cityName := r.PathValue("name")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	msgChan := Listen(ctx, cityName)

	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)

	for done := false; !done; {
		var msg TempMessage
		select {
		case <-ctx.Done():
			done = true
			continue

		case msg = <-msgChan:
		}

		jsonMsg, _ := json.Marshal(msg)
		content := string(jsonMsg)

		_, err := w.Write([]byte("data: " + content + "\n\n"))
		if err != nil {
			fmt.Println("ERROR: failed to marshal data:", err)
			break
		}
		w.(http.Flusher).Flush()
	}
}
