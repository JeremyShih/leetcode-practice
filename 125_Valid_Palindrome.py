# 2018/11/13 #2
class Solution:
	def isPalindrome(self, s):
		if len(s)<=1: return True
		s=s.lower()
		# s = [e for e in s if e.isalnum()]
		i,j=0,len(s)-1
		while i<j:
			if not s[i].isalnum():
				i+=1
			elif not s[j].isalnum():
				j-=1
			elif s[i]==s[j]:
				i,j=i+1,j-1
			else:
				return False
		return True

slt=Solution()
inp="A man, a plan, a canal: Panama"
ans=True
print(slt.isPalindrome(inp)==ans)
inp="race a car"
ans=False
print(slt.isPalindrome(inp)==ans)
