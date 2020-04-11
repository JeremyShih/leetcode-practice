class Solution:
    def threeSum(self, nums):
        ans=set()
        nums.sort()
        i,j,k=0,1,len(nums)-1
        for i in range(len(nums)-2):
            if i>0 and nums[i]==nums[i-1]:
                continue
            j,k=i+1,len(nums)-1
            while j<k:
                sum_=nums[i]+nums[j]+nums[k]
                if sum_>0:
                    k-=1
                elif sum_<0:
                    j+=1
                else:
                    tri=(nums[i],nums[j],nums[k])
                    ans.add(tri)
                    j,k=j+1,k-1
        return list(ans)

sol=Solution()
inp=[-4, -1, -1, 0, 1, 2]
#inp=[-2, 0, 1, 1, 2]
#inp=[-4, -4, -4, -4, -3, -3, -2, -1, 0, 2, 2, 2, 3, 3]
print(sorted(inp))
import time
start_time = time.time()
print(sol.threeSum(inp))
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)
"""
class Solution:
    def threeSum(self, nums):
        ans=[]
        nums=sorted(nums)
        for i in range(len(nums)-2):
            if nums[i]>=0: return ans
            subLst=nums[i+1:]
            for j in range(len(subLst)-1):
                if -nums[i]-subLst[j] in subLst[j+1:]:
                    tri=[nums[i],subLst[j],-nums[i]-subLst[j]]
                    if tri not in ans:
                        ans.append(tri)
        return ans

"""