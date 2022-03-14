# math
import math


def is_prime(n):
    if n < 2:
        return False

    end = int(math.sqrt(n)) + 1
    for i in range(2, end):
        if n % i == 0:
            return False

    return True


n = int(input())

while n:
    n -= 1
    x = int(input())
    if is_prime(x):
        print("Yes")
    else:
        print("No")
