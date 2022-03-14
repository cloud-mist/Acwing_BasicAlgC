N = 510
w = [[0 for _ in range(N)] for _ in range(N)]
f = [[0 for _ in range(N)] for _ in range(N)]
n = int(input())

for i in range(1, n + 1):
    tmp = list(map(int, input().split()))
    for j in range(1, i + 1):
        w[i][j] = tmp[j - 1]

for i in range(1, n + 1):
    f[n][i] = w[n][i]

for i in range(n - 1, 0, -1):
    for j in range(1, i + 1):
        f[i][j] = max(f[i + 1][j] + w[i][j], f[i + 1][j + 1] + w[i][j])

print(f[1][1])
