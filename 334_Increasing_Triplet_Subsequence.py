# 2021/1/24
import time
import math
from fractions import Fraction


class Solution:
    def increasingTriplet(self, nums) -> bool:
        first, second = float('inf'), float('inf')
        for n in nums:
            if n > second:
                return True
            elif n < first:
                first = n
            elif n > first:
                second = n
        return False


start_time = time.time()
slt = Solution()

nums = [1, 2, 3, 4, 5]
print(slt.increasingTriplet(nums))

nums = [5, 4, 3, 2, 1]
print(not slt.increasingTriplet(nums))

nums = [2, 1, 5, 0, 4, 6]
print(slt.increasingTriplet(nums))

nums = [2, 1, 5, 0, 6]
print(slt.increasingTriplet(nums))

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
