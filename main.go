package main

import (
	"client"
	"flag"
	"generators/har"
	"har"
	"log"
	"net/http"
	"sync"
	"time"
)

type DefaultStats struct {
	start time.Time
}

func (stats DefaultStats) Send(req *http.Request, res *http.Response, time time.Duration) {
	log.Printf("response time: %f\n", time.Seconds())
}

func main() {
	var host = flag.String("host", "localhost", "Specify the endpoint for issuing requests.")
	var harFile = flag.String("har", "", "The HAR file to be used for playback.")
	var numClients = flag.Int("count", 1, "The number of clients that should run in parallel.")
	var duration = flag.Int("duration", 10, "Duration of test run (in seconds).")
	flag.Parse()

	stop := make(chan struct{})
	var done sync.WaitGroup

	harLog, err := har.FromFile(harFile)
	if err != nil {
		log.Panicf("Failed to read in the file: %s\n", *harFile)
	}

	for i := 0; i < *numClients; i++ {
		generator := hargenerator.Create(harLog, *host)
		client := client.Create(1, DefaultStats{}, generator, time.Since(time.Now()), stop, &done)
		done.Add(1)

		go client.Start()
	}

	time.Sleep(time.Duration(*duration) * time.Second)

	close(stop)
	done.Wait()
}
