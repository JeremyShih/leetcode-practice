# 2019/10/13
class Solution:
	def solve(self, board) -> None:
		# if len(board) == 0 or len(board[0]) == 0:
		# 	return
		# def floodFill(board, r, c, target='-', replace='O'):
		# 	if board[r][c] != target:
		# 		return
		# 	else:
		# 		board[r][c] = replace
		# 		if r + 1 < len(board) and board[r + 1][c] == target:
		# 			floodFill(board, r + 1, c)
		# 		if r - 1 >= 0 and board[r - 1][c] == target:
		# 			floodFill(board, r - 1, c)
		# 		if c + 1 < len(board[0]) and board[r][c + 1] == target:
		# 			floodFill(board, r, c + 1)
		# 		if c - 1 >= 0 and board[r][c - 1] == target:
		# 			floodFill(board, r, c - 1)
		
		# def replace(matrix, target, value):
		# 	for i in range(len(matrix)):
		# 		for j in range(len(matrix[0])):
		# 			if matrix[i][j] == target:
		# 				matrix[i][j] = value
		# # replace all "O" with "-"
		# replace(board, 'O', '-')
		# print(board)

		# # mark all "O" connected to border
		# # Row
		# for i in range(len(board[0])):
		# 	# if board[0][i] == '-':
		# 		floodFill(board, 0, i)
		# 	# if board[-1][i] == '-':
		# 		floodFill(board, len(board) - 1, i)
		# # Column
		# for i in range(len(board)):
		# 	if board[i][0] == '-':
		# 		floodFill(board, i, 0)
		# 	if board[i][len(board[0]) - 1] == '-':
		# 		floodFill(board, i, len(board[0]) - 1)
				
		# # replace all left "-" with "X", which are all the "O"s not connected to border
		# replace(board, '-', 'X')
		
		if len(board)==0 or len(board[0])==0:
			return

		def checkSurround(matrix, row, col):
			t,r="-","O"
			# print(matrix,row,col)
			if not matrix[row][col]==t:
				return
			matrix[row][col]=r
			if row+1<len(matrix) and matrix[row+1][col]==t:
				checkSurround(matrix, row+1 ,col)
			if row-1>=0 and matrix[row-1][col]==t:
				checkSurround(matrix, row-1 ,col)
			if col+1<len(matrix[0]) and matrix[row][col+1]==t:
				checkSurround(matrix, row, col+1)
			if col-1>=0 and matrix[row][col-1]:
				checkSurround(matrix, row, col-1)
		def replace(matrix, t, r):
			for i in range(len(matrix)):
				for j in range(len(matrix[0])):
					if matrix[i][j]==t:
						matrix[i][j]=r
		replace(board,"O","-")
		# print(board)

		for i in range(len(board[0])):
			checkSurround(board,0,i)
			checkSurround(board,len(board)-1,i)
		# print(board)
		for i in range(len(board)):
			checkSurround(board,i,0)
			checkSurround(board,i,len(board[0])-1)
		# print(board)
		
		replace(board,"-","X")

slt=Solution()
board=[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
ans=[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
# slt.solve(board)
print(board==ans)
board=[["X","X","X","X"],["X","O","O","X"],["X","O","X","X"],["X","O","X","X"]]
ans=[["X","X","X","X"],["X","O","O","X"],["X","O","X","X"],["X","O","X","X"]]
# slt.solve(board)
print(board==ans)
board=[["X","O","X","O","X","O"],["O","X","O","X","O","X"],["X","O","X","O","X","O"],["O","X","O","X","O","X"]]
slt.solve(board)
print(board)
