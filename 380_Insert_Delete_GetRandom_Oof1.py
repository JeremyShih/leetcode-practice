# 2018/11/6 #4
class RandomizedSet:
    def __init__(self):
        self.ds,self.ind=[],{}
    def insert(self, val):
        if val not in self.ind:
            self.ds+=[val]
            self.ind[val]=len(self.ds)-1
            return True
        return False
    def remove(self, val):
        if val in self.ind:
            # copy the last element to the place of the element to be removed
            # and then remove the last element
            ind,last=self.ind[val],self.ds[-1]
            self.ds[ind],self.ind[last]=last,ind
            self.ds.pop()
            self.ind.pop(val)
            return True
        return False
    def getRandom(self):
        import random
        return random.choice(self.ds)


# Your RandomizedSet object will be instantiated and called as such:
obj = RandomizedSet()
val=1
param_1 = obj.insert(val)
obj.insert(2)
val=3
param_2 = obj.remove(val)
param_3 = obj.getRandom()
print(param_3)