class Solution:
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        maxIndex=nums[0]
        for i in range(len(nums)):
        	if maxIndex<i:
        		return False
        	maxIndex=max(maxIndex,i+nums[i])
        return True


slt=Solution()
inp=[3,2,1,1,4]
print(slt.canJump(inp))
inp=[2,3,1,1,4,2]
print(slt.canJump(inp))
