class Solution:

    def pList(self, p: str) -> list:
        plist = []
        for i in range(len(p)):
            if p[i] != '*':
                if plist:
                    if len(plist[-1]) > 1 and plist[-1][0] != p[i]:
                        plist.append(p[i])
                    else:
                        plist.append(p[i])
                else:
                    plist.append(p[i])

            else:
                plist[-1] += '*'
        return plist

    def checkCharacter(self, s: str, p: str) -> int:
        if s[0] == p or p == '.':
            return 0
        elif len(p) == 2:
            if s[0] == p[0]:
                return 1
            if p[0] == '.':
                return 2
            else:
                return 3
        else:
            return -1

    def isMatch(self, s: str, p: str) -> bool:
        if s == p:
            return True
        p = self.pList(p)
        try:
            while len(s) != 0:

                r = self.checkCharacter(s, p[0])
                if r == 0:  # caracter simple se corresponde o . sin *
                    s = s[1:]
                    p = p[1:]
                elif r == 1:  # el caracter conincide con el primero de x*
                    c = s[0]
                    p = p[1:]
                    i = len(s)
                    while i >= 0:
                        if s == "" or s[0] == c:
                            if self.isMatch(s, p):
                                return True
                            s = s[1:]
                            i -= 1
                        elif s[0]!= c:
                            break
                        else:
                            return False

                elif r == 2:  # el primer caracter es ., entra todo. Hay que comprobar el final. Si el final coincide con algúna posibilidad entonces hay match
                    p = p[1:]
                    i = len(s)
                    while i >= 0:
                        if self.isMatch(s, p):
                            return True
                        s = s[1:]
                        i -= 1
                    return False

                elif r == 3:  # el caracter no coincide con el primero de x*
                    p = p[1:]
                elif r == -1:
                    return False
            while len(p) > 0:
                if '*' in p[0]:
                    p = p[1:]
                else:

                    return False
            return True
        except Exception:
            return False


if __name__ == "__main__":
    print(Solution().isMatch("aab", "c*a*b"))
