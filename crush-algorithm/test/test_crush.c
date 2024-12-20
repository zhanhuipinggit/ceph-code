#include <assert.h>
#include <stdio.h>
#include "crush.h"
#include "cluster.h"

// 测试 CRUSH 算法
void test_crush_mapping() {
    Cluster cluster;
    init_cluster(&cluster);

    // 添加节点
    add_node(&cluster, 1, "Node A");
    add_node(&cluster, 2, "Node B");
    add_node(&cluster, 3, "Node C");
    add_node(&cluster, 4, "Node D");

    // 打印集群信息
    print_cluster(&cluster);

    // 测试对象映射
    const char* object_name = "object123";
    int mapped_node_id = crush_map(&cluster, object_name);
    printf("Object '%s' is mapped to node %d\n", object_name, mapped_node_id);

    // 假设映射值应在 [0, 3] 之间
    assert(mapped_node_id >= 0 && mapped_node_id < cluster.node_count);
}

int main() {
    // 运行测试
    test_crush_mapping();
    printf("All tests passed.\n");
    return 0;
}
