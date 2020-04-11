# 2019/10/20
class Solution:
	def balancedString(self, s: str) -> int:
		bal,dic="QWER",{}
		for c in bal:
			dic[c]=0
		for c in s:
			dic[c]+=1
		min_,ans,sum_=len(s),0,0
		for k in dic:
			sum_+=dic[k]
			min_=min(min_,dic[k])
		for k in dic:
			dic[k]-=min_
		# print(min_)
		sum_/=4
		for k in dic:
			if dic[k]>sum_:
				ans+=dic[k]-sum_

		print(dic)
		return int(ans)

slt=Solution()
s = "QWER"
r=0
print(slt.balancedString(s)==r)
s = "QQWE"
r=1
print(slt.balancedString(s)==r)
s = "QQQW"
r=2
print(slt.balancedString(s)==r)
s = "QQQQ"
r=3
print(slt.balancedString(s)==r)
s = "WQWRQQQW"
r=3
print(slt.balancedString(s)==r)
s="WWEQERQWQWWRWWERQWEQ"
r=4
print(slt.balancedString(s)==r)
