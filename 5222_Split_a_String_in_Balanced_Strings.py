# 2019/10/13
class Solution:
	def balancedStringSplit(self, s: str) -> int:
		r,l,ans=0,0,0

		while s:
			if s[0]=="R":
				r+=1
			else:
				l+=1
			s=s[1:]
			if r==l:
				r,l=0,0
				ans+=1
		return ans

slt=Solution()
s=""
ans=0
print(slt.balancedStringSplit(s)==ans)
s="RLRRLLRLRL"
ans=4
print(slt.balancedStringSplit(s)==ans)
s="RLLLLRRRLR"
ans=3
print(slt.balancedStringSplit(s)==ans)
s="LLLLRRRR"
ans=1
print(slt.balancedStringSplit(s)==ans)
