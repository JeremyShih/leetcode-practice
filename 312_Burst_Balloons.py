# 2020/7/4
import time


class Solution:
    def jobScheduling(self, startTime, endTime, profit) -> int:
        # return dp[-1]
        return 0


slt = Solution()
startTime = [1, 2, 3, 3]
endTime = [3, 4, 5, 6]
profit = [50, 10, 40, 70]
ans = 120
print(slt.jobScheduling(startTime, endTime, profit) == ans)
end = time.time()
print("It cost %f sec" % (end - start))
