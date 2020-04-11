class Solution:
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        def getN(n, index):
            factor=1
            for i in range(index):
                factor *= (n-i)
            return factor/math.factorial(index)
        ans=[]
        for i in range(rowIndex+1):
        	ans.append(int(getN(rowIndex,i)))
        return ans

import math
slt=Solution()
n=3
print(n, slt.getRow(n))
n=4
print(n, slt.getRow(n))