def push(x):
    global tt
    tt+=1
    que[tt]=x

def pop():
    global hh
    hh+=1 

def empty():
    if hh<=tt:
        print('NO')
    else:
        print('YES')

def query():
    print(que[hh])


m=int(input())
N=100010
tt,hh=-1,0
que=[0]*N
while m:
    m-=1
    op,*s=input().split()
    if op=='push':
        push(s[0])
    elif op=='pop':
        pop()
    elif op=='empty':
        empty()
    else:
        query()


    

