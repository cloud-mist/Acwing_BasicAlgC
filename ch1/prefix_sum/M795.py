# 一维前缀和

n,m=map(int, input().split())
s,temp=[0],[0]
s.extend(list(map(int, input().split())))

for i in range(1,n+1):
    temp.append(temp[i-1]+s[i])

while m:
    m-=1
    l,r=map(int, input().split())
    print(temp[r]-temp[l-1])
        


