# 2019/10/20
class Solution:
	def removeSubfolders(self, folder):
		dic={}
		folder.sort()
		i,count=0,len(folder)
		while i<count:
			j,key=i+1,folder[i]+"/"
			while j<count:
				if key in folder[j] and folder[j].index(key)==0:
					folder.pop(j)
					count-=1
				else:
					j+=1
			i+=1
		return folder

slt=Solution()
folder = ["/a","/a/b","/c/d","/c/d/e","/c/d/f","/c/f"]
ans = ["/a","/c/d","/c/f"]
# print(slt.removeSubfolders(folder))
print(slt.removeSubfolders(folder)==ans)
folder = ["/a","/a/b/c","/a/b/d"]
ans = ["/a"]
print(slt.removeSubfolders(folder)==ans)
folder = ["/a/b/c","/a/b/ca","/a/b/d"]
ans = ["/a/b/c","/a/b/ca","/a/b/d"]
# print(slt.removeSubfolders(folder))
print(slt.removeSubfolders(folder)==ans)
