# 第k小数
# v1 : 直接使用快排 nlogn
def qsort(s,l,r):
    if l>=r:
        return
    
    x=s[l+r>>1]
    i,j=l-1,r+1
    while i<j:
        while True:
            i+=1
            if s[i]>=x:
                break
        while True:
            j-=1
            if s[j]<=x:
                break
        if i<j: #重点
            s[i],s[j] = s[j],s[i]
    
    qsort(s, l, j)
    qsort(s, j+1, r)

n,k=map(int, input().split())
s = list(map(int, input().split()))
qsort(s, 0, n-1)
print(s[k-1])


