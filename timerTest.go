package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println("Time: ", time.Now(), "i:", i)
		time.Sleep(time.Millisecond * 250)
		if i == 100 {
			return
		}
	}
}
