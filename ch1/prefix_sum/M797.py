# 一维差分

n,m = map(int, input().split())
a=[0]*(n+10)
b=[0]*(n+10)
temp = list(map(int, input().split()))
for i in range(n):
    a[i+1] = temp[i]
    
def fd(l,r,c):
    b[l]+=c
    b[r+1]-=c

for i in range(1,n+1):
    fd(i, i, a[i])

while m:
    m-=1
    l,r,c=map(int, input().split())
    fd(l, r, c)

for i in range(1,n+1):
    b[i] += b[i-1]

for i in range(1,n+1):
    print(b[i], end=' ')

