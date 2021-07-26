# 单调栈
n=int(input())
s=list(map(int, input().split()))
stk,tt = [0]*n,0
for i in range(n):
    # 当tt不等于0并且栈顶元素大于当前元素，栈顶元素出栈
    while tt and stk[tt]>=s[i]:
        tt-=1
    # 如果栈不为空，那么栈顶元素就是小于当前元素左边最近的元素
    if tt:
        print(stk[tt],end=' ')
    else:
        print(-1,end=' ')
    # 当前元素入栈
    tt+=1
    stk[tt]=s[i]

    
