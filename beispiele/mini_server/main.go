package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var allCities = map[string]int{
	"Berlin":  123,
	"Hamburg": 345,
}

func main() {
	fmt.Println("Mini HTTP-Server")

	http.Handle("/", http.FileServer(http.Dir("./html")))

	http.HandleFunc("GET /cities/{name}", getCitiesName)
	http.HandleFunc("PUT /cities/{name}", putCitiesName)
	http.HandleFunc("GET /updates", getUpdates)

	http.ListenAndServe(":8080", nil)
}

func getCitiesName(w http.ResponseWriter, r *http.Request) {
	cityName := r.PathValue("name")

	if value, ok := allCities[cityName]; !ok {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(404)
		w.Write([]byte("Unable to find city '" + cityName + "'."))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte(strconv.Itoa(value)))
	}
}

func putCitiesName(w http.ResponseWriter, r *http.Request) {
	cityName := r.PathValue("name")

	lenStr := r.Header.Get("Content-Length")
	len, err := strconv.Atoi(lenStr)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Content-Length invalid or not given."))
		return
	}

	buffer := make([]byte, min(len, 1024))
	n, err := r.Body.Read(buffer[:])
	fmt.Println(n, err)

	body := string(buffer[:n])
	value, err := strconv.Atoi(body)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Unable to parse '" + body + "'."))
		return
	}

	allCities[cityName] = value
	w.WriteHeader(200)
}

func getUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	ticker := time.NewTicker(time.Second)

	for {
		var t time.Time
		select {
		case <-r.Context().Done():
			fmt.Println("connection closed")
			return

		case t = <-ticker.C:
		}

		timeStr := t.Format(time.RFC3339)
		_, err := w.Write([]byte("data: " + timeStr + "\n\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
		w.(http.Flusher).Flush()
	}
}
