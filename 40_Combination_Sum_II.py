# 2018/11/8 #5
class Solution:
	def combinationSum2(self, candidates, target):
		def dfs(t,res,cur,numList):
			if t==0:
				res.append(cur)
			for i in range(len(numList)):
				if t-numList[i]>=0:
					if i>0 and numList[i]==numList[i-1]:
						continue
					else:
						# print(t-numList[i],res,cur+[numList[i]],numList[i+1:])
						dfs(t-numList[i],res,cur+[numList[i]],numList[i+1:])
		candidates.sort()
		res=[]
		dfs(target,res,[],candidates)
		return res

slt=Solution()
inp=[10,1,2,7,6,1,5]
t=8
ans=[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
print(slt.combinationSum2(inp,t))
inp=[2,5,2,1,2]
t=5
ans=[
  [1,2,2],
  [5]
]
print(slt.combinationSum2(inp,t))
