#include <bits/stdc++.h>
using namespace std;
typedef unsigned long long ULL;
typedef long long LL;
const int N = 100010;

int e[N], ne[N];  // 节点的值和next指针
int head = -1, idx;

void add_to_head(int x) {
    e[idx] = x;
    ne[idx] = head;
    head = idx++;
}
void add(int k, int x) {
    e[idx] = x;
    ne[idx] = ne[k];
    ne[k] = idx++;
}
void remove(int k) { ne[k] = ne[ne[k]]; }

int main() {
    int m;
    scanf("%d", &m);
    char op[5];
    int k, x;

    while (m--) {
        scanf("%s %d", op, &k);

        if (!strcmp(op, "H")) {
            add_to_head(k);
        } else if (!strcmp(op, "D")) {
            if (!k)  // 如果k==0,就是删除 head指向的点
                head = ne[head];
            else
                remove(k - 1);
        } else {
            scanf("%d", &x);
            add(k - 1, x);
        }
    }

    for (int i = head; i != -1; i = ne[i])
        printf("%d ", e[i]);

    return 0;
}
