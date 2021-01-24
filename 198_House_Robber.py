# 2021/1/23
import time
import math
from fractions import Fraction


class Solution:
    def rob(self, nums) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) < 3:
            return max(nums)

        dp = nums[:2]
        dp.append(nums[0]+nums[2])
        for i in range(3, len(nums)):
            n = max(dp[i-2], dp[i-3])+nums[i]
            dp.append(n)

        return max(dp[-1], dp[-2])


start_time = time.time()
slt = Solution()

nums = [1, 2, 3, 1]
print(slt.rob(nums) == 4)

nums = [2, 7, 9, 3, 1]
print(slt.rob(nums) == 12)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
