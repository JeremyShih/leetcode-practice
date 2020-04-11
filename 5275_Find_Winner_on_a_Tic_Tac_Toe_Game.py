# 2019/12/1
class Solution:
	def tictactoe(self, moves) -> str:
		if len(moves)<5:
			return "Pending"
		board=[[" "]*3 for _ in range(3)]
		# print(board)
		i=0
		while i<len(moves):
			# print(moves[i])
			board[moves[i][0]][moves[i][1]]="X"
			if i+1<len(moves):
				board[moves[i+1][0]][moves[i+1][1]]="O"
			i+=2
		# for n in board:
		# 	print(n)
		for i in range(3):
			if board[i][0]==board[i][1]==board[i][2]!=" ":
				if board[i][0]=="X":
					return "A"
				else:
					return "B"
			elif board[0][i]==board[1][i]==board[2][i]!=" ":
				if board[0][i]=="X":
					return "A"
				else:
					return "B"
		if board[0][0]==board[1][1]==board[2][2]!=" ":
			if board[0][0]=="X":
				return "A"
			else:
				return "B"
		elif board[2][0]==board[1][1]==board[0][2]!=" ":
			if board[2][0]=="X":
				return "A"
			else:
				return "B"
		if len(moves)==9:
			return "Draw"
		else:
			return "Pending"

slt=Solution()
moves = [[0,0],[2,0],[1,1],[2,1],[2,2]]
a="A"
print(slt.tictactoe(moves)==a)
moves = [[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]
a="B"
print(slt.tictactoe(moves)==a)
moves = [[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]]
a="Draw"
print(slt.tictactoe(moves)==a)
moves = [[0,0],[1,1]]
a="Pending"
print(slt.tictactoe(moves)==a)
moves = [[0,2],[2,0],[2,1],[0,1],[1,2]]
a="Pending"
print(slt.tictactoe(moves)==a)
