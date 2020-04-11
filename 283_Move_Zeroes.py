class Solution:
	def moveZeroes(self, nums):
		if not nums: return
		i,j,n=0,len(nums)-1,0
		while i<j:
			if nums[i]==0:
				nums.pop(i)
				n+=1
				j-=1
			else:
				i+=1
		nums+=[0]*n
		# i,j=0,1
		# while j<len(nums):
		# 	if nums[i]!=0:
		# 		i,j=i+1,i+2
		# 	else:
		# 		while nums[j]==0:
		# 			j+=1
		# 		print(nums[i],nums[j])
		# 		nums[i],nums[j]=nums[j],0
		# 		print(nums[i],nums[j])
		# 		i,j=j,j+1

slt=Solution()
inp=[0,1,0,3,12]
ans=[1,3,12,0,0]
slt.moveZeroes(inp)
print(inp==ans)
inp=[0,1]
ans=[1,0]
slt.moveZeroes(inp)
print(inp==ans)
