class Solution:
	def combine(self, n, k):
		def recur(k,nums,res,cur):
			if k==len(cur):
				res.append(cur)
			else:
				for i in range(len(nums)):
					#print(k,nums[i+1:],res,cur+[nums[i]])
					recur(k,nums[i+1:],res,cur+[nums[i]])
		res=[]
		if k==1:
			return [ [j] for j in range(1,n+1) ]
		if n==k:
			return [list(range(1,n+1))]
		for i in range(1,n):
			recur(k,range(i+1,n+1),res,[i])
		return res

slt=Solution()
k=2
n=4
ans=[
  [1,2],
  [1,3],
  [1,4],
  [2,3],
  [2,4],
  [3,4],
]
print(slt.combine(n,k)==ans)
k=1
n=1
print(slt.combine(n,k)==[[1]])
k=1
n=3
print(slt.combine(n,k)==[[1],[2],[3]])
k=3
n=3
print(slt.combine(n,k)==[[1,2,3]])
