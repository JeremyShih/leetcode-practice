# 2021/1/10
import time
import math
from fractions import Fraction


class Solution:
    def decode(self, encoded, first: int):
        ans = [first]
        for enc in encoded:
            ans.append(enc ^ ans[-1])
        return ans


start_time = time.time()
slt = Solution()

encoded = [1, 2, 3]
first = 1
ans = [1, 0, 2, 1]
print(slt.decode(encoded, first) == ans)

encoded = [6, 2, 7, 3]
first = 4
ans = [4, 2, 0, 7, 4]
print(slt.decode(encoded, first) == ans)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
