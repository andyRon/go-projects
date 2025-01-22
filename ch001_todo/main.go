package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	type TODOs struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}
	counter := 0
	data := make([]TODOs, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("CMDs: show create remove done\n>")

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {
		case "show":
			for _, todo := range data {
				fmt.Println("-------")
				statusLine := "[ ]"
				if todo.Status {
					statusLine = "[x]"
				}
				fmt.Println(statusLine, todo.Description, " ID:", todo.Id)
			}
		case "create":
			data = append(data, TODOs{
				Id:          counter,
				Description: strings.Join(split[1:], " "),
				Status:      false,
			})
			counter++
			fmt.Println("Created TODO Item")
		case "remove":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range data {
				if todo.Id == index {
					data = slices.Delete(data, i, i+1)
					break
				}
			}
		case "done":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range data {
				if todo.Id == index {
					data[i].Status = true
				}
			}
		}
		fmt.Print(">")
	}
}
