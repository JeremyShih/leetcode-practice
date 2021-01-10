# 2021/1/9
import time
import math
from fractions import Fraction


class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        # test
        return 0


start_time = time.time()
slt = Solution()
# 1
version1 = "1.01"
version2 = "1.001"
print(slt.compareVersion(version1, version2) == 0)
# 2
version1 = "1.0"
version2 = "1.0.0"
print(slt.compareVersion(version1, version2) == 0)
# 3
version1 = "0.1"
version2 = "1.1"
print(slt.compareVersion(version1, version2) == -1)
# 4
version1 = "1.01"
version2 = "1.001"
print(slt.compareVersion(version1, version2) == 0)
# 5
version1 = "1.01"
version2 = "1.001"
print(slt.compareVersion(version1, version2) == 0)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
