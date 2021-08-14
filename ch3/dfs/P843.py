# n皇后
def dfs(u):
    if u == n:
        for i in range(n):
            for j in range(n):
                print(res[i][j], end="")
            print()
        print()

    for i in range(n):
        if not col[i] and not dg1[u + i] and not dg2[u - i + n]:
            res[u][i] = "Q"
            col[i] = dg1[u + i] = dg2[u - i + n] = True
            dfs(u + 1)
            col[i] = dg1[u + i] = dg2[u - i + n] = False
            res[u][i] = "."


n = int(input())
N = 20  # 至少取两倍的n的范围
res = [["."] * n for i in range(n)]
dg1, dg2, col = (
    [False] * N,
    [False] * N,
    [False] * N,
)

dfs(0)
