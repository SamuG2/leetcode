from itertools import combinations


class Solution:
    def maxScoreWords(self, words: list[str], letters: list[str], score: list[int]) -> int:
        chars = "abcdefghijklmnopqrstuvwxyz"
        score = {chars[i]: score[i] for i in range(26)}
        letters_map = {c: letters.count(c) for c in set(letters)}
        max_value = 0
        for i in range(1, len(words)+1):
            combs = combinations(words, i)
            for comb in combs:
                new_word = "".join(comb)
                if len(new_word) <= len(letters):
                    value = 0
                    valid = True
                    for c in set(new_word):
                        num = new_word.count(c)
                        if num > letters_map.get(c, 0):
                            valid = False
                            break
                        value += num*score[c]
                    if valid:
                        max_value = max(value, max_value)

        return max_value


if __name__ == "__main__":
    words = ["dog", "cat", "dad", "good"]
    letters = ["a", "a", "c", "d", "d", "d", "g", "o", "o"]
    score = [1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0,
             0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    Solution().maxScoreWords(words, letters, score)
