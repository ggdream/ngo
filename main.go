package main

import "fmt"


func main() {
	if err := runCli(); err != nil {
		fmt.Println(err)
	}
}
