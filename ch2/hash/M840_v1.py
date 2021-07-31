# 散列表, 拉链法
def insert(x):
    global idx
    k=x%N

    # h[k]是head
    e[idx]  = x
    ne[idx] = h[k]
    h[k]    = idx
    idx+=1


def find(x):
    k=x%N

    i=h[k]
    while i!=-1:
        if e[i]==x:
            return True
        i=ne[i]
    return False


N=100003 # 取一个质数,为了减小冲突概率
h=[-1]*N
e,ne,idx=[0]*N,[0]*N,0
n=int(input())

while n:
    n-=1
    op,num=input().split()
    x=int(num)

    if op=='I':
        insert(x)
    else:
        if(find(x)): print('Yes')
        else: print('No')


