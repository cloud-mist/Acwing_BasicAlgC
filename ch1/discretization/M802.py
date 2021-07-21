n,m=map(int, input().split())

N=300010
a,s=[0]*N,[0]*N
add,query=[],[]
al=[]

def find(x):
    l,r=0,len(al)-1
    while l<r:
        mid = l+r>>1
        if al[mid] >=x:
            r=mid
        else:
            l=mid+1
    return l+1   # 映射从1开始，所以位置+1

for i in range(n):
    x,c=map(int, input().split()) # x:位置，c:要加的数
    add.append([x,c]) 
    al.append(x)

for i in range(m):
    l,r=map(int, input().split())
    query.append([l,r])
    al.append(l) 
    al.append(r)

# al数组，去重&排序
al = list(set(al))
al.sort()

# a数组，把加的数加完
for i in add:
    x = find(i[0])
    a[x] += i[1]

# s数组，前缀和存储
for i in range(1,len(al)+1):
    s[i] = s[i-1]+a[i]

for i in query:
    l,r=find(i[0]),find(i[1])
    print(s[r]-s[l-1])
