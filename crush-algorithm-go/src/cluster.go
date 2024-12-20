package main

import "fmt"

// 打印集群节点信息
func PrintCluster(cluster *Cluster) {
	fmt.Printf("Cluster contains %d nodes:\n", cluster.NodeCount)
	for _, node := range cluster.Nodes {
		fmt.Printf("Node %d: %s\n", node.ID, node.Name)
	}
}
