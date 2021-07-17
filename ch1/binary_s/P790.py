# 数的3次方根

x = float(input())
l,r=-1e4,1e4
while l+1e-8 < r:
    mid = (l+r)/2
    if mid**3>=x: # >=x的话，结果肯定在左半边，所以右端点更新成mid
        r=mid
    else:
        l=mid
print('{0:.6f}'.format(l))
