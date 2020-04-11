class Solution(object):
    def findMin(self, nums):
        if len(nums) == 1: return nums[0]
        left,right = 0,len(nums) - 1
        if nums[right] > nums[0]:
            return nums[0]
        while left+1<right:
            mid=(left+right)//2
            if nums[mid]>nums[right]:
                left=mid+1
            elif nums[mid]<nums[right]:
                right=mid
            else:
                while nums[mid]==nums[right] and mid<right:
                    mid+=1
                if mid<right:
                    return min(nums[mid],nums[right])
                else:
                    right=(left+right)//2
        return min(nums[right],nums[left])

slt=Solution()
inp=[1,3,5]
print(slt.findMin(inp)==1)
inp=[2,2,2,0,1]
print(slt.findMin(inp)==0)
inp=[2,2]
print(slt.findMin(inp)==2)
inp=[3,3,3,1]
print(slt.findMin(inp)==1)
inp=[3,3,3,3,3,3,3,3,1,3]
print(slt.findMin(inp)==1)
inp=[3,3,1,3,3,3,3,3,3,3]
print(slt.findMin(inp)==1)
inp=[-9,-9,-9,-8,-8,-7,-7,-7,-7,-6,-6,-6,-6,-6,-6,-6,-6,-6,-5,-5,-5,-5,-5,-4,-4,-4,-3,-3,-3,-3,-3,-3,-2,-2,-2,-2,-2,-2,-2,-2,-2,-2,-1,-1,-1,-1,-1,-1,0,0,0,1,1,2,2,2,2,2,2,2,3,3,3,3,4,4,4,4,4,5,5,5,5,5,5,5,6,6,6,7,7,7,7,7,8,9,9,9,10,10,10,10,10,10,10,-10,-9,-9,-9,-9]
print(slt.findMin(inp)==-10)
inp=[10,10,10,10,10,10,10,-10,-9,-9,-9,-9]
print(slt.findMin(inp)==-10)
inp=[1,1,3,1]
print(slt.findMin(inp)==1)
