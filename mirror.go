package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func cleanIp(r *http.Request) string {
	addr := r.RemoteAddr
	occur := strings.LastIndex(addr, ":")
	return addr[:occur]
}

func mirrorStatus(w http.ResponseWriter, r *http.Request) {

	type StatusResponse struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
		Ip         string `json:"ip"`
		Method     string `json:"method"`
	}

	status := strings.TrimRight(r.URL.Path[len("/status/"):], "/")

	// try to see if it can be converted to integer
	status_code, err := strconv.Atoi(status)
	if err != nil {
		fmt.Fprintf(w, "Invalid status code: %s", status)
		return
	}

	message := http.StatusText(status_code)

	if len(message) == 0 {
		fmt.Fprintf(w, "Status code does not exist: %s", status)
		return
	}

	response := StatusResponse{
		message,
		status_code,
		cleanIp(r),
		r.Method,
	}

	js, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	w.Write(js)
}

func mirrorIp(w http.ResponseWriter, r *http.Request) {

	type IpResponse struct {
		Ip string `json:"ip"`
	}

	response := IpResponse{cleanIp(r)}

	js, err := json.Marshal(response)
	writeResponse(w, js, err)
}

func mirrorNow(w http.ResponseWriter, r *http.Request) {

	type NowResponse struct {
		RFC3339         string `json:"rfc3339"`
		ANSIC           string `json:"ansi_c"`
		UnixDate        string `json:"unix_date"`
		UnixSeconds     int64  `json:"unix_seconds"`
		UnixNanoSeconds int64  `json:"unix_nano_seconds"`
		Ip              string `json:"ip"`
	}

	now := time.Now()
	response := NowResponse{
		now.Format(time.RFC3339),
		now.Format(time.ANSIC),
		now.Format(time.UnixDate),
		now.Unix(),
		now.UnixNano(),
		cleanIp(r),
	}

	js, err := json.Marshal(response)
	writeResponse(w, js, err)

}

func mirrorUserAgent(w http.ResponseWriter, r *http.Request) {

	type UserAgentResponse struct {
		UserAgent string `json:"user_agent"`
		Ip        string `json:"ip"`
	}

	response := UserAgentResponse{r.UserAgent(), cleanIp(r)}

	js, err := json.Marshal(response)
	writeResponse(w, js, err)
}

func writeResponse(w http.ResponseWriter, js []byte, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func mirrorHeaders(w http.ResponseWriter, r *http.Request) {
	type HeaderResponse struct {
		Headers map[string]string `json:"headers"`
		Ip      string            `json:"ip"`
	}

	response := HeaderResponse{make(map[string]string), cleanIp(r)}

	for key, values := range r.Header {
		response.Headers[key] = values[0]
	}
	js, err := json.Marshal(response)
	writeResponse(w, js, err)
}

func index(w http.ResponseWriter, r *http.Request) {
	type IndexResponse struct {
		Apis [5]string `json:"apis"`
		Ip   string    `json:"ip"`
	}

	response := IndexResponse{
		[5]string{"/status/:code/", "/ip/", "/now/", "/user-agent/", "/headers/"},
		cleanIp(r),
	}
	js, err := json.Marshal(response)
	writeResponse(w, js, err)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/status/", mirrorStatus)
	http.HandleFunc("/ip/", mirrorIp)
	http.HandleFunc("/now/", mirrorNow)
	http.HandleFunc("/user-agent/", mirrorUserAgent)
	http.HandleFunc("/headers/", mirrorHeaders)

	log.Printf("Listening at http://localhost:8799/")
	log.Fatal(http.ListenAndServe(":8799", nil))
}
