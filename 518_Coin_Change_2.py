class Solution:
	def change(self, amount, coins):
		record = [1]+[0]*(amount)
		coins=[0]+coins
		for i in range(1,len(coins)):
			#針對每個硬幣逐一計算，組成和為val的方法
			#組成和為val的方法數=val把各種硬幣減掉一個的方法數加總
			for val in range(coins[i],amount+1):
				#coins[i]相當於爬樓梯的最小跨距
				record[val]+=record[val-coins[i]]
				print(record)
		#print(record)
		return record[amount]

slt=Solution()
inp=[1, 2, 5]
s=5
print(slt.change(s,inp)==4)
inp=[10]
s=10
print(slt.change(s,inp)==1)
inp=[2]
s=3
print(slt.change(s,inp)==0)
s=500
inp=[3,5,7,8,9,10,11]
#print(slt.change(s,inp)==35502874)
