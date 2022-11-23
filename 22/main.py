class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        parenthesis =["()"]
        
        for _ in range(1, n):
            next = set()
            for s in parenthesis:
                for i in range(len(s)):
                    next.add(s[:i]+"()"+s[i:])
            parenthesis = next
        return parenthesis
    
    

if __name__ == "__main__":
    print(Solution().generateParenthesis(4))

