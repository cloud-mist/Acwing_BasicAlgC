// 朴素dijkstra
#include <bits/stdc++.h>
using namespace std;
const int N = 510;
int n, m;
int g[N][N], dist[N];
bool st[N];

int dijkstra() {
    // 初始化
    memset(dist, 0x3f, sizeof dist);
    dist[1] = 0;

    for (int i = 1; i <= n; i++) {
        int t = -1;
        // 找出未在dist中且距离最近的那个点
        for (int j = 1; j <= n; j++)
            if (!st[j] && (t == -1 || dist[t] > dist[j]))
                t = j;

        // 将其加入
        st[t] = true;

        // 更新所有点的距离
        for (int j = 1; j <= n; j++)
            dist[j] = min(dist[j], dist[t] + g[t][j]);
    }

    if (dist[n] == 0x3f3f3f3f)
        return -1;
    return dist[n];
}

int main() {
    scanf("%d%d", &n, &m);
    memset(g, 0x3f, sizeof g);
    while (m--) {
        int a, b, c;
        scanf("%d%d%d", &a, &b, &c);
        g[a][b] = min(g[a][b], c);  // 有重边的话，取最小的那个
    }

    int res = dijkstra();

    printf("%d", res);
    return 0;
}
