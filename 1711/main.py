from collections import defaultdict
import math

class Solution:
    def countPairs(self, deliciousness: list[int]) -> int:
        del_dic = defaultdict(lambda: 0)
        max_del = 0
        for f in deliciousness:
            max_del = max((max_del, f))
            del_dic[f] += 1
        powers = {1: True}
        i = 1
        while i < max_del:
            i *= 2
            powers[i] = True
        powers[i*2] = True

        keys = list(del_dic.keys())
        res = 0
        for idx, key1 in enumerate(keys):
            same_ones = self.binomial_cooefficient( del_dic[key1], 2)
            if del_dic[key1] > 1 and powers.get(key1, False):
                res += same_ones
            for key2 in keys[idx+1:]:
                if powers.get(key1+key2, False):
                    res += self.binomial_cooefficient(del_dic[key1]+del_dic[key2], 2)-same_ones-self.binomial_cooefficient(del_dic[key2], 2)
        return res

    def binomial_cooefficient(self, n: int, k: int) -> int:
        if n-k < 0:
            return 0
        n_fac = math.factorial(n)
        k_fac = math.factorial(k)
        n_minus_k_fac = math.factorial(n - k)
        return int(n_fac/(k_fac*n_minus_k_fac))
        # max_del = max(deliciousness)
        # powers = {1: True}
        # i = 1
        # while i < max_del:
        #     i = 2*i
        #     powers[i] = True
        # powers[2*i] = True
        # res = 0
        # for idx, food1 in enumerate(deliciousness):
        #     for food2 in deliciousness[idx+1:]:
        #         if powers.get(food1+food2, False):
        #             res += 1
        # return res % (10**9+7)


if __name__ == "__main__":
    deliciousness = [1, 3, 5, 7, 9]
    Solution().countPairs(deliciousness)
