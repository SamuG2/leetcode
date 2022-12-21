class Solution:
    def isHappy(self, n: int) -> bool:
        visited = {}
        n = str(n)
        while n != "1":
            if visited.get(n, False):
                return False
            visited[n] = True
            sum = 0
            for digit in n:
                sum += int(digit)**2
            n = str(sum)
        return True
