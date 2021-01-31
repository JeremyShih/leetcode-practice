# 2021/1/31
import time
import math
from fractions import Fraction
from typing import List


class Solution:
    def restoreArray(self, adjacentPairs: List[List[int]]) -> List[int]:
        if len(adjacentPairs) == 1:
            return adjacentPairs[0]
        pairs, count, head = {}, {}, 0
        for p in adjacentPairs:
            if p[0] not in count:
                count[p[0]] = 0
            if p[1] not in count:
                count[p[1]] = 0
            count[p[0]] += 1
            count[p[1]] += 1

            if p[0] not in pairs:
                pairs[p[0]] = []
            if p[1] not in pairs:
                pairs[p[1]] = []
            pairs[p[0]] += [p[1]]
            pairs[p[1]] += [p[0]]

        for k in count:
            if count[k] == 1:
                head = k
                break

        ans = [head]
        # print(count)
        # print(pairs)
        for i in range(len(adjacentPairs)):
            last = ans[-1]
            # print(ans)
            if len(pairs[last]) == 1:
                ans.append(pairs[ans[-1]][0])
            else:
                if pairs[last][0] == ans[-2]:
                    ans.append(pairs[ans[-1]][1])
                else:
                    ans.append(pairs[ans[-1]][0])
        # print(ans)
        return ans


start_time = time.time()
slt = Solution()

adjacentPairs = [[2, 1], [3, 4], [3, 2]]
ans = [1, 2, 3, 4]
print(slt.restoreArray(adjacentPairs) == ans)

adjacentPairs = [[4, -2], [1, 4], [-3, 1]]
ans = [-2, 4, 1, -3]
print(slt.restoreArray(adjacentPairs) == ans)

adjacentPairs = [[100000, -100000]]
ans = [100000, -100000]
print(slt.restoreArray(adjacentPairs) == ans)

adjacentPairs = [[4, -10], [-1, 3], [4, -3], [-3, 3]]
print(slt.restoreArray(adjacentPairs))

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
