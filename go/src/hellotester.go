package main

import (
	"fmt"
	"helloworld"
)

func main() {
	fmt.Println("hello tester")
	world := helloworld.NewHelloWorld()
	fmt.Println(world.GetHelloWorld())
}
