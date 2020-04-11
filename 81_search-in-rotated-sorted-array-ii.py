class Solution:
	def search(self, nums, target):
		l, r = 0, len(nums)-1
		while l<=r:
			m = (l+r)//2
			if target == nums[m]:
				return True
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
		return False

slt=Solution()
nums = [0,0,0,0,0,0,0]
target = 0
print(slt.search(nums,target))
nums = [2,5,6,0,0,1,2]
target = 3
print(not slt.search(nums,target))
nums = [2,5,6,0,0,1,2]
target = 2
print(slt.search(nums,target))
nums = [2,5,6,0,0,1,2]
target = 4
print(not slt.search(nums,target))
nums = []
target = 4
print(not slt.search(nums,target))
nums = [2,5,6,0,0,1,2]
target = 4
print(not slt.search(nums,target))
nums = [1, 1, 2, 2, 5, 7, 9, 11, 11, 13, 14, 15, 16, 17, 19, 20, 24, 24, 24, 24, 25, 27, 28, 31, 32, 32, 32, 32, 34, 35, 38, 41, 42, 43, 43, 44, 44, 45, 46, 47, 47, 47, 50, 50, 51, 52, 53, 55, 55, 57, 60, 62, 64, 65, 65, 67, 68, 68, 72, 72, 72, 73, 73, 75, 75, 77, 78, 79, 79, 80, 80, 80, 81, 81, 81, 86, 86, 86, 87, 87, 87, 88, 88, 89, 90, 91, 92, 92, 92, 93, 93, 94, 95, 95, 96, 96, 96, 97, 98, 99, 1, 1]
target = 1
print(slt.search(nums,target))
nums = [28, 31, 32, 32, 32, 32, 34, 35, 38, 41, 42, 43, 43, 44, 44, 45, 46, 47, 47, 47, 50, 50, 51, 52, 53, 55, 55, 57, 60, 62, 64, 65, 65, 67, 68, 68, 72, 72, 72, 73, 73, 75, 75, 77, 78, 79, 79, 80, 80, 80, 81, 81, 81, 86, 86, 86, 87, 87, 87, 88, 88, 89, 90, 91, 92, 92, 92, 93, 93, 94, 95, 95, 96, 96, 96, 97, 98, 99, 1, 1, 2, 2, 5, 7, 9, 11, 11, 13, 14, 15, 16, 17, 19, 20, 24, 24, 24, 24, 25, 27]
target = 9
print(slt.search(nums,target))
nums = [1, 3, 1, 1]
target = 1
print(slt.search(nums,target))
nums = [1, 1, 3, 1]
target = 3
print(slt.search(nums,target))
