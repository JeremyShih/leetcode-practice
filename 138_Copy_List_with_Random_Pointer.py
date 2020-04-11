# 2019/10/27
"""
# Definition for a Node.
class Node:
	def __init__(self, val, next, random):
		self.val = val
		self.next = next
		self.random = random
"""
class Solution:
	def copyRandomList(self, head: 'Node') -> 'Node':
		if not head:
			return None
		nodeDict={}
		temp=head
		res=new=Node(-1,None,None)
		while temp:
			nodeDict[temp]=Node(temp.val,None,None)
			new.next=nodeDict[temp]
			temp,new=temp.next,new.next
		temp=head
		for keyNode in nodeDict:
			if keyNode.random:
				nodeDict[keyNode].random=nodeDict[keyNode.random]
		return res.next

# {"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
# {"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
