# 2021/1/24
import time
import math
from fractions import Fraction


class Solution:
    def maximumTime(self, time: str) -> str:
        h, m = time.split(':')
        ans = ''
        if h[0] == '?':
            if h[1] != '?' and int(h[1]) < 4:
                ans = '2'
            elif h[1] == '?':
                ans = '2'
            else:
                ans = '1'
        else:
            ans = h[0]

        if h[1] == '?':
            if ans[0] == '2':
                ans += '3'
            else:
                ans += '9'
        else:
            ans += h[1]
        ans += ':'
        if m[0] == '?':
            ans += '5'
        else:
            ans += m[0]
        if m[1] == '?':
            ans += '9'
        else:
            ans += m[1]

        return ans


start_time = time.time()
slt = Solution()

t = "2?:?0"
print(slt.maximumTime(t) == "23:50")

t = "0?:3?"
print(slt.maximumTime(t) == "09:39")

t = "1?:22"
print(slt.maximumTime(t) == "19:22")

t = "?4:03"
print(slt.maximumTime(t) == "14:03")

t = "??:3?"
print(slt.maximumTime(t) == "23:39")

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
