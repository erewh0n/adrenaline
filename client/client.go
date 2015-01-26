package client

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type Stats interface {
	Send(*http.Request, *http.Response, time.Duration)
}

type Generator interface {
	Generate(*http.Response) (*http.Request, error)
}

type Client struct {
	id        int
	http      *http.Client
	stats     Stats
	generator Generator
	rate      time.Duration
	stop      <-chan struct{}
	done      *sync.WaitGroup
}

func Create(id int, stats Stats, generator Generator, rate time.Duration, stop <-chan struct{}, done *sync.WaitGroup) *Client {
	return &Client{
		id:        id,
		stats:     stats,
		generator: generator,
		rate:      rate,
		stop:      stop,
		done:      done,
	}
}

func (client *Client) SetRate(rate time.Duration) {
	client.rate = rate
}

func (client *Client) Start() {

	httpClient := &http.Client{}
	var response *http.Response
	for {
		select {
		case <-client.stop:
			log.Printf("Client %d is terminating.", client.id)
			client.done.Done()
			return
		default:
			loopStartTime := time.Now()
			request, err := client.generator.Generate(response)
			if err != nil {
				log.Printf("Couldn't generate request: %s\n", err.Error())
				continue
			}

			requestTime := time.Now()
			response, err := httpClient.Do(request)
			responseTime := time.Since(requestTime)

			if err != nil {
				log.Printf("Couldn't send request: %s\n", err.Error())
				continue
			}
			response.Body.Close()

			loopTime := time.Since(loopStartTime)
			if loopTime < client.rate {
				time.Sleep(client.rate - loopTime)
			}

			client.stats.Send(request, response, responseTime)
		}
	}
}
