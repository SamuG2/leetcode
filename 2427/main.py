from math import floor

class Solution:
    def commonFactors(self, a: int, b: int) -> int:
        f=0
        for i in range(1, min(a,b)+1):
            if a%i==0==b%i:
                f+=1
        return f