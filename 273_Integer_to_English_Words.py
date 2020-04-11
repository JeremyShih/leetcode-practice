# 2019/8/12
class Solution:
	def numberToWords(self, num: int) -> str:
		self.other = ["", "Thousand", "Million", "Billion"]
		if num == 0:
			return "Zero"
		ans = ""
		i = 0
		while num > 0:
			if num % 1000 > 0:
				ans = self.threeDigitNumberToWords(num % 1000) + " " + self.other[i] + " " + ans
			num = num // 1000
			i += 1
		return ans.rstrip()
	def threeDigitNumberToWords(self, num: int) -> str:
		self.singles = { 0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten", 11: "Eleven", 12:"Twelve", 13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18:"Eighteen", 19: "Nineteen"}
		self.double = {20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety"}
		ans = ""
		if num >= 100:
			ans = self.singles[num // 100] + " " + "Hundred "
		num = num % 100
		if num == 0:
			return ans.rstrip()
		if num % 100 < 20 and num % 100 > 0:
			ans += self.singles[num % 100]
		elif num % 10 == 0:
			ans += self.double[num % 100]
		else:
			ans += self.double[num % 100 - num % 10] + " " + self.singles[num % 10]
		return ans

slt=Solution()
n=1
print(slt.numberToWords(n) + "\t" + str(n))
n=9
print(slt.numberToWords(n) + "\t" + str(n))
n=10
print(slt.numberToWords(n) + "\t" + str(n))
n=12
print(slt.numberToWords(n) + "\t" + str(n))
n=20
print(slt.numberToWords(n) + "\t" + str(n))
n=99
print(slt.numberToWords(n) + "\t" + str(n))
n=100
print(slt.numberToWords(n) + "\t" + str(n))
n=111
print(slt.numberToWords(n) + "\t" + str(n))
n=150
print(slt.numberToWords(n) + "\t" + str(n))
n=200
print(slt.numberToWords(n) + "\t" + str(n))
n=999
print(slt.numberToWords(n) + "\t" + str(n))
n=1000
print(slt.numberToWords(n) + "\t" + str(n))
n=1001
print(slt.numberToWords(n) + "\t" + str(n))
n=1010
print(slt.numberToWords(n) + "\t" + str(n))
n=1101
print(slt.numberToWords(n) + "\t" + str(n))
n=100000
print(slt.numberToWords(n) + "\t" + str(n))
n=101000
print(slt.numberToWords(n) + "\t" + str(n))
n=100000000000
print(slt.numberToWords(n) + "\t" + str(n))
