# 滑动窗口，单调队列

N=1000010
n,k=map(int, input().split())
s=list(map(int, input().split()))
q=[0]*N 

# 最小值
hh,tt=0,-1
for i in range(n):
    # 因为q[hh]记录的是区间的最大/最小值的位置，只有窗口超过了他指向位置，他才失效.
    if hh<=tt and i>k-1:
        hh+=1
    
    # 把大于当前元素的出队
    while hh<=tt and s[q[tt]]>s[i]:
        tt-=1

    # 当前元素入队
    tt+=1
    q[tt]=i

    if i>=k-1:
        print(s[q[hh]],end=' ')
print()

# 最大值
hh,tt=0,-1
for i in range(n):
    if hh<=tt and i-k+1>q[hh]:
        hh+=1
    
    while hh<=tt and s[q[tt]]<s[i]:
        tt-=1

    tt+=1
    q[tt]=i

    if i>=k-1:
        print(s[q[hh]],end=' ')

