# 字符串前缀hash
def get(l,r):
    # 乘到对齐次数，再减
    return (h[r]-h[l-1]*p[r-l+1]) % Q


n,m=map(int,input().split())
s=input()
N=100010
h,p=[0]*N,[0]*N
P,Q=131,2**64

# h数组存储string前缀的hash值
p[0]=1
for i in range(1,n+1):
    p[i]=p[i-1]*P % Q
    h[i]=(h[i-1]*P + ord(s[i-1])) % Q

while m:
    m-=1
    l1,r1,l2,r2=map(int, input().split())

    if get(l1, r1)==get(l2, r2):
        print('Yes')
    else:
        print('No')
