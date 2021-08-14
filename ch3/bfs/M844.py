def bfs():
    q = deque()
    q.append([0, 0])
    dx, dy = [-1, 0, 1, 0], [0, 1, 0, -1]
    d[0][0] = 0

    while q:
        t = q.pop()

        for i in range(4):
            # 队头尝试往四个方向走
            x, y = t[0] + dx[i], t[1] + dy[i]

            # 如果在边界内，且位置上是0,且没有走过
            if x >= 0 and x < n and y >= 0 and y < m and g[x][y] == 0 and d[x][y] == -1:
                # 把当前这个坐标的距离+1,且入队
                d[x][y] = d[t[0]][t[1]] + 1
                q.appendleft([x, y])

    print(d[n - 1][m - 1])


from collections import deque

N = 110
g, d = [[0] * N for _ in range(N)], [[-1] * N for _ in range(N)]
n, m = map(int, input().split())

for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(m):
        g[i][j] = tmp[j]

bfs()
