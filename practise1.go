package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(now)
	rand.Seed(now)
	fmt.Println("my favorite number is", rand.Intn(10))
}
