class Solution:
    def rotate(self, matrix: list[list[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # m = [[0]*len(r) for r in matrix]
        # for i, row in enumerate(matrix):
        #     for j, v in enumerate(row):
        #         m[j][len(row)-i-1] = v

        # for i in range(len(matrix)):
        #     for j in range(len(matrix[i])):
        #         matrix[i][j] = m[i][j]
        for i in range(len(matrix)/2):
            matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
        
        for i in range(len(matrix)):
            for j in range(i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

        print(matrix)


if __name__ == '__main__':
    #matrix = [[5, 1, 9, 11], [2, 4, 8, 10], [13, 3, 6, 7], [15, 14, 12, 16]]
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

    Solution().rotate(matrix)
