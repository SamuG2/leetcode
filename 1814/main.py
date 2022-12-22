class Solution:
    def countNicePairs(self, nums: list[int]) -> int:
        seen = {}
        seen_revs = {}
        for i
    
    def revNum(self, n: int) -> int:
        s = str(n)
        list(s).reverse()
        n = ''.join(s)
        return int(n)