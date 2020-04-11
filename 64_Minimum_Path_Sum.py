class Solution:
    def minPathSum(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        height=len(grid)
        width=len(grid[0])
        dp=[[0]*width for _ in range(height)]
        dp[0][0]=grid[0][0]
        for r in range(height):
            for c in range(width):
                if r==0 and c==0:
                    continue
                elif r==0:
                    dp[r][c]=grid[r][c]+dp[r][c-1]
                elif c==0:
                    dp[r][c]=grid[r][c]+dp[r-1][c]
                else:
                    dp[r][c]=grid[r][c]+min(dp[r][c-1],dp[r-1][c])
        return dp[-1][-1]

slt=Solution()
inp=[
  [1,2,2],
  [3,8,1],
  [4,2,1]
]
print(slt.minPathSum(inp)==7)
