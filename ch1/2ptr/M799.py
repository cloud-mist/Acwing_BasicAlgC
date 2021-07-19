# 双指针
# 最长连续不重复子序列
n = int(input())
a = list(map(int, input().split()))
b=[0]*100010
ans=0

# b数组记录每个数出现的次数，如果i走到某个位置，这个数的次数>1的话，j就要往前移动，直到不再有重复,最后更新下答案
j=0
for i in range(n):
    b[a[i]]+=1
    while b[a[i]]>1:
        b[a[j]]-=1
        j+=1
    ans = max(ans, i-j+1)
print(ans)
