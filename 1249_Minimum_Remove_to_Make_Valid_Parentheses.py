# 2019/11/3
class Solution:
	def minRemoveToMakeValid(self, s: str) -> str:
		cur,l=0,[]
		for i in range(len(s)):
			if s[i]=="(":
				cur+=1
				l+=[s[i]]
			elif s[i]==")":
				if cur>0:
					cur-=1
					l+=[s[i]]
			else:
				l+=[s[i]]
		# print(l)
		ans=""
		for i in range(len(l)-1,-1,-1):
			if cur>0 and l[i]=="(":
				cur-=1
			else:
				ans=l[i]+ans
		# print(ans)
		return ans

slt=Solution()
s = "lee(t(c)o)de)"
a="lee(t(c)o)de"
print(slt.minRemoveToMakeValid(s)==a)
s = "a)b(c)d"
a="ab(c)d"
print(slt.minRemoveToMakeValid(s)==a)
s = "))(("
a=""
print(slt.minRemoveToMakeValid(s)==a)
s = "(a(b(c)d)"
a="(a(bc)d)"
print(slt.minRemoveToMakeValid(s)==a)
