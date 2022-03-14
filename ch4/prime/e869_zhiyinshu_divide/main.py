import math


def divide(n):
    end = math.sqrt(n) + 1

    for i in range(2, int(end)):
        if n % i == 0:
            s = 0
            while n % i == 0:
                n //= i
                s += 1

            print(i, s)

    if n > 1:
        print(n, 1)


n = int(input())

while n != 0:
    n -= 1

    x = int(input())
    divide(x)
    print()
