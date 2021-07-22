n,f=map(int, input().split())
cow=[0]
sum=[0]*(n+1)
for i in range(1,n+1):
    cow.append(int(input()))

def check(mid):
    # 每个前缀和-mid
    for i in range(1,n+1):
        sum[i] = sum[i-1]+cow[i]-mid

    minv=0 # 因为sum[0]是0
    i,j=0,f
    while j<=n:
        minv = min(minv, sum[i]) #记录之前的最小值
        if sum[j]-minv>=0: # 说明有合法方案
            return True
        j+=1;i+=1
    return False

l,r=0,2000
while l+1e-5<r:
    mid=(l+r)/2
    if check(mid): # 说明存在平均值不小于二分的值，所以要取右半边
        l=mid
    else:
        r=mid
print(int(r*1000))


