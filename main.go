package main

import (
	"fmt"
	"github.com/kabbali/go-examples/http_calls"
)

func main() {
	endpoints, err := http_calls.GetEndpoints()

	fmt.Println(err)
	fmt.Println(endpoints.EventsUrl)
}
