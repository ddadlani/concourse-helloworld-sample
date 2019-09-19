package main

import (
	"fmt"
	"os"
)

func SetTarget(target string) {
	os.Setenv("TARGET", target)
}

func main() {
	SetTarget("CF Onboarding Week September")
	fmt.Println("Hello World " + os.Getenv("TARGET"))
}
