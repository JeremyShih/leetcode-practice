class Solution:
    def threeSumClosest(self, nums, target):
        nums.sort()
        i,j,k=0,1,len(nums)-1
        ret=nums[i]+nums[j]+nums[k]
        for i in range(len(nums)-2):
        	j,k=i+1,len(nums)-1
        	while j<k:
        		s=nums[i]+nums[j]+nums[k]
        		if s==target:
        			return s

        		if abs(s-target)<abs(ret-target):
        			ret=s
        		if s>target:
        			k-=1
        		else:
        			j+=1
        return ret

sol=Solution()
nums=[-1, 2, 1, -4]
targ=1
print(nums)
import time
start_time = time.time()
ans=sol.threeSumClosest(nums,targ)
print(ans==2)
nums=[0,2,1]
targ=0
ans=sol.threeSumClosest(nums,targ)
print(ans==3)
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)