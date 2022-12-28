class Solution:
    def mergeTriplets(self, triplets: list[list[int]], target: list[int]) -> bool:
        a, b, c = False, False, False
        for t in triplets:
            if t[0] == target[0] and t[1] <= target[1] and t[2] <= target[2]:
                a = True
            if t[1] == target[1] and t[0] <= target[0] and t[2] <= target[2]:
                b = True
            if t[2] == target[2] and t[1] <= target[1] and t[0] <= target[0]:
                c = True
            if a and b and c:
                return True
        return False