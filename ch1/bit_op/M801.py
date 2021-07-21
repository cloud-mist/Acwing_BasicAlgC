# 二进制中1的个数

# 每次找到lowbit(i)(说明找到了1个1),然后用原数减去这个数(就减去了一个1)，update答案加1,直到i为0结束
n=int(input())
s=list(map(int, input().split()))

for i in s:
    ans=0
    while i:
        i -= i&(-i) 
        ans+=1
    print(ans,end=' ')
