# 并查集
def find(x):
    if p[x]!=x:
        p[x] = find(p[x])
    return p[x]

n,m=map(int, input().split())
p=[i for i in range(n+10)]

while m:
    m-=1
    op,a,b=input().split()
    a,b=int(a),int(b)

    if op=='M':
        if find(a)!=find(b):
            p[find(a)]=find(b)
    else:
        if find(a)!=find(b):
            print('No')
        else:
            print('Yes')


    

