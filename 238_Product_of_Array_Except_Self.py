# 2018/11/5 最後一題
class Solution:
	def productExceptSelf(self, nums):
		res=[1]*len(nums)
		tmp=[1]*len(nums)
		for i in range(1,len(nums)):
			res[i]=res[i-1]*nums[i-1]
			tmp[-i-1]=tmp[-i]*nums[-i]
		# print(res)
		# tmp[-1]=nums[-1]
		# for i in range(len(nums)-2,-1,-1):
		# 	tmp[i]=tmp[i+1]*nums[i+1]
		# print(tmp)
		for i in range(len(nums)):
			res[i]*=tmp[i]
		return res

slt=Solution()
inp=[1,2,3,4]
ans=[24,12,8,6]
print(slt.productExceptSelf(inp)==ans)
inp=[1,2,3,4,1]
ans=[24,12,8,6,24]
print(slt.productExceptSelf(inp)==ans)
inp=[5,2,3,4]
ans=[24,60,40,30]
print(slt.productExceptSelf(inp)==ans)
