# 2019/11/1
class Solution:
    def convertToBase7(self, num: int) -> str:
        if num==0: return "0"
        ans, ori="", num
        num=abs(num)
        while (num>0):
            ans=str(num%7)+ans
            num=num//7
        return ans if ori>0 else "-"+ans

slt=Solution()
n=100
a="202"
print(slt.convertToBase7(n)==a)
n=-7
a="-10"
print(slt.convertToBase7(n)==a)
