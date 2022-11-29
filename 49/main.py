from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        dic = defaultdict(list)
        for word in strs:
            dic[tuple(sorted(word))].append(word)
        return list(dic.values())