#include <bits/stdc++.h>
using namespace std;
const int N = 100010, M = 2 * N;
int h[N], e[M], ne[M], idx;  // h是链表的所有head
bool st[N];                  // 记录当前点是否被访问过
int ans = N, n;

// 链表头插法
void add(int a, int b) {
    e[idx] = b;
    ne[idx] = h[a];
    h[a] = idx++;
}

int dfs(int x) {
    st[x] = true;  // 标记此点已被搜过
    int sum = 1, res = 0;
    // res记录当前点分割的连通max

    for (int i = h[x]; i != -1; i = ne[i]) {
        // 遍历此点的所有出边
        int j = e[i];

        if (!st[j]) {
            int s = dfs(j);     // s:当前点子树的值
            res = max(res, s);  // 取一个max
            sum += s;           // sum 当前点为根的值
        }
    }
    res = max(res, n - sum);
    ans = min(ans, res);
    return sum;
}

int main() {
    // 输入
    scanf("%d", &n);
    memset(h, -1, sizeof h);  // head都指向-1
    int a, b;
    for (int i = 0; i < n - 1; i++) {
        scanf("%d%d", &a, &b);
        add(a, b);
        add(b, a);
    }

    dfs(1);

    printf("%d\n", ans);

    return 0;
}
