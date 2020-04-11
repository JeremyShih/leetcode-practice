class Solution:
	def combinationSum(self, candidates, target):
		def dfs(t,res,cur,numList):
			if t==0:
				res.append(cur)
			for i in range(len(numList)):
				if t-numList[i]>=0:
					#print(t-numList[i],res,cur+[numList[i]],numList[i:])
					dfs(t-numList[i],res,cur+[numList[i]],numList[i:])
		res=[]
		dfs(target,res,[],candidates)
		return res

slt=Solution()
inp=[2,3,6,7]
t=7
ans=[
  [2,2,3],
  [7]
]
print(slt.combinationSum(inp,t)==ans)
inp=[2,3,5]
t=8
ans=[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
print(slt.combinationSum(inp,t)==ans)
