from collections import Counter
class Solution:
	def majorityElement(self, nums):
		if not nums: return []
		dic=Counter(nums)
		#print(dic)
		l=len(nums)/3
		ans=[]
		for key in dic:
			#print(key,dic[key],l)
			if dic[key]>l:
				ans.append(key)
		return ans


slt=Solution()
inp=[3,2,3]
ans=[3]
print(slt.majorityElement(inp)==ans)
inp=[1,1,1,3,3,2,2,2]
ans=[1,2]
print(slt.majorityElement(inp)==ans)
