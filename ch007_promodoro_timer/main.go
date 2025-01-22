package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one argument.")
		return
	}

	repeat, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	for i := 1; i <= repeat; i++ {
		msg := fmt.Sprintf("Work for 25minutes [%v/%v]", i, repeat)
		fmt.Println(msg)
		//time.Sleep(25 * time.Minute)
		time.Sleep(3 * time.Second)

		if i == repeat {
			fmt.Println("Finished!")
		} else {
			fmt.Println("Break for 5 minutes")
			//time.Sleep(5 * time.Minute)
			time.Sleep(1 * time.Second)
		}
	}
}
