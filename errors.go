package main

import "fmt"

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
