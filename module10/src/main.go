package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/powerfj/learn_git/module10/src/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func index(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v1.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	//w.Header().Set("pinkyHead", "pinkyHeadV")
	fmt.Printf("os version: %s \n", version)

	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	// clientip := getCurrentIP(r)
	clientip := ClientIP(r)

	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's OK")
}

/*
func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
*/

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func main() {
	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/images", images)
	mux.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}
