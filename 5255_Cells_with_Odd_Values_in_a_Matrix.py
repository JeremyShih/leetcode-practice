# 2019/11/10
class Solution:
	def oddCells(self, n: int, m: int, indices) -> int:
		matrix=[[0]*m for _ in range(n)]
		# self.printMap(matrix)
		for i in indices:
			for x in range(m):
				# print("i[0],x =",i[0],x)
				matrix[i[0]][x]+=1
			# self.printMap(matrix)
			for y in range(n):
				# print("y,i[1] =",y,i[1])
				matrix[y][i[1]]+=1
			# self.printMap(matrix)
		
		# self.printMap(matrix)
		count=0
		for i in range(n):
			for j in range(m):
				if matrix[i][j]%2==1:
					count+=1
		return count

	def printMap(self, matrix):
		for i in range(len(matrix)):
			print(matrix[i])


slt=Solution()
n,m=2,3
indi=[[0,1],[1,1]]
ans=6
print(slt.oddCells(n,m,indi),"=>",ans)
n,m=2,2
indi=[[1,1],[0,0]]
ans=0
print(slt.oddCells(n,m,indi),"=>",ans)
n,m=48,37
indi=[[40,5]]
ans=83
print(slt.oddCells(n,m,indi),"=>",ans)
