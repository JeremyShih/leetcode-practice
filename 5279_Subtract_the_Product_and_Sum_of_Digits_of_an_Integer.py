# 2019/12/8
class Solution:
	def subtractProductAndSum(self, n: int) -> int:
		p,s=1,0
		for c in str(n):
			p*=int(c)
			s+=int(c)
		return p-s

slt=Solution()
n=234
a=15
print(slt.subtractProductAndSum(n)==a)
n=4421
a=21
print(slt.subtractProductAndSum(n)==a)
