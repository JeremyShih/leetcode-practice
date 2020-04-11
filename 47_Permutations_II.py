class Solution:
	def permuteUnique(self, nums):
		if len(nums)==1:
			return [nums]
		res=[]
		nums.sort()
		for i in range(len(nums)):
			if i>0 and nums[i-1]==nums[i]:
				continue
			for perm in self.permuteUnique(nums[:i]+nums[i+1:]):
				#print([nums[i]]+perm)
				res.append([nums[i]]+perm)
				#print(res)
		return res


slt=Solution()
inp=[1,2]
print(slt.permuteUnique(inp))
inp=[1,2,1]
ans=[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
print(slt.permuteUnique(inp))
