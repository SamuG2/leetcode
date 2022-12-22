class Solution:
    def sortPeople(self, names: list[str], heights: list[int]) -> list[str]:
        pairs = list(zip(names, heights))
        pairs.sort(key=lambda p: p[1], reverse=True)
        return [p[0] for p in pairs]
