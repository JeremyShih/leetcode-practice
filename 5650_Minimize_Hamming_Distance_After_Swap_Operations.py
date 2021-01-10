# 2021/1/10
import time
import math
from fractions import Fraction


class Solution:
    def minimumHammingDistance(self, source, target, allowedSwaps) -> int:
        if len(allowedSwaps) == 0:
            return self.distance(source, target)
        return self.distance(source, target)

    def distance(self, source, target):
        count = 0
        for i in range(len(source)):
            if source[i] != target[i]:
                count += 1
        return count


start_time = time.time()
slt = Solution()

source = [1, 2, 3, 4]
target = [2, 1, 4, 5]
allowedSwaps = [[0, 1], [2, 3]]
print(slt.minimumHammingDistance(source, target, allowedSwaps) == 1)

source = [1, 2, 3, 4]
target = [1, 3, 2, 4]
allowedSwaps = []
print(slt.minimumHammingDistance(source, target, allowedSwaps) == 2)

source = [5, 1, 2, 4, 3]
target = [1, 5, 4, 2, 3]
allowedSwaps = [[0, 4], [4, 2], [1, 3], [1, 4]]
print(slt.minimumHammingDistance(source, target, allowedSwaps) == 0)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
