# 2021/1/31
import time
import math
from fractions import Fraction
from typing import List


class Solution:
    def canEat(self, candiesCount: List[int], queries: List[List[int]]) -> List[bool]:
        return []


start_time = time.time()
slt = Solution()

candiesCount = [7, 4, 5, 3, 8]
queries = [[0, 2, 2], [4, 2, 4], [2, 13, 1000000000]]
ans = [True, False, True]
print(slt.canEat(candiesCount, queries) == ans)

candiesCount = [5, 2, 6, 4, 1]
queries = [[3, 1, 2], [4, 10, 3], [3, 10, 100], [4, 100, 30], [1, 3, 1]]
ans = [False, True, True, False, False]
print(slt.canEat(candiesCount, queries) == ans)


elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
