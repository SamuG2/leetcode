class Solution:
    def longestDecomposition(self, text: str) -> int:
        res = 0
        while text:
            found = False
            for i in range(1, len(text)+1//2):
                if text[0:i] == text[-i:len(text)]:
                    text = text[i:-i]
                    found = True
                    break
            if found:
                res += 2
            else:
                res += 1
                break
        return res

    # def possiblePalindrome(self, text: str) -> list[bool, str]:
    #     for i in range(1, len(text)+1//2):
    #         if text[0:i] == text[-i:len(text)]:
    #             text = text[i:-i]
    #             return True, text
    #     return False, text


if __name__ == "__main__":
    #text = "ghiabcdefhelloadamhelloabcdefghi"
    text = "aaa"
    Solution().longestDecomposition(text)
