class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid):
        if obstacleGrid[-1][-1]==1 or obstacleGrid[0][0]==1:
            return 0
        inp=[[-j for j in i] for i in obstacleGrid]
        for c in range(len(inp[0])):
            for r in range(len(inp)):
                if inp[r][c]==-1:
                    continue
                elif c==0 and r==0:
                    inp[r][c]=1
                elif c==0:
                    if inp[r-1][c]==-1:
                        inp[r][c]=0
                    else:
                        inp[r][c]=inp[r-1][c]
                elif r==0:
                    if inp[r][c-1]==-1:
                        inp[r][c]=0
                    else:
                        inp[r][c]=inp[r][c-1]
                elif inp[r-1][c]==-1 and inp[r][c-1]==-1:
                    inp[r][c]=0
                elif inp[r-1][c]==-1:
                    inp[r][c]=inp[r][c-1]
                elif inp[r][c-1]==-1:
                    inp[r][c]=inp[r-1][c]
                else:
                    inp[r][c]=inp[r-1][c]+inp[r][c-1]
        return inp[-1][-1]

slt=Solution()
inpu=[[1]]
print(slt.uniquePathsWithObstacles(inpu)==0)
inpu=[[0]]
print(slt.uniquePathsWithObstacles(inpu)==1)
inpu=[[0,1]]
print(slt.uniquePathsWithObstacles(inpu)==0)
inpu=[[0,0],[0,1]]
print(slt.uniquePathsWithObstacles(inpu)==0)
inpu=[[1,0],[0,0]]
print(slt.uniquePathsWithObstacles(inpu)==0)
inpu=[[0,1],[1,0]]
print(slt.uniquePathsWithObstacles(inpu)==0)
inpu=[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
print(slt.uniquePathsWithObstacles(inpu)==2)
inpu=[[0,1,0,0,0],[1,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0]]
print(slt.uniquePathsWithObstacles(inpu)==0)
