class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        if len(s) % 2 == 1:
            return False
        open_r, closed_r, variables_r, open_l, closed_l, variables_l = 0, 0, 0, 0, 0, 0

        for i in range(len(s)):
            if locked[i] == '0':
                variables_r += 1
            elif s[i] == "(":
                open_r += 1
            elif s[i] == ")":
                closed_r += 1
            if variables_r + open_r - closed_r < 0:
                return False

            if locked[-(i+1)] == '0':
                variables_l += 1
            elif s[-(i+1)] == "(":
                open_l += 1
            elif s[-(i+1)] == ")":
                closed_l += 1
            if variables_l - open_l + closed_l < 0:
                return False

        return True

if __name__ == '__main__':
    print(Solution().canBeValid("()", "11"))
