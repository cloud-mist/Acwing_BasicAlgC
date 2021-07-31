# 开放定址法

# 返回的k：如果x在，k就是他的位置，如果x不在，k就是可以插入的位置
def find(x):
    k=x%N
    while h[k]!=null and h[k]!=x:
        k+=1
        if k==N: k=0 # 如果看完了最后一个坑，要循环看第一个坑
    return k


N=200003 # 取一个质数,为了减小冲突概率
null=1e9+10
h=[null]*N
n=int(input())

while n:
    n-=1
    op,x=input().split()
    k = find(int(x))

    if op=='I': 
        h[k] = int(x)
    else:
        if(h[k]!=null): print('Yes')
        else: print('No')


