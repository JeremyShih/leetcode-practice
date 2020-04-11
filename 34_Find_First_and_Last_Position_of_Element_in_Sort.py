class Solution:
	def searchRange(self, nums, target):
		if not nums: return [-1,-1]
		l,r,mid=0,len(nums)-1,0
		ll,rr=l-1,r+1
		while l<r:
			mid=(l+r)//2
			if nums[mid]>target:
				r,rr=mid-1,mid
			elif nums[mid]<target:
				l,ll=mid+1,mid
			else:
				l=r=mid
		if nums[mid]!=target:
			return [-1,-1]
		while ll<l:
			mid=(ll+l+1)//2
			if nums[mid]==target:
				l=mid-1
			else:
				ll=mid
		while r<rr:
			mid=(r+rr)//2
			print(r,rr,mid)
			if nums[mid]==target:
				r=mid+1
			else:
				rr=mid
		print(ll+1,rr-1)
		return [ll+1,rr-1]

slt=Solution()
inp=[5,7,7,8,8,10]
tar=8
ans=[3,4]
print(slt.searchRange(inp,tar)==ans)
inp=[5,7,7,8,8,10]
tar=6
ans=[-1,-1]
#print(slt.searchRange(inp,tar)==ans)
inp=[5,7,7,8,8,10]
tar=7
ans=[1,2]
print(slt.searchRange(inp,tar)==ans)
inp=[5,7,7,8,8,10]
tar=5
ans=[0,0]
print(slt.searchRange(inp,tar)==ans)
inp=[1]
tar=1
ans=[0,0]
print(slt.searchRange(inp,tar)==ans)
inp=[1]
tar=0
ans=[-1,-1]
#print(slt.searchRange(inp,tar)==ans)
