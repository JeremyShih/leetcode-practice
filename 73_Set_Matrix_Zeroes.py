class Solution:
    def setZeroes(self, matrix):
        rows,cols=[],[]
        for row in range(len(matrix)):
        	for col in range(len(matrix[row])):
        		if matrix[row][col]==0:
        			rows.append(row)
        			cols.append(col)
        for r in rows:
        	for i in range(len(matrix[r])):
        		matrix[r][i]=0
        for c in cols:
        	for i in range(len(matrix)):
        		matrix[i][c]=0


sol=Solution()
inp=[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
print(inp)
import time
start_time = time.time()
sol.setZeroes(inp)
print(inp)
elapsed_time = time.time() - start_time
print('timp spent:',elapsed_time)
