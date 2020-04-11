class Solution:
    def isPowerOfFour(self, num):
        """
        :type num: int
        :rtype: bool
        """
        i=1
        while i<=num:
            if i==num: return True
            i*=4
        return False

sol=Solution()
n=32
ans=sol.isPowerOfFour(n)
print(n,ans)