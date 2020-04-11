class Solution:
	def findTargetSumWays(self, nums, S):
		L = len(nums)
		sumrecord = [0]*(L+1)
		for i in range(1,L+1):
			sumrecord[i] = sumrecord[i-1]+nums[i-1]
		sumrecord = sumrecord[1:]
		memo = {(0,-1):1}
		def dfs(target,end):
			if (target,end) in memo:
				return memo[(target,end)]
			if end==-1 or target>sumrecord[end] or target<-sumrecord[end]:
				return 0
			#用dfs的意思應該是分別考慮，數列中第n個數前面擺+/-的時候
			#，搭配前面n-1個數組合結果等於target減第n個數或加第n個數的狀況
			memo[(target, end)] = dfs(target+nums[end],end-1)+dfs(target-nums[end],end-1)
			return memo[(target,end)]
		return dfs(S,L-1)

slt=Solution()
inp=[1, 1, 1, 1, 1]
s=3
print(slt.findTargetSumWays(inp,s)==5)
