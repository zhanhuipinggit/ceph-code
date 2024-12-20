#include "cluster.h"

// 打印集群节点信息
void print_cluster(const Cluster* cluster) {
    printf("Cluster contains %d nodes:\n", cluster->node_count);
    for (int i = 0; i < cluster->node_count; i++) {
        printf("Node %d: %s\n", cluster->nodes[i].id, cluster->nodes[i].name);
    }
}
