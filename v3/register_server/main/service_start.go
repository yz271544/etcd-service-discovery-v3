package main

import (
	"fmt"
	dis "github.com/yz271544/etcd-service-discovery-v3/v3/register_server/discovery"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IP: "192.168.31.138"}

	s, err := dis.NewService(serviceName, serviceInfo, []string{
		//"http://192.168.1.17:2379",
		//"http://192.168.1.17:2479",
		//"http://192.168.1.17:2579",
		"http://127.0.0.1:2379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
