package distribute

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func ConnectToCluster(me, dest NodeInfo) bool {
	conn, err := net.DialTimeout("tcp", dest.NodeIpAddr+":"+dest.NodePort, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("没有连接到集群", me.NodeId)
			return false
		}
	} else {
		fmt.Println("成功连接到集群")
		text := "与节点（" + strconv.Itoa(me.NodeId) + ")进行连接"
		//构造请求
		req := GetAddToClusterMessage(me, dest, text)
		//
		json.NewEncoder(conn).Encode(&req)
		var resp AddToClusterMessage
		//
		json.NewDecoder(conn).Decode(&resp)
		fmt.Println("返回信息：", resp.String())
		return true
	}

	return false
}

func GetAddToClusterMessage(source, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			NodePort:   source.NodePort,
		},
		Dest: NodeInfo{
			NodeId:     dest.NodeId,
			NodeIpAddr: dest.NodeIpAddr,
			NodePort:   dest.NodePort,
		},
		Message: message,
	}
}
func ListenOnPort(me NodeInfo) {
	In, _ := net.Listen("tcp", ":"+me.NodePort)
	for {
		con, err := In.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("接受信息出错", me.NodeId)
			}
		} else {
			var req AddToClusterMessage
			//接收请求
			json.NewDecoder(con).Decode(&req)
			fmt.Println("接收到：", req.String())
			text := "回复"
			resp := GetAddToClusterMessage(me, req.Source, text)
			json.NewEncoder(con).Encode(&resp)
			con.Close()

		}
	}
}
