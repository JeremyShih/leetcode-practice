# 2019/11/3
class Solution:
	def minimumSwap(self, s1: str, s2: str) -> int:
		if (s1.count("x")+s2.count("x"))%2!=0:
			return -1
		t,i=0,0
		while i<len(s1):
			if s1[i]==s2[i]:
				s1=s1[:i]+s1[i+1:]
				s2=s2[:i]+s2[i+1:]
			else:
				i+=1
		# print(s1)
		# print(s2)

		i,j=s1.count("x"),s2.count("x")
		if i==j:
			if i%2==0:
				return i
			else:
				return i+1
		else:
			if i%2==0:
				return int((i+j)/2)
			else:
				return int((i+j)/2)+1

slt=Solution()
s1 = "xx"
s2 = "xy"
a=-1
# print(slt.minimumSwap(s1,s2)==a)
s1 = "xxyyxyxyxx"
s2 = "xyyxyxxxyx"
a=4
print(slt.minimumSwap(s1,s2)==a)
# print(slt.minimumSwap(s1,s2))
s1 = "xx"
s2 = "yy"
a=1
print(slt.minimumSwap(s1,s2))
