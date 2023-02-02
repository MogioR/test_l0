package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nats-io/stan.go"
)

const (
	cluster = "test-cluster"
	client  = "order-publisher"
	channel = "order-notification"
)

func main() {
	sc, err := stan.Connect(cluster, client)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Json patch: ")
		text, _ := reader.ReadString('\r')
		text = strings.Trim(text, "\r\n\"")

		file, err := os.ReadFile(text)
		if err != nil {
			fmt.Println("Read error")
			continue
		}

		err = sc.Publish(channel, file)
		if err != nil {
			fmt.Println("Publish error")
			continue
		}
	}
}
