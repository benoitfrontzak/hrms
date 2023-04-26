package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// send concurrently multiple post requests and get responses back
func makeConcurrentPostRequests(requests []Request, responseChan chan *MyResponse) ([]*MyResponse, error) {
	// Make a POST request for each request using Goroutines
	for _, req := range requests {
		go func(r Request) {
			// Convert data to JSON
			jsonData, err := json.Marshal(r.Data)
			if err != nil {
				responseChan <- &MyResponse{URL: r.URL, Error: true, Message: err.Error()}
				return
			}

			// Make a POST request with the JSON data
			resp, err := http.Post(r.URL, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				responseChan <- &MyResponse{URL: r.URL, Error: true, Message: err.Error()}
				return
			}

			// Read the response body and decode it to answer struct
			defer resp.Body.Close()
			var answer jsonResponse
			dec := json.NewDecoder(resp.Body)
			err = dec.Decode(&answer)

			// body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				responseChan <- &MyResponse{URL: r.URL, Error: true, Message: err.Error()}
				return
			}

			// Send the response to the channel
			responseChan <- &MyResponse{URL: r.URL, Error: false, Message: answer.Message, Data: answer.Data, CreatedAt: answer.CreatedAt, CreatedBy: answer.CreatedBy}
		}(req)
	}

	// store all responses
	allResp := make([]*MyResponse, len(requests))

	// Wait for all responses to be received from the channel
	for i := 0; i < len(requests); i++ {
		resp := <-responseChan
		if resp.Error {
			// Handle the error
			return nil, errors.New(resp.Message)
		} else {
			// Process the response
			allResp[i] = resp
		}
	}

	return allResp, nil
}
