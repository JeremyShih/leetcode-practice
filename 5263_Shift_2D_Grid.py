# 2019/11/17
class Solution:
	def shiftGrid(self, grid, k: int):
		if not grid or len(grid)*len(grid[0])==k:
			return grid
		h,w=len(grid),len(grid[0])
		ans=[[-1]*w for _ in range(h)]
		# print("h,w =",h,w)
		def shiftedIndex(i, j, height, width):
			# print(i,j,k,"->")
			j+=k
			while j>=width:
				i,j=(i+1)%height,j-width
			# print(i,j)
			return i,j
		for i in range(h):
			for j in range(w):
				m,n=shiftedIndex(i,j,h,w)
				# print(i,j,k,l)
				ans[m][n]=grid[i][j]
		return ans


slt=Solution()
grid = [[1,2,3],[4,5,6],[7,8,9]]
k=1
a=[[9,1,2],[3,4,5],[6,7,8]]
print(slt.shiftGrid(grid,k)==a)
grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]
k=4
a=[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
print(slt.shiftGrid(grid,k)==a)
grid = [[1,2,3],[4,5,6],[7,8,9]]
k=9
a=[[1,2,3],[4,5,6],[7,8,9]]
print(slt.shiftGrid(grid,k)==a)
