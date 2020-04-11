class Solution:
	def permute(self, nums):
		if len(nums)==1:
			return [nums]
		res=[]
		for i in range(len(nums)):
			for perm in self.permute(nums[:i]+nums[i+1:]):
				#print([nums[i]]+perm)
				res.append([nums[i]]+perm)
				#print(res)
		return res


slt=Solution()
inp=[1,2]
print(slt.permute(inp))
inp=[1,2,3]
ans=[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
print(slt.permute(inp))
