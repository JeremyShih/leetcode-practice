# 2020/11/29
import time
import math
from fractions import Fraction


class Solution:
    def maximumWealth(self, accounts) -> int:
        maxWealth = 0
        for acc in accounts:
            tmp = sum(acc)
            maxWealth = max(tmp, maxWealth)
        return maxWealth


start_time = time.time()
slt = Solution()

accounts = [[1, 2, 3], [3, 2, 1]]
print(slt.maximumWealth(accounts) == 6)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
