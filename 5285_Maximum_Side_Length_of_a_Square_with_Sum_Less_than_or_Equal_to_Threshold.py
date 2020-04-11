# 2019/12/15
class Solution:
	def maxSideLength(self, mat, threshold: int) -> int:
		return 0

slt=Solution()
mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]]
threshold = 4
a=2
print(slt.maxSideLength(mat,threshold)==a)
mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]]
threshold = 1
a=0
print(slt.maxSideLength(mat,threshold)==a)
mat = [[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]]
threshold = 6
a=3
print(slt.maxSideLength(mat,threshold)==a)
mat = [[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]]
threshold = 40184
a=2
print(slt.maxSideLength(mat,threshold)==a)
