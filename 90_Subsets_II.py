import collections
class Solution:
	def subsetsWithDup(self, nums):
		if not nums:
			return [[]]
		res = [[]]
		dic = collections.Counter(nums)
		for key, val in dic.items():
			tmp = []
			for lst in res:
				for i in range(1, val+1):
					tmp.append(lst+[key]*i)
			res += tmp
		return res

slt=Solution()
inp=[1,2,3]
ans=[
  [3],
  [1],
  [2],
  [1,3],
  [2,3],
  [1,2],
  [1,2,3],
  []
]
print(slt.subsetsWithDup(inp))
inp=[1,2,2]
ans=[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
print(slt.subsetsWithDup(inp))
