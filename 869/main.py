from itertools import permutations


class Solution:
    def reorderedPowerOf2(self, n: int) -> bool:
        powers = {"1":True}
        i =1
        while i<10**9:
            i*=2
            powers[''.join(sorted(str(i)))] = True
        return powers.get(''.join(sorted(str(n))), False)
        
        # n = list(str(n))
        # aux = permutations(n, len(n))
        # combs = []
        # max_num = 0
        # for c in aux:
        #     if c[0] != '0':
        #         num = int(''.join(c))
        #         max_num = max(max_num, num)
        #         combs.append(num)
        # i = 1
        # powers = {1: True}
        # while i < max_num:
        #     i *= 2
        #     powers[i] = True

        # for c in combs:
        #     if powers.get(c):
        #         return True
        # return False


if __name__ == "__main__":
    n = 16
    Solution().reorderedPowerOf2(n)
