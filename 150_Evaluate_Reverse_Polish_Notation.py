# 2019/11/3
class Solution:
	def evalRPN(self, tokens) -> int:
		oper,stack=["+","-","*","/"],[]
		for n in tokens:
			if n not in oper:
				stack.append(int(n))
			else:
				second=stack.pop()
				first=stack.pop()
				if n=="+":
					stack.append(first+second)
				elif n=="-":
					stack.append(first-second)
				elif n=="*":
					stack.append(first*second)
				else:
					stack.append(int(first/second))
		return stack[0]

slt=Solution()
n=["2", "1", "+", "3", "*"]
a=9
print(slt.evalRPN(n)==a)
n=["4", "13", "5", "/", "+"]
a=6
print(slt.evalRPN(n)==a)
n=["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
a=22
print(slt.evalRPN(n)==a)