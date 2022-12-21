
class Solution:
    def champagneTower(self, poured: int, query_row: int, query_glass: int) -> float:
        glasses = [[0 for _ in range(x)] for x in range(1, 101)]
        glasses[0][0] = poured

        for i in range(query_row):
            for j in range(len(glasses[i])):
                temp = (glasses[i][j] - 1) / 2.0
                if temp > 0:
                    glasses[i+1][j] += temp
                    glasses[i+1][j+1] += temp

        return glasses[query_row][query_glass] if glasses[query_row][query_glass] <= 1 else 1


if __name__ == "__main__":
    print(Solution().champagneTower(6, 3, 2))
