from collections import defaultdict


class Solution:
    def frequencySort(self, nums: list[int]) -> list[int]:
        d = defaultdict(lambda : 0)
        for n in nums:
            d[n] += 1
        nums.sort(key=lambda x: d[x], reverse=True)
