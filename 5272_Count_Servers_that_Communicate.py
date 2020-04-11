# 2019/11/24
class Solution:
	def countServers(self, grid) -> int:
		com,count=0,0
		for i in range(len(grid)):
			tmp,ser=0,[]
			for j in range(len(grid[0])):
				if grid[i][j]==1 or str(grid[i][j])=="+":
					ser.append(j)
					tmp+=1
			if tmp>=2:
				for j in ser:
					grid[i][j]="+"
		for i in range(len(grid[0])):
			tmp,ser=0,[]
			for j in range(len(grid)):
				if grid[j][i]==1 or str(grid[j][i])=="+":
					ser.append(j)
					tmp+=1
			if tmp>=2:
				for j in ser:
					grid[j][i]="+"
		# print(grid)
		for i in range(len(grid)):
			for j in range(len(grid[0])):
				if grid[i][j]=="+":
					count+=1

		return count

slt=Solution()
grid = [[1,0],[0,1]]
a=0
print(slt.countServers(grid),a)
grid = [[1,0],[1,1]]
a=3
print(slt.countServers(grid),a)
grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
a=4
print(slt.countServers(grid),a)
