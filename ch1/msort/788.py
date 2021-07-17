# 逆序对的数量

# 直接暴力，会超时
# 做法：归并排序，3种情况加起来.1 构成逆序对的数在左边，2.在右边。3.一边一个。。
# 情况1和2 直接加。3推到公式，s{j} = mid - i +1
def msort(s,l,r):
    if l>=r:
        return 0

    mid = l+r>>1
    ans = msort(s,l, mid) + msort(s,mid+1, r)

    i,j=l,mid+1
    temp=[]
    while i<=mid and j <=r:
        if s[i]<=s[j]:
            temp.append(s[i])
            i+=1
        else:
            temp.append(s[j])
            j+=1
            ans += mid - i+1
    if i<=mid:
        temp.extend(s[i:mid+1])
    if j<=r:
        temp.extend(s[j:r+1])
    
    i,j=l,0
    while i<=r:
        s[i] = temp[j]
        i+=1
        j+=1

    return ans
    
n = int(input())
s=list(map(int, input().split()))
print(msort(s,0,n-1))
