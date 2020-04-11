# 2019/8/17
class Solution:
	def letterCombinations(self, digits: str):
		dic = {"2":["a","b","c"],"3":["d","e","f"],"4":["g","h","i"],"5":["j","k","l"],"6":["m","n","o"],"7":["p","q","r","s"],"8":["t","u","v"],"9":["w","x","y","z"]}
		def dfs(digitList,res,cur):
			if not digitList:
				res.append(cur)
				return
			for let in dic[digitList[0]]:
				dfs(digitList[1:],res,cur+let)
		ans=[]
		if not digits:
			return ans
		dfs(digits,ans,"")
		return ans

slt=Solution()
n="234"
print(slt.letterCombinations(n))