class Solution:
    def reversePairs(self, nums):
        if nums==[]: return 0
        class BinaryIndexedTree(object):
            def __init__(self, arg):
                self.counts=[0]*(arg+1)
            def getSum(self, index):
                sum=0
                while index>0:
                    sum+=self.counts[index]
                    index-=index & -index
                return sum
            def update(self, index, delta):
                while index<len(self.counts):
                    self.counts[index]+=delta
                    index+=index & -index
                
        from bisect import bisect_left
        rankMapping={v:i for i,v in enumerate(sorted(set(nums)))}
        ranks=list(sorted(rankMapping.keys()))
        bitTree=BinaryIndexedTree(len(ranks))
        print('rankMapping:',rankMapping)
        print('ranks:',ranks)
        count=0
        for val in reversed(nums):
            rank=rankMapping[val]
            j=bisect_left(ranks,val/2)
            count+=bitTree.getSum(j)
            bitTree.update(rank+1,1)
        return count

class SolutionAns:
    def reversePairs(self, nums):
        # Similar to https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/
        # Use BIT to keep frequency count when reverse iterating
        # answer for each num is prefix sum of freq for 0 to i where i is rank of num//2
        class BinaryIndexedTree:
            def __init__(self, n):
                self._sums = [0] * (n+1)
            def update(self, i, delta):
                while i < len(self._sums):
                    self._sums[i] += delta
                    i += i & -i
            def query(self, i): # get prefixSum
                s = 0
                while i > 0:
                    s += self._sums[i]
                    i -= i & -i
                return s
            
        if not nums: return 0
        rank_mapping = {v: i for i, v in enumerate(sorted(set(nums)))}
        ranks = list(sorted(rank_mapping.keys())) # sorted original number list
        bit_tree = BinaryIndexedTree(len(rank_mapping)) # all 0 at first
        print('ranks=',ranks)
        print('original bit tree=',bit_tree._sums)
        ret = 0
        for val in reversed(nums):
            rank = rank_mapping[val] # the val's rank in original number list
            # find the rank where val/2 will fall in original number list
            idx = bisect.bisect_left(ranks, val/2)
            tmp=bit_tree.query(idx)
            print(val,rank,idx,tmp)
            ret += tmp
            bit_tree.update(rank+1, 1)
            print(bit_tree._sums,[bit_tree.query(i) for i in range(len(bit_tree._sums))])
        print('bit tree after=',bit_tree._sums)
        return ret

sol=Solution()
import time
import bisect
start_time = time.time()
#lst=[5,4,3,2,1]
#lst=[5,3,6,1,9,2,8,1,2]
lst=[233,2000000001,234,2000000006,235,2000000003,236,2000000007,237,2000000002,2000000005,233,233,233,233,233,2000000004]
#cur=[5,5,11,15,24,30,34,34,36]
res=sol.reversePairs(lst)
elapsed_time = time.time() - start_time
print("ans: ",res)
print(elapsed_time)