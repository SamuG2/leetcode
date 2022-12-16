class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        happy_strings = self.generateStrings(n, "a")+self.generateStrings(n, "b")+self.generateStrings(n, "c")
        if k > len(happy_strings):
            return ""
        happy_strings = sorted(happy_strings)
        return happy_strings[k-1]

    def generateStrings(self, n: int, s: str):
        if len(s) == n:
            return [s]
        if s[-1] == 'a':
            return self.generateStrings(n, s+'b')+self.generateStrings(n, s+'c')
        if s[-1] == 'b':
            return self.generateStrings(n, s+'a')+self.generateStrings(n, s+'c')
        if s[-1] == 'c':
            return self.generateStrings(n, s+'b')+self.generateStrings(n, s+'a')


if __name__ == "__main__":
    n = 3
    k = 9

    Solution().getHappyString(n, k)
