class Solution:
    def findContentChildren(self, g: list[int], s: list[int]) -> int:
        g.sort()
        s.sort()

        satisfied = 0
        i, j = 0, 0
        while i < len(g) and j < len(s):
            if s[j] < g[i]:
                j += 1
            elif s[j] >= g[i]:
                j += 1
                i += 1
                satisfied += 1
        return satisfied
