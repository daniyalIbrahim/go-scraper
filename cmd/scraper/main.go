package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"scraper/api"
	"scraper/helpers"
	"time"
)

func main() {
	fmt.Println("Initializing Rest Server " + time.Now().String())
	helpers.GetCPUInformation()
	srv := &http.Server{
		Addr: ":8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.GetMuxAPI(), // Pass our instance of go-chi/chi in.
	}
	// Run our server in a goroutine so that it doesn't block.
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error while running server %v", err)
	}
}

func GetCPUInformation() {
	log.Printf("Getting CPU Information")
	//Get MAX CPU CORES
	log.Printf("Max CPU Cores: %v", runtime.NumCPU())
	//Get MAX CPU FREQUENCY
	log.Printf("Runtime GOARCH: %v", runtime.GOARCH)
	//Get MAX CPU THREADS
	log.Printf("Max CPU Threads: %v", runtime.GOMAXPROCS(0))
}
