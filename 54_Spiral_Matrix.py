#2018/11/5 倒數第二題
class Solution:
	def spiralOrder(self, matrix):
		if not matrix: return []
		res=[]
		rowBegin,colBegin,rowEnd,colEnd=0,0,len(matrix)-1,len(matrix[0])-1
		while rowBegin<=rowEnd and colBegin<=colEnd:
			for i in range(colBegin,colEnd+1):
				res.append(matrix[rowBegin][i])
			rowBegin+=1
			for i in range(rowBegin,rowEnd+1):
				res.append(matrix[i][colEnd])
			colEnd-=1
			if rowBegin<=rowEnd:
				for i in range(colEnd,colBegin-1,-1):
					res.append(matrix[rowEnd][i])
			rowEnd-=1
			if colBegin<=colEnd:
				for i in range(rowEnd,rowBegin-1,-1):
					res.append(matrix[i][colBegin])
			colBegin+=1
		return res

slt=Solution()
inp=[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
ans=[1,2,3,6,9,8,7,4,5]
print(slt.spiralOrder(inp)==ans)
inp=[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
ans=[1,2,3,4,8,12,11,10,9,5,6,7]
print(slt.spiralOrder(inp)==ans)
