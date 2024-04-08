package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) (string, error) {
	return ConfigurableRacer(slowUrl, fastUrl, 10*time.Second)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ConfigurableRacer(slowUrl, fastUrl string, timeout time.Duration) (string, error) {
	select {
	case <-ping(slowUrl):
		return slowUrl, nil
	case <-ping(fastUrl):
		return fastUrl, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", slowUrl, fastUrl)
	}
}
