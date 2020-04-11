# 2019/10/13
class Solution:
	def queensAttacktheKing(self, queens, king):
		direc,ans=[[0,1],[1,0],[0,-1],[-1,0],[1,1],[1,-1],[-1,1],[-1,-1]],[]
		for vec in direc:
			lattice=[]
			lattice.append(king[0])
			lattice.append(king[1])
			while 0<=lattice[0]<=8 and 0<=lattice[1]<=8:
				lattice[0]+=vec[0]
				lattice[1]+=vec[1]
				if lattice in queens:
					ans.append(lattice)
					break
		return ans

slt=Solution()
queens,king = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], [0,0]
ans=[[0,1],[1,0],[3,3]]
# print(slt.queensAttacktheKing(queens,king))
print(slt.queensAttacktheKing(queens,king)==ans)
queens,king = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], [3,3]
ans=[[2,2],[3,4],[4,4]]
print(len(slt.queensAttacktheKing(queens,king))==len(ans))
# print(slt.queensAttacktheKing(queens,king)==ans)
queens,king = [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]], [3,4]
ans=[[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]
print(len(slt.queensAttacktheKing(queens,king))==len(ans))
# print(slt.queensAttacktheKing(queens,king)==ans)
