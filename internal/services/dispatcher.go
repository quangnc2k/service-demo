package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const someURL = "http://10.0.0.0/staff"

var DefaultDispatcher Dispatcher

type Dispatcher interface {
	DispatchDeleteStaff(id string) error
}

type HTTPDispatcher struct {
	client *http.Client
}

func (d *HTTPDispatcher) DispatchDeleteStaff(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(someURL, "/", id), nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf(someURL, " returned error code %d: %s", resp.StatusCode, string(bodyBytes))
		log.Println(err)
	}

	return nil
}

func NewHTTPDispatcher() Dispatcher {
	var httpDispatcher = new(HTTPDispatcher)
	httpDispatcher.client = http.DefaultClient
	return httpDispatcher
}

