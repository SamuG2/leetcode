class Solution:
    def countAndSay(self, n: int) -> str:
        r = "1"
        for _ in range(1, n):
            r = self._countAndSay(r)
            # l = self.count(r)
            # r = self.say(l)
        return r

    def _countAndSay(self, s: str) -> str:
        res = ""
        cur = [s[0]]
        for n in s[1:]:
            if cur[-1] == n:
                cur.append(n)
            else:
                res += str(len(cur))+str(cur[0])
                cur = [n]
        res += str(len(cur))+str(cur[0])
        return res

    # def count(self, s: str) -> list[list[int]]:
    #     res = []
    #     cur = []

    #     for n in s:
    #         if not cur or cur[-1] == n:
    #             cur.append(n)
    #         else:
    #             res.append([len(cur), cur[0]])
    #             cur = [n]

    #     res.append([len(cur), cur[0]])
    #     return res

    # def say(self, l: list[list[int]]) -> str:
    #     res = ""
    #     for pair in l:
    #         res += str(pair[0])+str(pair[1])

    #     return res
