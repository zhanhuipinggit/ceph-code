package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

const MaxNodes = 1000

// Node 结构体代表一个集群节点
type Node struct {
	ID   int
	Name string
}

// Cluster 结构体表示整个集群
type Cluster struct {
	Nodes     []Node
	NodeCount int
}

// 初始化集群
func InitCluster() *Cluster {
	return &Cluster{
		Nodes:     make([]Node, 0, MaxNodes),
		NodeCount: 0,
	}
}

// 向集群添加节点
func (cluster *Cluster) AddNode(id int, name string) {
	if cluster.NodeCount < MaxNodes {
		node := Node{ID: id, Name: name}
		cluster.Nodes = append(cluster.Nodes, node)
		cluster.NodeCount++
	}
}

// 简单的哈希函数（FNV-1a）
func hash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

// CRUSH 映射函数，将数据映射到集群节点
func (cluster *Cluster) CrushMap(objectName string) int {
	hashedValue := hash(objectName)
	return int(hashedValue) % cluster.NodeCount
}
