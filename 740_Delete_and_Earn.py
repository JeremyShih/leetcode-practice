class Solution:
    def deleteAndEarn(self, nums):
        if len(nums)==0:
        	return 0
        dic={}
        dp=[0]
        nums.sort()
        for n in nums:
        	dic[n]=dic.get(n,0)+n
        oldKey=0
        for key in dic:
        	if oldKey==0 or key!=oldKey+1: #dont need to delete
        		dp.append(dp[-1]+dic[key])
        	else:
        		dp.append(max(dp[-2]+dic[key],dp[-1])) #choose to earn or delete
        	oldKey=key
        return dp[-1]



import time
start_time = time.time()
slt=Solution()
print(slt.deleteAndEarn([3, 4, 2])==6)
print(slt.deleteAndEarn([2, 2, 3, 3, 3, 4])==9)
elapsed_time = time.time() - start_time
print(elapsed_time)