package distribute

import "strconv"

type NodeInfo struct {
	NodeId     int    `json:"nodeId"`
	NodeIpAddr string `json:"nodeIpAddr"`
	NodePort   string `json:"nodePort"`
}
type AddToClusterMessage struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}

func (node *NodeInfo) String() string {
	return "NodeInfo:{nodeId：" + strconv.Itoa(node.NodeId) +
		",nodeIpAddr:" + node.NodeIpAddr +
		",nodePort:" + node.NodePort + "}"
}
func (req *AddToClusterMessage) String() string {
	return "AddToClusterMessage:{Source:" + req.Source.String() +
		",dest:" + req.Dest.String() +
		",message:" + req.Message + "}"
}
