# 2019/9/21
class Solution:
	def myPow(self, x: float, n: int) -> float:
		if n == 0 or x == 1.0:
			return 1.0
		if x == 0:
			if n < 0:
				return float('inf')
			else:
				return 0.0
		
		if n < 0:
			x, n = 1/x, -n
		count=0
		res, num, power = 1.0, x, n
		while power != 1:
			count+=1
			if power%2 == 0:
				num = num*num
				power /= 2
			else:
				res *= num
				power -= 1
		print(count)
		return res*num

slt=Solution()
x,n=2,50
print(slt.myPow(x,n)==pow(x,n))
x,n=2.1,3
print(slt.myPow(x,n)-9.261<0.000001)
x,n=2,-2
print(slt.myPow(x,n)==0.25)
