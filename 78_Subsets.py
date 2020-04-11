class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        r=[[nums[0]]]
        for n in nums[1:]:
        	len_=len(r)
        	r.append([n])
        	for i in range(len_):
        		r.append(r[i]+[n])
        r.append([])
        return r


slt=Solution()
inp=[1,2,3]
ans=[
  [3],
  [1],
  [2],
  [1,3],
  [2,3],
  [1,2],
  [1,2,3],
  []
]
print(slt.subsets(inp))