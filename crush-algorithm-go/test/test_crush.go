package main

import "fmt"

// 测试 CRUSH 映射
func TestCrushMapping() {
	cluster := InitCluster()

	// 添加节点
	cluster.AddNode(1, "Node A")
	cluster.AddNode(2, "Node B")
	cluster.AddNode(3, "Node C")
	cluster.AddNode(4, "Node D")

	// 打印集群信息
	PrintCluster(cluster)

	// 测试对象映射
	objectName := "object123"
	mappedNodeID := cluster.CrushMap(objectName)
	fmt.Printf("Object '%s' is mapped to node %d\n", objectName, mappedNodeID)

	// 假设映射值应在 [0, 3] 之间
	if mappedNodeID >= 0 && mappedNodeID < cluster.NodeCount {
		fmt.Println("Test passed.")
	} else {
		fmt.Println("Test failed.")
	}
}

func main() {
	// 运行测试
	TestCrushMapping()
}
