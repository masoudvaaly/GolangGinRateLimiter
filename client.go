package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	limit := 300000

	for i := 1; i < limit; i++ {
		resp, err := c.Get("http://localhost:8080/")
		if err != nil {
			fmt.Printf("Error %s", err)
			fmt.Println("Count is: ", i)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("\n%d Body : %s", i, body)
	}
}
