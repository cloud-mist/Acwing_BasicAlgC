#include <bits/stdc++.h>
using namespace std;
const int N = 100010;
int       h[N], e[N], ne[N], idx;
int       d[N], q[N];  // q是队列，d是距离
int       n, m;

void add(int a, int b) {
    e[idx]  = b;
    ne[idx] = h[a];
    h[a]    = idx++;
}

int bfs() {
    int tt = 0, hh = 0;
    memset(d, -1, sizeof d);
    q[0] = 1;
    d[1] = 0;

    // 当队列不为空，取出队头，遍历队头可以到达的各个点，当遍历的这个点没有访问过，那么把这个点的距离更新，且入队
    while (hh <= tt) {
        int t = q[hh++];
        for (int i = h[t]; i != -1; i = ne[i]) {
            int j = e[i];
            if (d[j] == -1) {
                d[j]    = d[t] + 1;
                q[++tt] = j;
            }
        }
    }
    return d[n];
}

int main() {
    scanf("%d%d", &n, &m);
    int a, b;
    memset(h, -1, sizeof h);
    for (int i = 0; i < m; i++) {
        scanf("%d%d", &a, &b);
        add(a, b);
    }

    printf("%d", bfs());

    return 0;
}
