# 插入
def push(x):
    global tt
    tt+=1
    stk[tt]=x

# 查询
def query():
    print(stk[tt])

# 弹出
def pop():
    global tt
    tt-=1

# 是否为空
def empty():
    if tt<=0:
        print('YES')
    else:
        print('NO')


# input
N=100010
stk,tt=[0]*N,0
M=int(input())

# op
while M:
    M-=1
    op,*s=input().split()

    if op=='push':
        push(s[0])
    elif op=='query':
        query()
    elif op=='empty':
        empty()
    else:
        pop()

