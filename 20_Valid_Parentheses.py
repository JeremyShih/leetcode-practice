# 2019/8/19 rewrite in Python
class Solution:
	def isValid(self, s: str) -> bool:
		stack=[]
		def validInput(c1,c2):
			return (c1 == '(' and c2 == ')') or (c1 == '{' and c2 == '}') or (c1 == '[' and c2 == ']')
		for c in s:
			if c in "([{":
				stack.append(c)
			else:
				if len(stack)!=0 and validInput(stack[-1],c):
					stack.pop()
				else:
					return False
		return len(stack)==0

slt=Solution()
inp="(()())"
print(slt.isValid(inp))
inp="[][]{}"
print(slt.isValid(inp))
inp="[][){]"
print(not slt.isValid(inp))
inp="[{][}]"
print(not slt.isValid(inp))
