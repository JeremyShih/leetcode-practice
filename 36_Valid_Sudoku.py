# 2019/8/21
class Solution:
	def isValidSudoku(self, board) -> bool:
		row = {}
		col = {}
		block = {}
		for i, x in enumerate(board):
			for j, y in enumerate(x):
				if y != ".":
					# print(i/3,j/3,y)
					if (i,y) in row or (j,y) in col or (i//3,j//3,y) in block:
						print(block)
						return False
					else:
						row[i,y] = 1
						col[j,y] = 1
						block[i//3,j//3,y] = 1
		# print(block)
		return True

slt=Solution()
n=[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
print(slt.isValidSudoku(n))
n=[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4","7",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
print(not slt.isValidSudoku(n))



# class Solution:
# 	def isValidSudoku(self, board) -> bool:
# 		tmp,tmp2=[],[]
# 		for i in range(9):
# 			for j in range(9):
# 				if board[i][j]!=".":
# 					if board[i][j] in tmp:
# 						return False
# 					else:
# 						tmp.append(board[i][j])
# 				if board[j][i]!=".":
# 					if board[j][i] in tmp2:
# 						return False
# 					else:
# 						tmp2.append(board[j][i])
# 			tmp,tmp2=[],[]
# 		for x in range(0,9,3):
# 			for y in range(0,9,3):
# 				for i in range(x,x+3):
# 					for j in range(y,y+3):
# 						if board[i][j]!=".":
# 							if board[i][j] in tmp:
# 								return False
# 							else:
# 								tmp.append(board[i][j])
# 				tmp=[]
# 		return True