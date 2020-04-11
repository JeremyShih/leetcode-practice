# Definition for an interval.
# class Interval:
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e

class Solution:
	def merge(self, intervals):
		if not intervals: return []
		intervals.sort(key=lambda x: x.start)
		res=[intervals[0]]
		for i in range(1,len(intervals)):
			if res[-1].end<intervals[i].start:
				res.append(intervals[i])
			elif res[-1].end<intervals[i].end:
				res[-1].end=intervals[i].end
		return res

slt=Solution()
inp=[[1,3],[2,6],[8,10],[15,18]]
ans=[[1,6],[8,10],[15,18]]
print(slt.merge(inp)==ans)
inp=[[1,4],[4,5]]
ans=[[1,5]]
print(slt.merge(inp)==ans)
