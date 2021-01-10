# 2021/1/9
import time
import math
from fractions import Fraction


class Solution:
    def reverseWords(self, s: str) -> str:
        splitStr = s.split()
        splitStr.reverse()
        return " ".join(splitStr)


start_time = time.time()
slt = Solution()

s = "the sky is blue"
ans = "blue is sky the"
print(slt.reverseWords(s) == ans)

s = "  hello world  "
ans = "world hello"
print(slt.reverseWords(s) == ans)

s = "a good   example"
ans = "example good a"
print(slt.reverseWords(s) == ans)

s = "  Bob    Loves  Alice   "
ans = "Alice Loves Bob"
print(slt.reverseWords(s) == ans)

s = "Alice does not even like bob"
ans = "bob like even not does Alice"
print(slt.reverseWords(s) == ans)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
