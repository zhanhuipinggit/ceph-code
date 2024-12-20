#include "crush.h"

// 初始化集群
void init_cluster(Cluster* cluster) {
    cluster->node_count = 0;
}

// 向集群中添加节点
void add_node(Cluster* cluster, int id, const char* name) {
    if (cluster->node_count < MAX_NODES) {
        cluster->nodes[cluster->node_count].id = id;
        strncpy(cluster->nodes[cluster->node_count].name, name, sizeof(cluster->nodes[cluster->node_count].name) - 1);
        cluster->node_count++;
    }
}

// 简单的哈希函数，用于根据对象名称进行映射
int hash(const char* str) {
    int hash = 0;
    while (*str) {
        hash = (hash * 31) + *str++;
    }
    return hash;
}

// 使用 CRUSH 算法映射数据对象到集群节点
int crush_map(Cluster* cluster, const char* object_name) {
    int hashed_value = hash(object_name);
    return hashed_value % cluster->node_count;
}
