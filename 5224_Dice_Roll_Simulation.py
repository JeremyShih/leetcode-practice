# 2019/10/13 contest not solved
class Solution:
	def dieSimulator(self, n: int, rollMax: List[int]) -> int:
		return 0

slt=Solution()
n, rollMax = 2, [1,1,2,2,2,3]
ans=34
print(slt.dieSimulator(n,rollMax)==ans)
n, rollMax = 2, [1,1,1,1,1,1]
ans=30
print(slt.dieSimulator(n,rollMax)==ans)
n, rollMax = 3, [1,1,1,2,2,3]
ans=181
print(slt.dieSimulator(n,rollMax)==ans)
