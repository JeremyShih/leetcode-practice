# 2021/1/17
import time
import math
from fractions import Fraction


class Solution:
    def countGoodRectangles(self, rectangles) -> int:
        maxLens = {}
        maxLen = 0
        for rec in rectangles:
            side = min(rec)
            maxLen = max(maxLen, side)
            if side in maxLens:
                maxLens[side] += 1
            else:
                maxLens[side] = 1

        return maxLens[maxLen]


start_time = time.time()
slt = Solution()

rectangles = [[5, 8], [3, 9], [5, 12], [16, 5]]
print(slt.countGoodRectangles(rectangles) == 3)

rectangles = [[2, 3], [3, 7], [4, 3], [3, 7]]
print(slt.countGoodRectangles(rectangles) == 3)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
