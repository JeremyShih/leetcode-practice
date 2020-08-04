# 2020/6/7
class Solution:
	def findNumbers(self, nums) -> int:
		ans=0
		for n in nums:
			if len(str(n)) % 2 == 0:
				ans+=1
		return ans

slt=Solution()
nums = [12,345,2,6,7896]
a=2
print(slt.findNumbers(nums)==a)
nums = [555,901,482,1771]
a=1
print(slt.findNumbers(nums)==a)
