# 防线
T = int(input())
while T:
    T-=1

    s=[]
    n=int(input())
    l,r=0,0
    for i in range(n):
        ss,e,d=map(int, input().split())
        s.append([ss,e,d])
        r=max(r, e)

    def getsum(x):
        res=0
        for i in s:
            if i[0]<=x:
                res += (min(i[1], x) - i[0]) // i[2] + 1
        return res

    while l<r:
        mid = l+r>>1
        if getsum(mid)%2: #如果是奇数,答案肯定在左半边
            r=mid
        else:
            l=mid+1

    sum = getsum(r)-getsum(r-1)
    if sum%2:
        print(r,sum)
    else:
        print("There's no weakness.")
