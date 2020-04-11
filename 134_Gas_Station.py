# 2019/10/27
# 1. if sum(gas) > sum(cost) --> there must be a route for car to go through all the station.
# 2. start from index == 0 and if the total passed gas amount is smaller than the total cost amount at any index n, the car cannot go anymore.
# --> the start should be at some index after n.
class Solution:
	def canCompleteCircuit(self, gas, cost) -> int:
		if sum(gas)<sum(cost):
			return -1
		ans,tank=0,0
		for i in range(len(gas)):
			left=gas[i]-cost[i]
			if tank+left<0:
				ans,tank=i+1,0
			else:
				tank+=left
		return ans

slt=Solution()
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
a=3
print(slt.canCompleteCircuit(gas,cost)==a)
gas  = [2,3,4]
cost = [3,4,3]
a=-1
print(slt.canCompleteCircuit(gas,cost)==a)
gas=[67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66]
cost=[27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26]
a=74
print(slt.canCompleteCircuit(gas,cost))
