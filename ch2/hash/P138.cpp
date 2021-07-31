#include <cstdio>
#include <cstring>

typedef unsigned long long ULL;
const int N=1000010, P=131;
ULL p[N], h[N];

ULL get(int l,int r){
	return h[r]-h[l-1]*p[r-l+1];
}

int main(){
	char s[N];
	int m;
	scanf("%s%d", s,&m);
	int n=strlen(s);

	p[0]=1;
	for (int i = 1; i <=n; ++i) {
		p[i] = p[i-1]*P;
		h[i] = h[i-1]*P+s[i-1];
	}

	int l1,r1,l2,r2;
	while (m--){
		scanf("%d%d%d%d",&l1,&r1,&l2,&r2);
		if (get(l1, r1)==get(l2, r2)) {
			printf("Yes\n");
		}else{
			printf("No\n");
		}
	}
	return 0;
}
