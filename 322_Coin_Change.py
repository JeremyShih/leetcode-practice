class Solution:
	def coinChange(self, coins, amount):
		coins.sort(reverse=True)
		len_,self.res=len(coins),10**10-1
		def dfs(pt,rem,count):
			#print(rem,count,self.res)
			if not rem:
				self.res=min(self.res,count)
			for i in range(pt,len_):
				#上限是已經計算出的其他組合，如果硬幣數量超過就沒有意義了，所以不繼續計算
				if coins[i]<=rem<coins[i]*(self.res-count):
					dfs(i,rem-coins[i],count+1)
		#for c in range(len_):
		dfs(0,amount,0)
		return self.res if self.res<10**10-1 else -1

slt=Solution()
inp=[1, 2, 5]
s=11
print(slt.coinChange(inp,s)==3)
s=14
print(slt.coinChange(inp,s)==4)
inp=[2]
s=3
print(slt.coinChange(inp,s)==-1)
