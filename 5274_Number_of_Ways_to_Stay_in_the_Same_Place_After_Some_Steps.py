# 2019/11/24
class Solution:
	def numWays(self, steps: int, arrLen: int) -> int:
		self.count=0
		def dfs(index, leftStep):
			if leftStep==0 or index>leftStep:
				if index==0:
					self.count+=1
				return
			if index+1<arrLen:
				dfs(index+1,leftStep-1)
			if index-1>-1:
				dfs(index-1,leftStep-1)
			dfs(index,leftStep-1)
		dfs(0,steps)
		return self.count % (pow(10,9)+7)

slt=Solution()
steps = 3
arrLen = 2
a=4
print(slt.numWays(steps,arrLen)==a)
steps = 2
arrLen = 4
a=2
print(slt.numWays(steps,arrLen)==a)
steps = 4
arrLen = 2
a=8
print(slt.numWays(steps,arrLen)==a)
steps = 27
arrLen = 7
a=8
print(slt.numWays(steps,arrLen))
