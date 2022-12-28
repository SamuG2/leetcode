class Solution:
    def replaceWords(self, dictionary: list[str], sentence: str) -> str:
        sentence = sentence.split(" ")
        roots = {r: True for r in dictionary}
        res = []
        for word in sentence:
            for i in range(1, len(word)+1):
                if roots.get(word[:i], False):
                    word = word[:i]
                    break
            res.append(word)
        return " ".join(res)
                