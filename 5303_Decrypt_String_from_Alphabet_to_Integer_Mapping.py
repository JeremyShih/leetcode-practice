# 2020/1/5
class Solution:
	def freqAlphabets(self, s: str) -> str:
		d1={"1": "a", "2": "b", "3": "c", "4": "d", "5": "e", "6": "f", "7": "g", "8": "h", "9": "i"}
		d2={"10#": "j", "11#": "k", "12#": "l", "13#": "m", "14#": "n", "15#": "o", "16#": "p", "17#": "q", "18#": "r", "19#": "s", "20#": "t", "21#": "u", "22#": "v", "23#": "w", "24#": "x", "25#": "y", "26#": "z"}
		for k in d2:
			s=s.replace(k,d2[k])
		for k in d1:
			s=s.replace(k,d1[k])
		return s

slt=Solution()
s = "10#11#12"
a="jkab"
print(slt.freqAlphabets(s)==a)
s = "1326#"
a="acz"
print(slt.freqAlphabets(s)==a)
s = "25#"
a="y"
print(slt.freqAlphabets(s)==a)
s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
a="abcdefghijklmnopqrstuvwxyz"
print(slt.freqAlphabets(s)==a)
