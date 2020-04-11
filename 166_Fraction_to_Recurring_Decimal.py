# 2019/11/30
class Solution:
	def fractionToDecimal(self, numerator: int, denominator: int) -> str:
		ans=""
		if numerator*denominator<0:
			ans+="-"
		return ans

slt=Solution()
numerator = 1
denominator = 2
a="0.5"
print(slt.fractionToDecimal(numerator,denominator)==a)
numerator = 2
denominator = 1
a="2"
print(slt.fractionToDecimal(numerator,denominator)==a)
numerator = 2
denominator = 3
a="0.(6)"
print(slt.fractionToDecimal(numerator,denominator)==a)
