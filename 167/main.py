class Solution:
    def twoSum(self, numbers: list[int], target: int) -> list[int]:
        m = {}
        for i, n in enumerate(numbers):
            if m.get(n, False):
                return [m[n], i+1]
            m[target-n] = i+1