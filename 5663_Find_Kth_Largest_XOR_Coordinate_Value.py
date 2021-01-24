# 2021/1/24
import time
import math
from fractions import Fraction


class Solution:
    def kthLargestValue(self, matrix, k: int) -> int:
        dp = [[0 for i in range(len(matrix[0]))] for i in range(len(matrix))]
        dp[0][0] = matrix[0][0]
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if i == 0 and j == 0:
                    continue
                cur = matrix[i][j]
                if i > 0:
                    cur ^= dp[i-1][j]
                if j > 0:
                    cur ^= dp[i][j-1]
                if i > 0 and j > 0:
                    cur ^= dp[i-1][j-1]
                dp[i][j] = cur
        print(dp)
        l = []
        for i in range(len(dp)):
            for j in range(len(dp[0])):
                l.append(dp[i][j])
        l.sort()
        print(l)
        return l[-k]


start_time = time.time()
slt = Solution()

matrix = [[5, 2], [1, 6]]
k = 1
print(slt.kthLargestValue(matrix, k) == 7)

matrix = [[5, 2], [1, 6]]
k = 2
print(slt.kthLargestValue(matrix, k) == 5)

matrix = [[5, 2], [1, 6]]
k = 3
print(slt.kthLargestValue(matrix, k) == 4)

matrix = [[5, 2], [1, 6]]
k = 4
print(slt.kthLargestValue(matrix, k) == 0)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
