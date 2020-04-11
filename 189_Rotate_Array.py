class Solution:
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        tmp=list(nums)
        len_=len(nums)
        for i in range(len_):
        	nums[i]=tmp[(i-k)%len_]

#Input: [1,2,3,4,5,6,7] and k = 3
#indexs: [4,5,6,7,1,2,3]
#Output: [5,6,7,1,2,3,4]
slt=Solution()
inp=[1,2,3,4,5,6,7]
k=3
print(inp,k)
slt.rotate(inp,k)
print(inp)