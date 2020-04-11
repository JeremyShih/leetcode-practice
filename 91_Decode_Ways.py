# 2019/9/21
class Solution:
	def numDecodings(self, s: str) -> int:
		if len(s)==0 or s[0]=="0": return 0
		if len(s)==1: return 1
		dp=[0]*len(s)
		dp[0]=1

		if int(s[:2])<27:
			if not s[1]=="0": dp[1]=2
			else: dp[1]=1
		elif not s[1]=="0": dp[1]=1
		
		for i in range(2,len(s)):
			if int(s[i])>0:
				dp[i]+=dp[i-1]
			if 9<int(s[i-1:i+1])<27:
				dp[i]+=dp[i-2]
			if dp[i]==0:
				return 0
		# print(dp)
		return dp[-1]
	def numDecodingsdfs(self, s):
		l=[]
		if s[0]=="0" or int(s[:2])>26 or int(s[:2])==0: return 0
		def dfs(code, s):
			if len(s)==0:
				l.append(code)
			elif len(s)==1:
				l.append(code+chr(int(s)+96))
			else:
				dfs(code+chr(int(s[0])+96), s[1:])
				if int(s[:2])<27:
					dfs(code+chr(int(s[:2])+96), s[2:])
		dfs("", s)
		# print(l)
		return len(l)

slt=Solution()
s="0"
print(slt.numDecodings(s)==0)
s=""
print(slt.numDecodings(s)==0)
s="20"
print(slt.numDecodings(s)==1)
s="2"
print(slt.numDecodings(s)==1)
s="27"
print(slt.numDecodings(s)==1)
s="26"
print(slt.numDecodings(s)==2)
s="80"
print(slt.numDecodings(s)==0)
s="10"
print(slt.numDecodings(s)==1)
s="202"
print(slt.numDecodings(s)==1)
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
# s="226"
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
# s="226378"
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
# s="1238135"
# # print(slt.numDecodings(s))
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
# s="1822181"
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
# s="13061"
# print(slt.numDecodings(s)==slt.numDecodingsdfs(s))
