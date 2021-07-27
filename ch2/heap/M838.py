# 堆排序

def down(x):
    # 求当前点和两个儿子（如果有的话）的最小值,t代表最小值
    t=x                 
    l,r=x*2,x*2+1
    if l<=size and h[l]<h[t]: t=l
    if r<=size and h[r]<h[t]: t=r
    
    # 如果最小不是x的话，就交换。然后递归处理
    if t!=x:
        h[t],h[x]=h[x],h[t]
        down(t)

def up(x):
    # 如果父节点存在且父节点较大，就交换，然后层层往上
    while x//2 and h[x//2]>h[x]:
        h[x//2],h[x]=h[x],h[x//2]
        x//=2

n,m=map(int, input().split())
h,size=[0],n
h.extend(map(int, input().split()))

# 建堆
for i in range(n//2,0,-1):
    down(i)

while m:
    m-=1
    print(h[1],end=' ')
    h[1]=h[size]
    size-=1
    down(1)
