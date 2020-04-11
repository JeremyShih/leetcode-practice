class Solution:
	def search(self, nums, target):
		l, r = 0, len(nums)-1
		if not nums: return -1
		if nums[l]==target:
			return l
		if nums[r]==target:
			return r
		while l<=r:
			m = (l+r)//2
			if target == nums[m]:
				return m
			elif nums[m] < target:
				# target is larger than the middle number
				# but the smallest number might be on the left
				if nums[r]<target and nums[m]<=nums[r]:
					#smallest number is on the left
					r = r-1
				else:
					l = l+1
			else:
				if nums[l]>target and nums[m]>=nums[r]:
					l = l+1
				else:
					r = r-1
		return -1

slt=Solution()
nums = [4,5,6,7,0,1,2]
target = 0
print(slt.search(nums,target)==4)
nums = [4,5,6,7,0,1,2]
target = 3
print(slt.search(nums,target)==-1)
