class Solution:
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 1:
            return nums[0]
        max = 0
        for i in range(1, len(nums) + 1):
            for j in range(len(nums) - i + 1):
                sum = 0
                for k in range(j, j + i):
                    sum = sum + nums[k]
                if sum > max:
                    max = sum
        return max
