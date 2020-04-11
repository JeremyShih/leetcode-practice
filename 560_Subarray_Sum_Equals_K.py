# 2018/11/11 #8
class Solution:
    def subarraySum(self, nums, target):
        dic = {0:1}
        res = pre_sum = 0
        for num in nums:
            pre_sum += num
            res += dic.get(pre_sum - target, 0)
            dic[pre_sum] = dic.get(pre_sum, 0) + 1
        return res