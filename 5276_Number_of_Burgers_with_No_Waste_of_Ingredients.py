# 2019/12/1
# 二元一次方程式？
# 4a + 2b = tomatoSlices
# a + b = cheeseSlices
class Solution:
	def numOfBurgers(self, tomatoSlices: int, cheeseSlices: int):
		ans=[]
		if tomatoSlices>cheeseSlices*4 or tomatoSlices<cheeseSlices*2:
			return ans
		if (tomatoSlices-cheeseSlices*2) % 2 != 0 and (tomatoSlices-cheeseSlices*2) > 0:
			# print("hit")
			return ans
		j=(tomatoSlices-cheeseSlices*2) // 2
		s=cheeseSlices-j
		return [j,s]

slt=Solution()
tomatoSlices = 16
cheeseSlices = 7
a=[1,6]
print(slt.numOfBurgers(tomatoSlices,cheeseSlices)==a)
tomatoSlices = 17
cheeseSlices = 4
a=[]
print(slt.numOfBurgers(tomatoSlices,cheeseSlices)==a)
tomatoSlices = 4
cheeseSlices = 17
a=[]
print(slt.numOfBurgers(tomatoSlices,cheeseSlices)==a)
tomatoSlices = 0
cheeseSlices = 0
a=[0,0]
print(slt.numOfBurgers(tomatoSlices,cheeseSlices)==a)
tomatoSlices = 2
cheeseSlices = 1
a=[0,1]
print(slt.numOfBurgers(tomatoSlices,cheeseSlices)==a)
tomatoSlices = 1511274
cheeseSlices = 594732
print(slt.numOfBurgers(tomatoSlices,cheeseSlices))
tomatoSlices,cheeseSlices = 2537427,860448
print(slt.numOfBurgers(tomatoSlices,cheeseSlices))
