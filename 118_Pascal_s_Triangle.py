class Solution:
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        if rowIndex==0: return [1]
        def getN(n, index):
        	return math.factorial(n)/(math.factorial(index)*math.factorial(n-index))
        ans=[]
        for i in range(rowIndex+1):
        	ans.append(int(getN(rowIndex,i)))
        return ans
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        ans=[]
        for i in range(numRows):
        	ans.append(self.getRow(i))
        return ans

import math
slt=Solution()
n=5
print(n, slt.generate(n))
