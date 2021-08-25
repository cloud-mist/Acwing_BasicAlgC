import heapq

maxi = 0x3F3F3F3F
N = 150010
h, e, ne, st = (
    [-1] * N,
    [0] * N,
    [0] * N,
    [False] * N,
)
dist, w = [maxi] * N, [0] * N
idx = 0
heap = []


def add(a, b, c):
    global idx
    w[idx] = c
    e[idx] = b
    ne[idx] = h[a]
    h[a] = idx
    idx += 1


def dijkstra():
    # 初始化
    dist[1] = 0
    heapq.heappush(heap, (0, 1))

    # ----------------------------------------------
    # TODO: 重要部分  <24-08-21, cloud mist> #
    while heap:
        t = heapq.heappop(heap)  # 拿到距离最小的点, <距离，点>
        ver, distance = t[1], t[0]
        if st[ver]:
            continue
        st[ver] = True
        i = h[ver]
        while i != -1:
            j = e[i]
            if dist[j] > distance + w[i]:
                dist[j] = distance + w[i]
                heapq.heappush(heap, (dist[j], j))
            i = ne[i]
    # -----------------------------------------------
    if dist[n] == maxi:
        return -1
    return dist[n]


n, m = map(int, input().split())
while m:
    m -= 1
    a, b, c = map(int, input().split())
    add(a, b, c)
res = dijkstra()
print(res)
