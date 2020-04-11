# 2018/11/11 #10
class Solution:
	def triangleNumber(self, nums):
		if not nums or len(nums)<3: return 0
		nums.sort()
		res=0
		i,k=0,len(nums)-1
		j=k-1
		while i<k:
			while i<j:
				if nums[i]+nums[j]<=nums[k]:
					i+=1
				else:
					res+=j-i
					j-=1
			i,k=0,k-1
			j=k-1
		return res

slt=Solution()
import time
start_time = time.time()
inp=[2,2,3,4]
ans=3
print(slt.triangleNumber(inp)==ans)
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)
start_time = time.time()
inp=[86,2,54,7,10,60,25,87,66,7,68,10,19,70,40,15,46,71,9,34,74,24,8,33,80,71,1,29,19,38,20,56,14,67,70,100,74,33,79,6,92,83,70,22,14,88,48,38,61,1,76,53,67,6,23,39,66,50,54,17,1,61,43,57,0,45,55,56,21,94,70,22,100,20,75,25,28,5,80,48,8,33,100,69,10,34,52,21,38,20,41,84,89,99,83,18,43,11,61,65]
ans=66946
print(slt.triangleNumber(inp))
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)
