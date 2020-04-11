# 2018/11/13 #?
class Solution:
	def countArrangement(self, N):
		numsList=list(range(1,N+1))
		# res=[]
		def dfs(cur,nums):
			if cur==N or len(nums)==0:
				# print(cur)
				return 1
			res=0
			c=cur+1
			for i in range(len(nums)):
				if nums[i]%c==0 or c%nums[i]==0:
					# print(nums[i],c)
					res+=dfs(c,nums[:i]+nums[i+1:])
					# print(res)
			return res
		ans=0
		for i in range(len(numsList)):
			ans+=dfs(1,numsList[:i]+numsList[i+1:])
		return ans

slt=Solution()
n=2
ans=2
print(slt.countArrangement(n)==ans)
n=3
ans=3
print(slt.countArrangement(n)==ans)
