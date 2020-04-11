# 2019/11/24
class Solution:
	def minTimeToVisitAllPoints(self, points) -> int:
		step=0
		if len(points)==0: return step
		
		for i in range(1,len(points)):
			print(i)
			ver=[abs(points[i][0]-points[i-1][0]),abs(points[i][1]-points[i-1][1])]
			step+=max(ver)
		return step


slt=Solution()
points = [[1,1],[3,4],[-1,0]]
print(slt.minTimeToVisitAllPoints(points)==7)
points = [[3,2],[-2,2]]
print(slt.minTimeToVisitAllPoints(points)==5)
