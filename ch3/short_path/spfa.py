import heapq

N, big_num = 100010, 0x3F3F3F3F
e, ne, w, h, idx = [0] * N, [0] * N, [0] * N, [-1] * N, 0
dist, st = [big_num] * N, [False] * N
heap = []


def spfa():
    dist[1] = 0
    heapq.heappush(heap, 1)
    st[1] = True

    while heap:
        t = heapq.heappop(heap)
        st[t] = False

        i = h[t]
        while i != -1:
            j = e[i]
            if dist[j] > dist[t] + w[i]:
                dist[j] = dist[t] + w[i]
                if not st[j]:
                    heapq.heappush(heap, j)
                    st[j] = True
            i = ne[i]

    if dist[n] == big_num:
        return "impossible"
    return dist[n]


def add(a, b, c):
    global idx
    e[idx] = b
    w[idx] = c
    ne[idx] = h[a]
    h[a] = idx
    idx += 1


n, m = map(int, input().split())
for _ in range(m):
    a, b, c = map(int, input().split())
    add(a, b, c)
res = spfa()
print(res)
