# trie字符串统计


N=100010
son=[[0]*26 for i in range(N)]
idx=0
cnt=[0]*N

def insert(s):
    global idx # 记录最后一个新建结点的编号
    p=0
    for i in range(len(s)):
        u = ord(s[i]) - ord('a')
        # 如果没有找到已经存在的结点，就新建一个Q节点,u指向Q
        if son[p][u] == 0: 
            idx+=1
            son[p][u]=idx
        # p指向Q,继续下一个节点
        p=son[p][u]
    cnt[p]+=1 # 标记以此节点结尾的字符串数量

def query(s):
    p=0
    for i in range(len(s)):
        u = ord(s[i]) - ord('a')
        if son[p][u]==0:
            return 0
        p=son[p][u]
    return cnt[p]

T=int(input())
while T:
    T-=1
    op,s=input().split()
    if op=='I':
        insert(s)
    else:
       print(query(s))
