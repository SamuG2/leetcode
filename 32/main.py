class Solution:
    def longestValidParentheses(self, s: str) -> int:
        stack = [-1]
        max_len = 0
        for i, v in enumerate(s):
            if v == '(':
                stack.append(i)
            else:
                stack.pop()
                if not stack:
                    stack.append(i)
                else:
                    max_len = max(max_len, i-stack[-1])
        return max_len


if __name__ == '__main__':
    # s1 = "))(("
    # s2 = "()"
    # s3 = "((()"
    #s4 = "())()(()"

    # print(Solution().isValid(s1))
    # print(Solution().isValid(s2))
    # print(Solution().isValid(s3))
    # print(Solution().isValid(s4))

    print(Solution().longestValidParentheses("(())()(()(("))
