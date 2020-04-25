from heapq import heappush, heappop, nsmallest
class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.q1, self.q2 = [], []
        
    def get_median(self):
        v1 = -nsmallest(1, self.q1)[0] if self.q1 else None
        v2 =  nsmallest(1, self.q2)[0] if self.q2 else None
        print(v1, v2)
        return v1, v2
    
    def addNum(self, num: int) -> None:
        v1, v2 = self.get_median()
        if len(self.q1) == len(self.q2): 
            if v2 and num > v2: heappop(self.q2); heappush(self.q2, num); heappush(self.q1, -v2)
            else              : heappush(self.q1, -num)
        else:
            if num >= v1: heappush(self.q2, num)
            else: heappop(self.q1); heappush(self.q1, -num); heappush(self.q2, v1); 

    def findMedian(self) -> float:
        v1, v2 = self.get_median()
        if v1 and v2 and len(self.q1) == len(self.q2): return (v1+v2)/2
        elif v1: return v1
        return 0

sol = MedianFinder()
sol.addNum(1)
sol.addNum(2)
print(sol.findMedian()==1.5)
sol.addNum(3)
print(sol.findMedian()==2)
