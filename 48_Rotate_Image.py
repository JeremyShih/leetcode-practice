#2018/11/4 倒數第三題
class Solution:
	def rotate(self, matrix):
		def verticalFlip(m):
			for j in range(len(m)//2):
				for i in range(len(m[0])):
					m[j][i],m[-j-1][i]=m[-j-1][i],m[j][i]
		def diagonalFlip(m):
			for j in range(len(m)):
				for i in range(j):
					m[j][i],m[i][j]=m[i][j],m[j][i]
		verticalFlip(matrix)
		diagonalFlip(matrix)

slt=Solution()
inp=[
  [1,2,3],
  [4,5,6],
  [7,8,9]
]
ans=[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
slt.rotate(inp)
print(inp)
inp=[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
]
ans=[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
slt.rotate(inp)
print(inp)
