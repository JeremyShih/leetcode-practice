# 2018/11/5 第一題
class Solution:
	def sortColors(self, nums):
		if not nums: return []
		zero,one,two=0,0,0
		for n in nums:
			if n==0:
				zero+=1
			elif n==1:
				one+=1
			else:
				two+=1
		for i in range(len(nums)):
			if i<zero:
				nums[i]=0
			elif i<zero+one:
				nums[i]=1
			else:
				nums[i]=2

slt=Solution()
inp=[2,0,2,1,1,0]
ans=[0,0,1,1,2,2]
slt.sortColors(inp)
print(inp==ans)
