class Solution:
    # def checkCharacter(self, c: str, p: str, last_star: bool) -> int:
    #     if p[0] == "*":
    #         return 1
    #     elif c[0] == p[0] or p[0] == '?':
    #         return 2
    #     return 0

    def isMatch(self, s: str, p: str) -> bool:
        s = list(s)
        p = list(p)
        last_star = False
        while s and p:
            
            if s[0] == p[0] or p[0] == '?':
                p.pop(0)
                s.pop(0)
                
            elif p[0] == "*":
                p.pop(0)
                last_star = True
                # while s:
                #     if self.isMatch(s, p):
                #         return True
                #     else:
                #         s = s[1:]

                # return self.isMatch(s, p) if len(p) != 0 else True
            elif last_star:
                while s and (s[0] != p[0] or s[0]!= '?'):
                    s.pop(0)
                last_star = False
            else:
                return False

        while p and p[0] == '*':
            p = p[1:]
        return len(s) == 0 and len(p) == 0

    # def isMatch(self, s: str, p: str) -> bool:
    #     i = j = 0
    #     back_j = -1
    #     match_i = 0
    #     m = len(s)
    #     n = len(p)
    #     while i < m:
    #         if j < n and (s[i] == p[j] or p[j] == "?"):
    #             i += 1
    #             j += 1
    #         elif j < n and p[j] == "*":
    #             back_j = j
    #             match_i = i
    #         elif back_j != -1:
    #             j = back_j+1
    #             match_i += 1
    #             i = match_i
    #         else:
    #             return False
    #     return list(p[j:]).count('*') == len(p[j:])


if __name__ == "__main__":
    # s = "babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab"
    # p = "***bba**a*bbba**aab**b"

    s = "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba"
    p = "a*******b"
    print(Solution().isMatch(s, p))
