class NumArray:
    sum_=[]
    len_=0
    nums_=[]
    def __init__(self, nums):
        """
        :type nums: List[int]
        """
        self.nums_=[0]*len(nums)
        self.len_=len(nums)+1
        #print('len:',self.len_)
        self.sum_=[0]*self.len_

        for i in range(len(nums)):
            #print(i)
            self.update(i, nums[i])

    def update(self, i, val):
        """
        :type i: int
        :type val: int
        :rtype: void
        """
        dif=val-self.nums_[i]
        self.nums_[i]=val
        i+=1
        while i<self.len_:
            self.sum_[i]+=dif
            #print(i)
            i+=i&-i
            
    def getSum(self, i):
        ret=0
        i+=1
        while i>0:
            ret+=self.sum_[i]
            i-=i&-i
        return ret

    def sumRange(self, i, j):
        """
        :type i: int
        :type j: int
        :rtype: int
        """
        return self.getSum(j)-self.getSum(i-1)


# Your NumArray object will be instantiated and called as such:
nums=[1, 3, 5]
i_=1
val=2
obj = NumArray(nums)
#print(obj.sum_)
#print(obj.getSum(-1))
#print(obj.getSum(2))
print(obj.sumRange(0,2))
obj.update(1,2)
print(obj.sum_)
param_2 = obj.sumRange(0,2)
print(param_2)