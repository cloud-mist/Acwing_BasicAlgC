# 归并排序 模板

def merge_sort(s,l,r):
    # 只有1个或0个元素，不需要排
    if l>=r:
        return 

    mid = l+r>>1
    merge_sort(s, l, mid)
    merge_sort(s, mid+1, r)

    # ⭐合并两个有序序列,
    i,j=l,mid+1 #i,j指向两个序列的起点
    temp=[]

    while i<=mid and j <=r:
        if s[i] <= s[j]:
            temp.append(s[i])
            i+=1
        else:
            temp.append(s[j])
            j+=1
    if i<=mid:
        temp.extend(s[i:mid+1])
    if j<=r:
        temp.extend(s[j:r+1])
    # 结果复制回原数组, s是从l开始的。
    i,j = l,0 
    while i<=r:
        s[i] = temp[j]
        i+=1
        j+=1

n = int(input())
s = list(map(int, input().split()))
merge_sort(s, 0, n-1)
for i in s:
    print(i,end=' ')

