#include <algorithm>
#include <cstdio>
#include <cstring>
using namespace std;
typedef unsigned long long ULL;

const int N = 2000010;
char s[N];
ULL hl[N], hr[N], p[N];
int P = 131;

ULL get(ULL h[], int l, int r) { return h[r] - h[l - 1] * p[r - l + 1]; }

int main() {
    int m = 1;
    // while 看,后面的语句的真假
    while (scanf("%s", s + 1), strcmp(s + 1, "END")) {
        int n = strlen(s + 1);

        // 预处理s,在两个字母间加同样的东西，目的：使偶数长度变为奇数长度，方便只做一种处理
        for (int i = 2 * n; i >= 0; i -= 2) {
            s[i] = s[i / 2];
            s[i - 1] = 'z' + 1;
        }
        n *= 2;
        // 求正序和逆序的Hash前缀和
        p[0] = 1;
        for (int i = 1, j = n; i <= n; i++, j--) {
            p[i] = p[i - 1] * P;
            hl[i] = hl[i - 1] * P + (s[i] - 'a' + 1);
            hr[i] = hr[i - 1] * P + (s[j] - 'a' + 1);
        }

        // 枚举中点，二分半径
        int res = 0;
        for (int i = 1; i <= n; i++) {
            int l = 0, r = min(i - 1, n - i);
            while (l < r) {
                int mid = l + r + 1 >> 1;
                if (get(hl, i - mid, i - 1) != get(hr, n - (i + mid) + 1, n - (i + 1) + 1))
                    r = mid - 1;
                else
                    l = mid;
            }

            if (s[i - l] <= 'z')
                res = max(res, l + 1);
            else
                res = max(res, l);
        }
        printf("Case %d: %d\n", m++, res);
    }

    return 0;
}
