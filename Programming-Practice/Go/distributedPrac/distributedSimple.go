package main

import (
	"./distribute"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("test start")
	makeMasterOnError := flag.Bool("makeMaseterOnErr", false, "ip没有连接到集群，设为主节点")
	clusterip := flag.String("cluster ip", "127.0.0.1:8080", "任何节点都可以连接本节点")
	myport := flag.String("myport", "8001", "正在运行节点，端口8001")
	flag.Parse()
	fmt.Println(*makeMasterOnError) //输出
	fmt.Println(*clusterip)
	fmt.Println(myport)

	rand.Seed(time.Now().UnixNano())
	myid := rand.Intn(9999999)
	fmt.Println(myid)
	myip, _ := net.InterfaceAddrs()
	fmt.Println("myip:", myip[0])
	//创建当前节点信息
	me := distribute.NodeInfo{NodeId: myid, NodeIpAddr: myip[0].String(), NodePort: *myport}
	fmt.Println("当前节点：", me)

	dest := distribute.NodeInfo{NodeId: -1, NodeIpAddr: strings.Split(*clusterip, ":")[0], NodePort: strings.Split(*clusterip, ":")[1]}
	fmt.Println("目标节点：", dest)
	ableToConnect := distribute.ConnectToCluster(me, dest)

	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("本节点为初始节点")
		}
		distribute.ListenOnPort(me)
	} else {
		fmt.Println("异常退出")
	}
}
