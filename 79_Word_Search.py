class Solution:
    def find(self,board,word,i,j):
    	s=board[i][j]
    	board[i][j]=' '
    	for r,c in [[-1,0],[0,-1],[1,0],[0,1]]:
    		if (i+r)>=0 and (i+r)<len(board) and (j+c)>=0 and (j+c)<len(board[i]):
    			if len(word)==1:
    				return True
    			if board[i+r][j+c]==word[1]:
    				#if len(word)==1:
    					#return True
    				if self.find(board,word[1:],i+r,j+c):
    					return True
    				else:
    					board[i][j]=word[0]
    			else:
    				board[i][j]=s
    	return False
    def exist(self, board, word):
    	if not word:
    		return True
    	for i in range(len(board)):
    		for j in range(len(board[i])):
    			if board[i][j]==word[0]:
    				if len(word)==1:
    					return True
    				if self.find(board,word,i,j):
    					return True
    				#else:
    					#board[i][j]=word[0]
    	return False

slt=Solution()
bo=[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
boa=bo
wo="ABCCED"
print(slt.exist(boa,wo))
print(boa)
boa=[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
wo="SEE"
print(slt.exist(boa,wo))
print(boa)
'''
boa=bo
wo=""
print(slt.exist(boa,wo))
boa=bo
wo="ABCB"
print(not slt.exist(boa,wo))
bo=[["a","a"]]
wo="aa"
print(slt.exist(bo,wo))
bo=[["a","a","a"]]
wo="aa"
print(slt.exist(bo,wo))
bo=[["a","a"]]
wo="aaa"
print(not slt.exist(bo,wo))
bo=[["a"],["a"]]
wo="aa"
print(slt.exist(bo,wo))
'''