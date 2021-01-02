# 2020/11/8
import time
import math
from fractions import Fraction


class Solution:
    def maxProfit(self, inventory, orders: int) -> int:
        mod = math.pow(10, 9)+7
        val, l = 0, len(inventory)
        inventory.sort(reverse=True)
        # print(inventory)
        while orders > 0:
            print(inventory, "orders:", orders, "val:", val)
            i = 0
            while i+1 < l and inventory[i] <= inventory[i+1]:
                i += 1
            print("i", i)
            t = inventory[i]-1
            if i+1 < l:
                t = inventory[i+1]
            if l == 1 or orders < inventory[i] - t:
                t = inventory[i] - orders

            print("invent:", inventory[i], "\ntarget:", t, "\norders:", orders)
            v = inventory[i]
            diff = Fraction(v+t+1) * Fraction(v-t)/Fraction(2)
            print("differ:", diff)
            diff = diff % Fraction(mod)
            val += diff
            val %= mod
            orders -= v - t
            inventory[i] = t

        # print(inventory, "orders:", orders, "val:", val)
        return val


start_time = time.time()
slt = Solution()

inv, ors = [2, 5], 4
print(slt.maxProfit(inv, ors) == 14)
inv, ors = [3, 5], 6
print(slt.maxProfit(inv, ors) == 19)
inv, ors = [2, 8, 4, 10, 6], 20
print(slt.maxProfit(inv, ors) == 110)
inv, ors = [1000000000], 1000000000
print(slt.maxProfit(inv, ors) == 21)
inv, ors = [76], 22
print(slt.maxProfit(inv, ors) == 1441)
inv, ors = [773160767], 252264991
print(slt.maxProfit(inv, ors) == 70267492)
inv, ors = [497978859, 167261111, 483575207, 591815159], 836556809
print(slt.maxProfit(inv, ors))
# print(slt.maxProfit(inv, ors) == 373219333)

elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
