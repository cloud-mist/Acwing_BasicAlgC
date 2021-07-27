# 模拟堆
#ph 第k个点在堆里的位置
#hp 堆里的点是第几个点

def h_swap(a,b):
    ph[hp[a]],ph[hp[b]]=ph[hp[b]],ph[hp[a]]
    hp[a],hp[b]=hp[b],hp[a]
    h[a],h[b]= h[b],h[a]

def down(x):
    # 求当前点和两个儿子（如果有的话）的最小值,t代表最小值
    t=x                 
    l,r=x*2,x*2+1
    if l<=size and h[l]<h[t]: t=l
    if r<=size and h[r]<h[t]: t=r
    
    # 如果最小不是x的话，就交换。然后递归处理
    if t!=x:
        h_swap(t, x)
        down(t)

def up(x):
    # 如果父节点存在且父节点较大，就交换，然后层层往上
    while x//2 and h[x//2]>h[x]:
        h_swap(x//2, x)
        x//=2

N=100010
h,ph,hp=[0]*N,[0]*N,[0]*N
size,m=0,0  # m记录第几个插入的数
T=int(input())

while T:
    T-=1
    op,*s=input().split()
    
    if op=='I':
        x=int(s[0])
        size+=1;m+=1

        ph[m]=size
        hp[size]=m
        h[size]=x

        up(size)
    elif op=='PM':
        print(h[1])
    elif op=='DM':
        h_swap(1, size)
        size-=1
        down(1)
    elif op=='D':
        k=int(s[0])
        k=ph[k] # 堆中位置
        h_swap(k, size)
        size-=1
        down(k);up(k)
    
    else: # 第k个插入的数修改成x
        k,x=int(s[0]),int(s[1])
        k=ph[k]
        h[k]=x
        down(k);up(k)
         
