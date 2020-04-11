# 2019/10/13
import collections 
class Solution(object):
	def ladderLength(self, beginWord, endWord, wordList):
		adjacency = collections.defaultdict(list)
		for word in wordList:
			for i in range(len(word)):
				adjacency[word[:i]+'*'+word[i+1:]].append(word)

		seen = set()
		q = collections.deque([(beginWord, 1)])
		while q:
			word, steps = q.popleft()
			if word==endWord: return steps

			for i in range(len(word)):
				for next_word in adjacency[word[:i]+'*'+word[i+1:]]:
					if next_word not in seen:
						q.append((next_word, steps+1))
						seen.add(next_word)
		return 0

slt=Solution()
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]
ans=5
print(slt.ladderLength(beginWord,endWord,wordList)==ans)
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]
ans=0
print(slt.ladderLength(beginWord,endWord,wordList)==ans)
