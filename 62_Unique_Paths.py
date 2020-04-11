class Solution:
    def uniquePaths(self, m, n):
        if m*n==0: return 0
        dp=[1]*n
        for i in range(1,m):
        	for j in range(1,n):
        		dp[j]+=dp[j-1]
        return dp[-1]

slt=Solution()
m, n=3, 2
print(slt.uniquePaths(m,n)==3)
m, n=7, 3
print(slt.uniquePaths(m,n)==28)