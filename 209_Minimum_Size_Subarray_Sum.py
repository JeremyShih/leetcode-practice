class Solution:
	def minSubArrayLen(self, s, nums):
		left = right = 0
		csum = 0 #current sum
		minlen = 0
		while right <= len(nums):
			if csum < s:
				if right == len(nums):
					return minlen
				csum += nums[right]
				right += 1
			else:
				minlen = min(minlen, right-left) if minlen else right-left
				csum -= nums[left]
				left += 1

slt=Solution()
s=7
inp=[2,3,1,2,4,3]
print(slt.minSubArrayLen(s,inp)==2)
inp=[2,2,2]
print(slt.minSubArrayLen(s,inp)==0)
