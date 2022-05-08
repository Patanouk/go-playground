package _select

import (
	"errors"
	"net/http"
	"time"
)

var errorUrlTooSlow = errors.New("url is too slow")

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, 10*time.Second)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errorUrlTooSlow
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
