# 2019/10/27
"""
   This is the custom function interface.
   You should not implement it, or speculate about its implementation
   class CustomFunction:
       # Returns f(x, y) for any given positive integers x and y.
       # Note that f(x, y) is increasing with respect to both x and y.
       # i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
       def f(self, x, y):
  
"""
class Solution:
	def findSolution(self, customfunction: 'CustomFunction', z: int) -> List[List[int]]:
		ans=[]
		def search(target):
			l,r=1,1000
			while l<=r:
				mid=(l+r)//2
				if customfunction.f(1,mid)>target:
					r=mid-1
				elif customfunction.f(1,mid)<target:
					l=mid+1
				else:
					ans.append([i,j])
					return mid
			return l
		def searchLeft(fix,target):
			l,r=1,1000
			while l<=r:
				mid=(l+r)//2
				if customfunction.f(mid,fix)>target:
					r=mid-1
				elif customfunction.f(mid,fix)<target:
					l=mid+1
				else:
					return mid
			return -1
		i,j=1,search(z)
		# print(i,j,z,search(z))
		# return
		# while i<=j:
		# 	mid=(i+j)//2
		# 	if customfunction.f(i,j)>z:
		# 		j=mid-1
		# 	elif customfunction.f(i,j)<z:
		# 		i=mid+1
		# return [i,j]
		# if j==-1:
		# 	return ans
		while 1<j:
			j-=1
			i=searchLeft(j,z)
			if i>-1:
				ans.append([i,j])
		return ans

ans=[[1,29],[2,28],[3,27],[4,26],[5,25],[6,24],[7,23],[8,22],[9,21],[10,20],[11,19],[12,18],[13,17],[14,16],[15,15],[16,14],[17,13],[18,12],[19,11],[20,10],[21,9],[22,8],[23,7],[24,6],[25,5],[26,4],[27,3],[28,2],[29,1]]
exp=[[1,29],[2,28],[3,27],[4,26],[5,25],[6,24],[7,23],[8,22],[9,21],[10,20],[11,19],[12,18],[13,17],[14,16],[15,15],[16,14],[17,13],[18,12],[19,11],[20,10],[21,9],[22,8],[23,7],[24,6],[25,5],[26,4],[27,3],[28,2],[29,1]]
print(ans==exp)
ans=[[1,29],[2,28],[3,27],[4,26],[5,25],[6,24],[7,23],[8,22],[9,21],[10,20],[11,19],[12,18],[13,17],[14,16],[15,15],[16,14],[17,13],[18,12],[19,11],[20,10],[21,9],[22,8],[23,7],[24,6],[25,5],[26,4],[27,3],[28,2],[29,1]]
