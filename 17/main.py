class Solution:
    def letterCombinations(self, digits: str) -> list[str]:
        if not digits:
            return []

        num_map = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g', 'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z']
        }

        res = ['']
        for dig in digits:
            temp_res = [seq+c for c in num_map[dig] for seq in res]
            res = temp_res

        return res


if __name__ == "__main__":
    print(Solution().letterCombinations("23"))
