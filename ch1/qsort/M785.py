# 快排模板题
def qsort(l,r,s):
    if l>=r:
        return 

    i,j = l-1 ,r+1 # 因为下面是先往中间走一次，所以先扩展一下边界,这样就是真正的边界。
    x = s[(l+r)//2] 

    while i < j:
        # python版 do while,whiletrue break
        while True:
            i +=1 
            if s[i] >= x:
                break
        while True:
            j-=1
            if s[j] <= x:
                break
        if i <j: # 这个条件不成立的话，就不需要交换了。
            s[i],s[j] = s[j],s[i]

            
n = int(input())
s = list(map(int, input().split()))
qsort(0, n-1, s)
for i in s:
    print(i,end=' ')

