# 2019/12/8
class Solution:
	def groupThePeople(self, groupSizes):
		d={}
		for i in range(len(groupSizes)):
			if groupSizes[i] not in d:
				d[groupSizes[i]]=[i]
			else:
				d[groupSizes[i]].append(i)
		ans=[]
		for k in d:
			if len(d[k])==k:
				ans.append(d[k])
			else:
				for i in range(k,len(d[k])+1,k):
					# print(i,d[k])
					ans.append(d[k][i-k:i])
		return ans

slt=Solution()
groupSizes = [3,3,3,3,3,1,3]
print(slt.groupThePeople(groupSizes))
groupSizes = [2,1,3,3,3,2]
print(slt.groupThePeople(groupSizes))
