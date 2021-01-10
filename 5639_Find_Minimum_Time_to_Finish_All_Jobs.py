# 2021/1/10
import time
import math
from fractions import Fraction


class Solution:
    def minimumTimeRequired(self, jobs, k: int) -> int:
        return 0


start_time = time.time()
slt = Solution()

jobs, k = [3, 2, 3],  3
print(slt.minimumTimeRequired(jobs, k) == 3)

jobs, k = [1, 2, 4, 7, 8],  2
print(slt.minimumTimeRequired(jobs, k) == 11)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
