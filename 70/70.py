class Solution:
    def climbStairs(self, n: int) -> int:
        solutions = {1: 1, 2: 2}
        for i in range(3, n):
            solutions[i] = solutions[i-1]+solutions[i-2]
        return solutions[n-1]+solutions[n-2] if n > 2 else solutions[n]
