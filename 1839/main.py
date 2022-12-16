class Solution:
    def longestBeautifulSubstring(self, word: str) -> int:
        chars = " aeiou "
        longest, cur_word = 0, 0
        cur_char = 0

        for c in word:
            if c == chars[cur_char]:
                cur_word += 1
            elif c == chars[cur_char+1]:
                cur_word += 1
                cur_char += 1
            else:
                if chars[cur_char] == 'u':
                    longest = max(longest, cur_word)
                if c == 'a':
                    cur_word = 1
                    cur_char = 1
                else:
                    cur_word = 0
                    cur_char = 0
        return max(longest, cur_word) if chars[cur_char] == 'u' else longest


if __name__ == "__main__":
    word = "auoeioueiaaioeuieuoaieuaoeuoaiaoueioiaeuiuaeouaoie"
    Solution().longestBeautifulSubstring(word)
