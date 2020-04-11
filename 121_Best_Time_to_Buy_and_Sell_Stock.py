# 2019/9/22
class Solution:
	def maxProfit(self, prices: list) -> int:
		buy,profit=float('inf'),0

		for p in prices:
			if p>buy:
				profit=max(profit, p-buy)
			else:
				buy=p

		return profit

slt=Solution()
l=[7,1,5,3,6,4]
print(slt.maxProfit(l)==5)
l=[]
print(slt.maxProfit(l)==0)
l=[2]
print(slt.maxProfit(l)==0)
l=[7,6,4,3,1]
print(slt.maxProfit(l)==0)
l=[7,6,4,3,5]
print(slt.maxProfit(l)==2)
