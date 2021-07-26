# 将x插到head后面
def add_to_head(x):
    global head,idx
    # global这行用途：为了说明用的全局变量，而不是函数内部的同名局部变量，否则会被程序报认为是局部变量，且没声明就使用了的错误
    ne[idx] = head  # 先把idx 指向 head所指向的位置
    head    = idx   # 再把head指向idx指向的位置
    e[idx]  = x     # idx位置储存x的值
    idx+=1          # idx指向下个位置

# 将x插到k的后面
def add(k,x):
    global idx

    ne[idx] = ne[k]
    ne[k]   = idx
    e[idx]  = x
    idx+=1

# 将k后面的点删掉
def remove(k):
    ne[k] = ne[ne[k]]

# -------------------------  main ----------------------------------
n=int(input())
N=100010
e,ne=[0]*N,[0]*N
head,idx=-1,0

# op
while n:
    n-=1

    op,*s=input().split()
    s=list(map(int, s))
     
    if op=='H':
        add_to_head(s[0])
    elif op=='I':
        add(s[0]-1, s[1])
    else:
        if op[1]=='0':
            head = ne[head]
        else:
            remove(s[0]-1)

# output：从head开始到-1结束,变化是下一个位置值
i=head
while i!=-1:
    print(e[i],end=' ')
    i = ne[i]
