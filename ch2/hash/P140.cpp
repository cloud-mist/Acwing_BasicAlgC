#include <bits/stdc++.h>
using namespace std;
typedef unsigned long long ULL;
typedef long long LL;
const int N = 100010;
char s[N];
ULL h[N], p[N];
int P = 131, sa[N];

ULL get(int l, int r) { return h[r] - h[l - 1] * p[r - l + 1]; }

int get_prefix(int a, int b) {}

bool cmp(int a, int b) {}

int main() {
    scanf("%s", s + 1);
    int n = strlen(s) + 1;

    // 前缀
    p[0] = 1;
    for (int i = 1; i <= n; i++) {
        p[i] = p[i - 1] * P;
        h[i] = h[i - 1] * P + s[i] - 'a' + 1;
        sa[i] = i;
    }
    sort(sa, sa + n, cmp);

    return 0;
}
