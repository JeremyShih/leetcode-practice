# 2019/11/10
class Solution:
	def reconstructMatrix(self, upper: int, lower: int, colsum):
		if sum(colsum)!=upper+lower: return []
		width=len(colsum)
		matrix=[[0]*width for _ in range(2)]
		# self.printMap(matrix)
		for i in range(width):
			if colsum[i]==2:
				matrix[0][i],matrix[1][i]=1,1
				upper-=1
				lower-=1
				colsum[i]=0
		print(upper,lower)
		for i in range(width):
			if upper<0 or lower<0:
				return []
			if colsum[i]==1:
				if upper>0:
					matrix[0][i]=1
					upper-=1
				elif lower>0:
					# print(i,1)
					matrix[1][i]=1
					lower-=1
				else:
					return []
		# self.printMap(matrix)
		return matrix

	def printMap(self, matrix):
		for i in range(len(matrix)):
			print(matrix[i])


slt=Solution()
u,l=2,1
col=[1,1,1]
# print(slt.reconstructMatrix(u,l,col))
# u,l=2,3
# col=[2,2,1,1]
# print(slt.reconstructMatrix(u,l,col))
# u,l=5,5
# col=[2,1,2,0,1,0,1,2,0,1]
# print(slt.reconstructMatrix(u,l,col))
u,l=9,2
col=[0,1,2,0,0,0,0,0,2,1,2,1,2]
print(slt.reconstructMatrix(u,l,col))
