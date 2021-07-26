# 在k的后面插入一个点
def add(k,x):
    global idx
    e[idx] = x

    r[idx] = r[k]
    l[idx] = k
    l[r[k]] = idx
    r[k] = idx

    idx+=1
    
# 删除k
def remove(k):
    l[r[k]] = l[k]
    r[l[k]] = r[k]

# input
N=100010
e,l,r=[0]*N,[0]*N,[0]*N
r[0],l[1]=1,0
idx=2
M=int(input())

# op
while M:
    M-=1

    op,*s=input().split()
    s=list(map(int, s))
    if len(s)==1: x=s[0]
    else: k,x=s[0],s[1]

    if op=='L':
        add(0, x)
    elif op=='R':
        add(l[1], x)
    elif op=='D':
        remove(x+1)
    elif op=='IL':
        add(l[k+1], x)
    else:
        add(k+1, x)
        
#output     
i=r[0]
while i!=1:
    print(e[i],end=' ')
    i=r[i]
