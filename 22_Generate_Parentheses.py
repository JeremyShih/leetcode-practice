# 2019/8/18
class Solution:
	def generateParenthesis(self, n: int) -> list:
		if n==0:
			return []
		ans=[]
		def dfs(cur,leftCount,rightCount):
			print(cur)
			if len(cur)==n*2:
				ans.append(cur)
				return
			if leftCount==n:
				dfs(cur+")",leftCount,rightCount+1)
			elif leftCount==rightCount:
				dfs(cur+"(",leftCount+1,rightCount)
			else:
				dfs(cur+"(",leftCount+1,rightCount)
				dfs(cur+")",leftCount,rightCount+1)
		dfs("",0,0)
		return ans

slt=Solution()
n=0
print(slt.generateParenthesis(n))
n=1
print(slt.generateParenthesis(n))
n=2
print(slt.generateParenthesis(n))
n=3
print(slt.generateParenthesis(n))



# runtime 176s
# class Solution:
# 	def generateParenthesis(self, n: int) -> list:
# 		if n==0:
# 			return []
# 		ans=[]
# 		def dfs(ans,cur,pareList):
# 			print(cur,pareList)
# 			if len(cur)==n*2:
# 				# print(cur)
# 				if self.isValid(cur):
# 					print('add')
# 					ans.append(cur)
# 				return
# 			dfs(ans,cur+pareList[0],pareList[1:])
# 			print(len(pareList))
# 			# if len(pareList)>1:
# 			dfs(ans,cur+pareList[-1],pareList[:-1])
# 		dfs(ans,"","("*n+")"*n)
# 		return ans
# 	def isValid(self, s: str) -> bool:
# 		stack=[]
# 		def validInput(c1,c2):
# 			return (c1 == '(' and c2 == ')')
# 		for c in s:
# 			if c in "(":
# 				stack.append(c)
# 			else:
# 				if len(stack)!=0 and validInput(stack[-1],c):
# 					stack.pop()
# 				else:
# 					return False
# 		return len(stack)==0