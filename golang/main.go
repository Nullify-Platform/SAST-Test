package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	host := os.Args[0]
	url := fmt.Sprintf("http://%s/", host)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d", res.StatusCode)
}
