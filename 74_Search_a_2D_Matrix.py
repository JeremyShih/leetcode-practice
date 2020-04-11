class Solution:
    def searchMatrix(self, matrix, t):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix or not matrix[0]:
        	return False
        if t<matrix[0][0] or t>matrix[-1][-1]:
        	return False
        for r in range(len(matrix)):
        	for c in range(len(matrix[0])):
        		if t<matrix[r][0]:
        			return False
        		elif matrix[r][0]<=t and t<=matrix[r][-1]:
        			if t in matrix[r]:
        				return True
        			else:
        				return False
        return False



slt=Solution()
inp=[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
t=3
print(slt.searchMatrix(inp,t))
inp=[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
t=13
print(not slt.searchMatrix(inp,t))
inp=[]
print(not slt.searchMatrix(inp,t))
inp=[[]]
print(not slt.searchMatrix(inp,t))
