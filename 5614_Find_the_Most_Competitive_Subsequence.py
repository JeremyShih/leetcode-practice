# 2020/11/29
import time
import math
from fractions import Fraction


class Solution:
    def mostCompetitive(self, nums, k: int):
        ans, cur = [], 0
        import numpy
        # print(nums)
        sortIndex = list(numpy.argsort(nums))
        # print(sortIndex[1:3])

        while len(ans) < k:
            minIndex = cur
            print(minIndex, sortIndex.index(minIndex))
            while sortIndex.index(minIndex) < cur or sortIndex.index(minIndex) > len(nums)+len(ans)-k+1:
                minIndex += 1
            # for i in range(cur, len(nums)+len(ans)-k+1):
            #     if nums[i] < minNum:
            #         minIndex, minNum = i, nums[i]
            minNum = nums[minIndex]

            ans += [minNum]
            cur = minIndex+1

        return ans


start_time = time.time()
slt = Solution()

nums, k = [3, 5, 2, 6],  2
ans = [2, 6]
print(slt.mostCompetitive(nums, k))
# print(slt.mostCompetitive(nums, k) == ans)
nums, k = [2, 4, 3, 3, 5, 4, 9, 6],  4
ans = [2, 3, 3, 4]
print(slt.mostCompetitive(nums, k) == ans)
nums, k = [71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2], 3
ans = [8, 80, 2]
print(slt.mostCompetitive(nums, k) == ans)


elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
