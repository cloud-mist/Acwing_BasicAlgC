# 二分模板题，数的范围

n,m=map(int, input().split())
s=list(map(int, input().split()))

while m:
    m-=1
    x=int(input())

    l,r=0,n-1
    while l<r:
        mid=l+r>>1
        if s[mid]>=x:
            r=mid
        else:
            l=mid+1

    if s[l]!=x:
        print("-1 -1")
    else:
        print(l,end=' ')
        l,r=0,n-1
        while l<r:
            mid=l+r+1>>1
            if s[mid]<=x:
                l=mid
            else:
                r=mid-1
        print(l)


        

