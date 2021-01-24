# 2021/1/24
import time
import math
from fractions import Fraction


class Solution:
    def minCharacters(self, a: str, b: str) -> int:
        l = min(len(a), len(b))
        for i in range(l):
            print(ord(b[i])-ord(a[i]))
        return 0


start_time = time.time()
slt = Solution()

a = "aba"
b = "caa"
print(slt.minCharacters(a, b) == 2)

a = "dabadd"
b = "cda"
print(slt.minCharacters(a, b) == 3)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
