# 2019/10/20
class Solution:
	def checkStraightLine(self, coordinates) -> bool:
		coordinates.sort()
		vector=[coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]]
		for i in range(2,len(coordinates)-1):
			dx=(coordinates[i+1][0]-coordinates[i][0])
			dy=(coordinates[i+1][1]-coordinates[i][1])
			if dx*vector[1]==dy*vector[0]:
				continue
			else:
				return False
		return True

slt=Solution()
coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
print(slt.checkStraightLine(coordinates))
coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
print(not slt.checkStraightLine(coordinates))
coordinates = [[-3,-2],[-1,-2],[2,-2],[-2,-2],[0,-2]]
print(slt.checkStraightLine(coordinates))
