class Solution:
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        ans=set()
        nums.sort()

sol=Solution()
inp=[1, 0, -1, 0, -2, 2]
#inp=[-2, 0, 1, 1, 2]
#inp=[-4, -4, -4, -4, -3, -3, -2, -1, 0, 2, 2, 2, 3, 3]
print(sorted(inp))
import time
start_time = time.time()
print(sol.threeSum(inp))
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)