from bisect import bisect_left
class Solution:
	def maxEnvelopes(self, envelopes):
		tails = []
		for _,h in sorted(envelopes, key = lambda env : (env[0], -env[1])):
			pos = bisect_left(tails, h)
			if pos == len(tails):
				tails.append(h)
			else:
				tails[pos] = h
			print(tails)
		return len(tails)

slt=Solution()
inp=[[5,4],[6,4],[6,7],[2,3]]
#print(sorted(inp, key = lambda env : (env[0], -env[1])))
print(slt.maxEnvelopes(inp)==3)