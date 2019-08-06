package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker()
}

func ticker()  {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <- ticker.C:
			fmt.Println("hello katy")
		}
	}
}
