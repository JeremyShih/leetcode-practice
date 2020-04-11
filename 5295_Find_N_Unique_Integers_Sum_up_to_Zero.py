# 2019/12/29
class Solution:
	def sumZero(self, n: int):
		if n<=1: return [0]
		ans=[i for i in range(n-1)]
		ans.append(sum(ans)*-1)
		return ans


slt=Solution()
n=5
print(sum(slt.sumZero(n))==0)
n=3
print(sum(slt.sumZero(n))==0)
n=1
print(sum(slt.sumZero(n))==0)
