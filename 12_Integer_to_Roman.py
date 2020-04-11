# 2018/11/15 #4
class Solution:
	def intToRoman(self, num):
		numericals=[1000,900,500,400,100,90,50,40,10,9,5,4,1]
		roman=['M','CM','D','CD','C','XC','L','XL','X','IX','V','IV','I']
		i=0
		s=""
		while 0<num:
			if (num-numericals[i])>=0:
				s+=roman[i]
				num-=numericals[i]
			else:
				i+=1
			# if 1000<=num:
			# 	s+="M"*(num//1000)
			# 	num%=1000
			# elif 900<=num:
			# 	s+="CM"
			# 	num%=900
			# elif 500<=num:
			# 	s+="D"+"C"*((num-500)//100)
			# 	num%=100
			# elif 400<=num:
			# 	s+="CD"
			# 	num%=400
			# elif 100<=num:
			# 	s+="C"*(num//100)
			# 	num%=100
			# elif 90<=num:
			# 	s+="XC"
			# 	num%=90
			# elif 50<=num:
			# 	s+="L"+"X"*((num-50)//10)
			# 	num%=10
			# elif 40<=num:
			# 	s+="XL"
			# 	num%=40
			# elif 10<=num:
			# 	s+="X"*(num//10)
			# 	num%=10
			# elif 9<=num:
			# 	s+="IX"
			# 	num%=9
			# elif 5<=num:
			# 	s+="V"+"I"*(num-5)
			# 	num%=1
			# elif 4<=num:
			# 	s+="IV"
			# 	num%=1
			# else:
			# 	s+="I"*num
			# 	num%=1
		return s

slt=Solution()
inp=3
ans="III"
print(slt.intToRoman(inp)==ans)
inp=4
ans="IV"
print(slt.intToRoman(inp)==ans)
inp=9
ans="IX"
print(slt.intToRoman(inp)==ans)
inp=58
ans="LVIII"
print(slt.intToRoman(inp)==ans)
inp=1994
ans="MCMXCIV"
print(slt.intToRoman(inp)==ans)
