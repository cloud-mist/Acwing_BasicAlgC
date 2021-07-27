def find(x):
    if p[x]!=x:
        p[x]=find(p[x])
    return p[x]

n,m=map(int, input().split())
p=[i for i in range(n+10)]
size=[1]*(n+10)

while m:
    m-=1
    op,*s=input().split()
    a=int(s[0])

    if op=='C':
        b=int(s[1])
        if find(a)!=find(b):
            size[find(b)]+=size[find(a)]
            p[find(a)] = find(b)
    elif op=='Q1':
        b=int(s[1])
        if find(a)==find(b):
            print('Yes')
        else:
            print('No')
    else:
        print(size[find(a)])









