N, M = 510, 10010
big_num = 0x3F3F3F3F
dist = [big_num] * N


def bf():
    # 初始化
    dist[1] = 0

    # 遍历k次，更新所有边
    for _ in range(k):
        bak = dist[:]  # 备份一下距离,防止被错误更新(copylist的方式还有 bak=list(dist))
        for j in range(m):
            a, b, w = edge[j]
            dist[b] = min(dist[b], bak[a] + w)

    # 因为有可能有负权，所以dist[n]可能是一个和bignum差不多大的数
    if dist[n] > big_num // 2:
        return "impossible"
    return dist[n]


n, m, k = map(int, input().split())
edge = []
for i in range(m):
    a, b, c = map(int, input().split())
    edge.append([a, b, c])

t = bf()
print(t)
