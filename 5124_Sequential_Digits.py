# 2019/12/15
class Solution:
	def sequentialDigits(self, low: int, high: int):
		ans,num=[],0
		def next(n):
			# n=n+pow(10,len(str(n))-1)
			s=str(n)
			if int(s[0])+len(s)>10:
				s="1"+s
			ans=int(s[0])
			for i in range(1,len(s)):
				m=ans%10
				ans=ans*10+m+1
			# ans=int(s)
			if ans<n:
				ans=next(ans+pow(10,len(str(ans))-1))
			return ans
		num=next(low)
		while num<high:
			# print(ans,num)
			ans.append(num)
			num=next(num+1)
		return ans

slt=Solution()
low = 100
high = 300
a = [123,234]
print(slt.sequentialDigits(low,high)==a)
low = 1000
high = 13000
a = [1234,2345,3456,4567,5678,6789,12345]
print(slt.sequentialDigits(low,high)==a)
low = 234
high = 2314
a = [234,345,456,567,678,789,1234]
print(slt.sequentialDigits(low,high)==a)
