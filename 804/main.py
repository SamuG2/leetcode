class Solution:
    def uniqueMorseRepresentations(self, words: list[str]) -> int:
        char_to_morse = {'a': '.-', 'b': '-...', 'c': '-.-.', 'd': '-..', 'e': '.', 'f': '..-.', 'g': '--.', 'h': '....', 'i': '..', 'j': '.---', 'k': '-.-', 'l': '.-..','m': '--', 'n': '-.', 'o': '---', 'p': '.--.', 'q': '--.-', 'r': '.-.', 's': '...', 't': '-', 'u': '..-', 'v': '...-', 'w': '.--', 'x': '-..-', 'y': '-.--', 'z': '--..'}
        diffs = set(["".join([char_to_morse[c] for c in word]) for word in words])
        return len(diffs)


if __name__ == "__main__":
    words = ["gin", "zen", "gig", "msg"]
    Solution().uniqueMorseRepresentations(words)
