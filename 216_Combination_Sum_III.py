class Solution:
	def combinationSum3(self, k, n):
		def dfs(k,n,res,cur,numList):
			if k==1:
				if n in numList:
					cur+=[n]
					res.append(cur)
			else:
				for i in range(len(numList)):
					dfs(k-1,n-numList[i],res,cur+[numList[i]],numList[i+1:])
		res=[]
		dfs(k,n,res,[],[i for i in range(1,10)])
		print(res)
		return res


slt=Solution()
k=3
n=7
print(slt.combinationSum3(k,n)==[[1,2,4]])
n=9
print(slt.combinationSum3(k,n)==[[1,2,6], [1,3,5], [2,3,4]])

