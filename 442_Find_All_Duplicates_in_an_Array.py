# 2018/11/10 #7
class Solution:
	def findDuplicates(self, nums):
		res=[]
		for n in nums:
			ind=abs(n)-1
			if nums[ind]<0:
				res+=[abs(n)]
			nums[ind]=-nums[ind]
		return res


slt=Solution()
inp=[4,3,2,7,8,2,3,1]
ans=[2,3]
print(slt.findDuplicates(inp))
inp=[10,2,5,10,9,1,1,4,3,7]
ans=[10,1]
print(slt.findDuplicates(inp))
