# 2021/1/31
import time
import math
from fractions import Fraction


class Solution:
    def countBalls(self, lowLimit: int, highLimit: int) -> int:
        boxes = {}
        for i in range(lowLimit, highLimit+1):
            t = i
            s = t % 10
            while t >= 10:
                t /= 10
                t = math.floor(t)
                s += t % 10
            # print(s)
            if s not in boxes:
                boxes[s] = 0
            boxes[s] += 1

        m = float('-inf')
        for k in boxes:
            m = max(m, boxes[k])
        return m


start_time = time.time()
slt = Solution()

lowLimit = 1
highLimit = 10
print(slt.countBalls(lowLimit, highLimit) == 2)

lowLimit = 5
highLimit = 15
print(slt.countBalls(lowLimit, highLimit) == 2)

lowLimit = 19
highLimit = 28
print(slt.countBalls(lowLimit, highLimit) == 2)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
