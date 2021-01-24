# 2021/1/19
import time
import math
from fractions import Fraction


class Solution:
    def findRepeatedDnaSequences(self, s: str):
        m, ans = {}, []

        for i in range(len(s)-9):
            if s[i:i+10] not in m:
                m[s[i:i+10]] = 0
            m[s[i:i+10]] += 1
        for k in m:
            if m[k] > 1:
                ans.append(k)
        return ans


start_time = time.time()
slt = Solution()

s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
ans = ["AAAAACCCCC", "CCCCCAAAAA"]
print(slt.findRepeatedDnaSequences(s) == ans)

s = "AAAAAAAAAAAAA"
ans = ["AAAAAAAAAA"]
print(slt.findRepeatedDnaSequences(s) == ans)

s = "AAAAAAAAAAA"
ans = ["AAAAAAAAAA"]
print(slt.findRepeatedDnaSequences(s) == ans)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
