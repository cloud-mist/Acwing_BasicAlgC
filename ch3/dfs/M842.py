# 全排列
def dfs(u):
    # 说明走到头了,该输出了
    if u==n: 
        for i in range(n):
            print(path[i],end=' ')
        print()

    for i in range(1,n+1):
        if not st[i]: # 如果这个点没有被用过
            path[u] = i 
            st[i] = True
            dfs(u+1)
            st[i] = False # 恢复现场

    
N=10
path,st=[0]*N,[False]*N
n=int(input())

dfs(0)
