# 2019/11/2
class LRUCache:

	def __init__(self, capacity):
		# Current time
		self.t = -1
		# Capacity
		self.c = capacity
		# Dictionary key->[value, time of last access]
		self.d = {}
		# List of accessed keys (t -> key accessed at time t)
		self.l = []
		# Relevance pointer to l: all elements in l[:p] are not relevant anymore
		self.p = 0
		

	def get(self, key):
		if key not in self.d:
			return -1
		self.t+=1
		# Getting value
		v = self.d[key][0]
		# Updating time
		self.d[key][1] = self.t
		# Updating access history
		self.l.append(key)
		return v

	def put(self, key, value):
		self.t+=1
		# If we need to delete something...
		if key not in self.d and len(self.d)==self.c:
			# We check the element accessed at time p to see if it was accessed again afterwards. If not, the key is deleted. If yes, keep searching.
			while self.l[self.p] not in self.d or self.d[self.l[self.p]][1]!=self.p:
				print("self.p+=1")
				self.p+=1
			del self.d[self.l[self.p]]
			  
		self.d[key] = [value,self.t]
		self.l.append(key)


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)