package zk

import (
	"fmt"
)

var (
	ZKClient *SdClient
)

func init() {
	// zk uri
	servers := []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}
	var err error
	ZKClient, err = NewClient(servers, "/rmq", 10)
	if err != nil {
		panic(err)
	}

	//register service node
	node1 := &ServiceNode{"node1", "127.0.0.1", 5672}
	node2 := &ServiceNode{"node2", "127.0.0.1", 5673}
	node3 := &ServiceNode{"node3", "127.0.0.1", 5674}

	if err = ZKClient.Register(node1); err != nil {
		panic(err)
	}
	if err = ZKClient.Register(node2); err != nil {
		panic(err)
	}
	if err = ZKClient.Register(node3); err != nil {
		panic(err)
	}

	fmt.Println("zk init success")
}
