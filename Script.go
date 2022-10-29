package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var url string
	var period int
	fmt.Print("Enter the URL of a website: ")
	fmt.Scan(&url)
	fmt.Print("Enter the period of checking service availability (seconds): ")
	fmt.Scan(&period)
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print(err, "\n")
			main()
		}
		status_code := resp.StatusCode
		fmt.Println(status_code, ":", url)
		time.Sleep(time.Duration(period) * time.Second)
	}
}
