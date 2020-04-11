# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:
	def firstBadVersion(self, n):
		if n==0: return 0
		def isBadVersion(ver):
			#test=[0,0,0,1,1]
			if test[ver-1]==1:
				return True
			else:
				return False
		l,r=1,n
		while l<r:
			mid=(l+r)//2
			if isBadVersion(mid):
				r=mid
			else:
				l=mid+1
		return r

slt=Solution()
n=5
test=[0,0,0,1,1]
ans=4
print(slt.firstBadVersion(n)==ans)
n=5
test=[0,0,1,1,1]
ans=3
print(slt.firstBadVersion(n)==ans)
n=5
test=[1,1,1,1,1]
ans=1
print(slt.firstBadVersion(n)==ans)
