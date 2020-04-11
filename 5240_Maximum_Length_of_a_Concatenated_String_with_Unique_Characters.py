# 2019/10/27
class Solution:
	def maxLength(self, arr) -> int:
		l=[]
		def unique(s):
			for i in range(len(s)):
				for j in range(i+1,len(s)):
					if s[i]==s[j]:
						return False
			return True
		for s in arr:
			if unique(s):
				for t in l:
					if unique(t+s):
						l.append(t+s)
				l.append(s)

		ans=0
		# print(l)
		for s in l:
			ans=max(ans,len(s))
		return ans

slt=Solution()
arr = ["un","iq","ue"]
a=4
print(slt.maxLength(arr)==a)
arr = ["cha","r","act","ers"]
a=6
print(slt.maxLength(arr)==a)
arr = ["abcdefghijklmnopqrstuvwxyz"]
a=26
print(slt.maxLength(arr)==a)
arr=["ab","ba","cd","dc","ef","fe","gh","hg","ij","ji","kl","lk","mn","nm","op","po"]
print(slt.maxLength(arr))
arr=["yy","bkhwmpbiisbldzknpm"]
print(slt.maxLength(arr))
