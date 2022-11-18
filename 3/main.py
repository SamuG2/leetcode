class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) < 2:
            return len(s)

        for i in range(min(len(s), 190), 0, -1):
            for j in range(0, min(len(s), 190)-i+1):
                if self.noRepeatingCharacters(s[j:i+j]):
                    return i

    def noRepeatingCharacters(self, s: str) -> bool:
        chars = {}
        for c in s:
            if not chars.get(c, False):
                chars[c] = True
            else:
                return False
        return True


if __name__ == "__main__":
    print(Solution().lengthOfLongestSubstring("bbb"))
