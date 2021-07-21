# input
n=int(input())
s=[]
while n:
    n-=1
    s.append(list(map(int, input().split())))
    
def merge(s):
    s.sort()
    ans=0               # 存储结果
    st,ed=-2e9,-2e9

    # 测试过，如果区间为空，不会执行for
    for i in s:
        if ed < i[0]:   # 如果是无交集区间
            ans+=1
            st,ed = i[0],i[1]
        else:
            ed = max(ed, i[1]) # 如果有交集，看一下哪个ed最长，取哪个
    return ans
    
# output    
print(merge(s))
