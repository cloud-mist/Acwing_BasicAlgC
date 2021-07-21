# 数组元素的目标和
# i从前往后，j从后往前，如果ai+bj大于x,j就--

n,m,x=map(int, input().split())
a=list(map(int, input().split()))
b=list(map(int, input().split()))

j=m-1
for i in range(n):
    while j>=0 and a[i]+b[j]>x:
        j-=1
    if a[i]+b[j]==x:
        print(i,j,end=' ')
        break

