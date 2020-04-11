from functools import reduce
class Solution:
	def superPow(self, a, b):
		return 0 if a % 1337 == 0 else pow(a, reduce(lambda x, y: (x * 10 + y) % 1140, b), 1337)

slt=Solution()
a=2
b=[3]
print(slt.superPow(a,b))
a=2
b=[1,0]
print(slt.superPow(a,b))
a=5
b=[6]
print(slt.superPow(a,b))