package main

import (
	"fmt"
	dis "github.com/yz271544/etcd-service-discovery-v3/v3/register_server/discovery"
	"log"
	"time"
)

func main() {

	m, err := dis.NewMaster([]string{
		//"http://192.168.1.17:2379",
		//"http://192.168.1.17:2479",
		//"http://192.168.1.17:2579",
		"http://127.0.0.1:2379",
	}, "services/")

	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
