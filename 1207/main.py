from collections import defaultdict, List

class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        dic = defaultdict(int)
        for v in arr:
            dic[v]+=1
        ocurrences=[]
        for v in dic.values():
            if v in ocurrences:
                return False
            ocurrences.append(v)
        return True