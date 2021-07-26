# 某个面积区域内加c
def insert(x1,y1,x2,y2,c):
    b[x1][y1]     += c
    b[x1][y2+1]   -= c
    b[x2+1][y1]   -= c
    b[x2+1][y2+1] += c

# input
n,m,q=map(int, input().split())
a=[[0 for i in range(m+2)] for j in range(n+2)]
b=[[0 for i in range(m+2)] for j in range(n+2)]
for i in range(1,n+1):
    temp=list(map(int, input().split()))
    for j in range(1,m+1):
        a[i][j] = temp[j-1]
        
# b数组差分初始化
for i in range(1,n+1):
    for j in range(1,m+1):
        insert(i,j,i,j,a[i][j])
# op       
while q:
    q-=1
    x1,y1,x2,y2,c=map(int, input().split())
    insert(x1,y1,x2,y2,c)
    
# 求前缀和
for i in range(1,n+1):
    for j in range(1,m+1):
        b[i][j] += b[i-1][j] + b[i][j-1] -b[i-1][j-1]
        
# output
for i in range(1,n+1):
    for j in range(1,m+1):
        print(b[i][j],end=' ')
    print()
