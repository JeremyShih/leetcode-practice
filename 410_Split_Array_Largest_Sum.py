class Solution:
	def splitArray(self, nums, m):
		max_=max(nums)
		sum_=sum(nums)
		if m==1:
			return sum_
		l,r=max_,sum_
		def valid(tar, nums, m):
			count=1
			total=0
			for n in nums:
				total+=n
				if total>tar:
					total=n
					count+=1
					if count>m:
						return False
			return True
		while l<=r:
			mid=(l+r)//2
			if valid(mid,nums,m):
				r=mid-1
			else:
				l=mid+1
		return l

slt=Solution()
nums=[7,2,5,10,8]
m=2
print(slt.splitArray(nums,m)==18)