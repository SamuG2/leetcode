from collections import defaultdict, Counter


class Solution:
    def exist(self, board: list[list[str]], word: str) -> bool:
        board_dic = defaultdict(int)
        for r in board:
            for c in r:
                board_dic[c] += 1
        word_dic = Counter(word)
        for c in word_dic:
            if c not in board_dic or board_dic[c] < word_dic[c]:
                return False
            
        for i in range(len(board)):
            for j in range(len(board[i])):
                r = self.checkWord(board, word, i, j, [])
                if r:
                    return r
        return False

    def checkWord(self, board: list[list[str]], word: str, i: int, j: int, used: list[list[int]]) -> bool:

        if not word:
            return True
        if i < 0 or j < 0 or i >= len(board) or j >= len(board[i]) or board[i][j] != word[0] or [i, j] in used:
            return False

        for x, y in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            r = self.checkWord(board, word[1:], i+x, j+y, used+[[i, j]])
            if r:
                return r

        return False


if __name__ == '__main__':
    board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    word = "SEE"
    print(Solution().exist(board, word))
