class Solution:
    def calculate(self, s: str) -> int:
        stack, i = [], 0

        while i < len(s):
            if s[i] in ('+', '-', '('):
                stack.append(s[i])
            elif '9' >= s[i] >= '0':
                start = i
                while i < len(s) and s[start: i + 1].isnumeric():
                    i += 1
                i -= 1
                num = int(s[start: i + 1])

			# left operand
                if not stack or stack[-1] not in ('+', '-'):
                    stack.append(num)

			# right operand
                else:
                    stack.append(self.evalSingle(num, stack))

            elif s[i] == ')':
                num = stack.pop()
                stack.pop()
                if len(stack) > 0 and stack[-1] in ('+', '-'):
                    num = self.evalSingle(num, stack)
                stack.append(num)
            i += 1
        return stack[0]
    
    def evalSingle(self, num, stack):
        oper = stack.pop()
        num0 = stack.pop()
        num = num0 + num if oper == '+' else num0 - num
        return num

sol = Solution()
s = "1 + 1"
print(sol.calculate(s)==2)
s = " 2-1 + 2 "
print(sol.calculate(s)==3)
s = "(1+(4+5+2)-3)+(6+8)"
print(sol.calculate(s)==23)
s = "2-(5-6)"
print(sol.calculate(s)==3)
