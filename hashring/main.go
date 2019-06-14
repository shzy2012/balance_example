package main

import (
	"balance_example/hashring"
	"fmt"
)

func main() {
	// memcacheServers := []string{
	// 	"192.168.0.246:11212",
	// 	"192.168.0.247:11212",
	// 	"192.168.0.249:11212"}

	// ring := hashring.New(memcacheServers)
	// for i:=0;i<10;i++{
	// 	server, _ := ring.GetNode(fmt.Sprintf("key_%v",i))
	// 	fmt.Println(server)
	// }

	weights := make(map[string]int)
	weights["192.168.0.246:11212"] = 1
	weights["192.168.0.247:11212"] = 2
	weights["192.168.0.249:11212"] = 10

	ring := hashring.NewWithWeights(weights)

	for i := 0; i < 10; i++ {
		server, _ := ring.GetNode(fmt.Sprintf("key_%v", i))
		fmt.Println(server)
	}
}
