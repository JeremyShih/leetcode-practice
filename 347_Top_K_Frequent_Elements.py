# 2021/1/24
import time
import math
from collections import Counter
from typing import List
import heapq


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if k == len(nums):
            return nums

        count = Counter(nums)
        print(count)
        return heapq.nlargest(k, count.keys(), key=count.get)


start_time = time.time()
slt = Solution()

nums = [1, 1, 1, 2, 2, 3]
k = 2
ans = [1, 2]
print(slt.topKFrequent(nums, k) == ans)

nums = [1]
k = 1
ans = [1]
print(slt.topKFrequent(nums, k) == ans)

nums = [1, 2]
k = 2
ans = [1, 2]
print(slt.topKFrequent(nums, k) == ans)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
