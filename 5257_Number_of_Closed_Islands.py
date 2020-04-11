# 2019/11/10
class Solution:
	def closedIsland(self, grid) -> int:
		def replace(matrix, t, r):
			for i in range(len(matrix)):
				for j in range(len(matrix[0])):
					if matrix[i][j]==t:
						matrix[i][j]=r
		replace(grid,"O","-")
		# print(grid)

		for i in range(len(grid[0])):
			self.checkSurround(grid,0,i,0,"-")
			self.checkSurround(grid,len(grid)-1,i,0,"-")
		for i in range(len(grid)):
			self.checkSurround(grid,i,0,0,"-")
			self.checkSurround(grid,i,len(grid[0])-1,0,"-")
		print(grid)

		return 0

	def checkSurround(self, matrix, row, col, cur, t):
		t,r=t,cur
		# print(matrix,row,col)
		if not matrix[row][col]==r:
			return
		matrix[row][col]=t
		if row+1<len(matrix) and matrix[row+1][col]==r:
			self.checkSurround(matrix, row+1 ,col, r, t)
		if row-1>=0 and matrix[row-1][col]==r:
			self.checkSurround(matrix, row-1 ,col, r, t)
		if col+1<len(matrix[0]) and matrix[row][col+1]==r:
			self.checkSurround(matrix, row, col+1, r, t)
		if col-1>=0 and matrix[row][col-1]==r:
			self.checkSurround(matrix, row, col-1, r, t)

slt=Solution()
grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
a=2
print(slt.closedIsland(grid),"=>",a)
grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
a=1
print(slt.closedIsland(grid),"=>",a)
grid = [[1,1,1,1,1,1,1],
               [1,0,0,0,0,0,1],
               [1,0,1,1,1,0,1],
               [1,0,1,0,1,0,1],
               [1,0,1,1,1,0,1],
               [1,0,0,0,0,0,1],
               [1,1,1,1,1,1,1]]
a=2
print(slt.closedIsland(grid),"=>",a)
