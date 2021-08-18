#include <bits/stdc++.h>
using namespace std;
const int N = 100010;
int h[N], e[N], ne[N], idx;
int q[N], d[N];  // d是入度
int n, m;

void add(int a, int b) {

    e[idx] = b;
    ne[idx] = h[a];
    h[a] = idx++;
}

int topsort() {
    // 把入度为0的都加入队列
    int tt = -1, hh = 0;
    for (int i = 1; i <= n; i++) {
        if (!d[i]) {
            q[++tt] = i;
        }
    }

    // 当队列不为空，取出队头，遍历所有出边，然后让入度-1，当这些点的入度为0,就可以入队了。
    while (hh <= tt) {
        int t = q[hh++];
        for (int i = h[t]; i != -1; i = ne[i]) {
            int j = e[i];
            d[j]--;
            if (!d[j]) {
                q[++tt] = j;
            }
        }
    }
    // 如果tt等于n-1,说明存在一个拓扑序
    return tt == n - 1;
}

int main() {
    scanf("%d%d", &n, &m);
    memset(h, -1, sizeof h);
    int a, b;
    for (int i = 0; i < m; ++i) {
        scanf("%d%d", &a, &b);
        add(a, b);
        d[b]++;
    }

    if (topsort()) {
        for (int i = 0; i < n; i++) {
            printf("%d ", q[i]);
        }
    } else {
        puts("-1");
    }

    return 0;
}
