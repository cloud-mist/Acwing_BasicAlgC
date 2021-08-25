N = 510
maxi = 0x3F3F3F3F
g, dist, st = [[maxi] * N for _ in range(N)], [maxi] * N, [False] * N


def dijkstra():
    dist[1] = 0
    for _ in range(1, n + 1):
        t = -1
        for j in range(1, n + 1):
            if not st[j] and (t == -1 or dist[t] > dist[j]):
                t = j
        st[t] = True
        for j in range(1, n + 1):
            dist[j] = min(dist[j], dist[t] + g[t][j])

    if dist[n] == maxi:
        return -1
    return dist[n]


n, m = map(int, input().split())
while m:
    m -= 1
    a, b, c = map(int, input().split())
    g[a][b] = min(g[a][b], c)
res = dijkstra()
print(res)
