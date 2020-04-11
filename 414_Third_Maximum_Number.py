# 2018/11/9 #6
class Solution:
	def thirdMax(self, nums):
		rank=[-10**10]*3
		for n in nums:
			if rank[2]<n and n<rank[1]:
				rank[2]=n
			elif rank[1]<n and n<rank[0]:
				rank[2],rank[1]=rank[1],n
			elif rank[0]<n:
				rank[2],rank[1],rank[0]=rank[1],rank[0],n
		# print(rank)
		return rank[-1] if rank[-1]>-10**10 else rank[0]


slt=Solution()
inp=[3, 2, 1]
ans=1
print(slt.thirdMax(inp)==ans)
inp=[1, 2]
ans=2
print(slt.thirdMax(inp)==ans)
inp=[2, 2, 3, 1]
ans=1
print(slt.thirdMax(inp)==ans)
inp=[1,1,2]
ans=2
print(slt.thirdMax(inp)==ans)
