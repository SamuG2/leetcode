class Solution:
    def firstMissingPositive(self, nums: list[int]) -> int:
        nums_dic = {i: False for i in range(1, len(nums)+1)}

        for n in nums:
            if n > 0 and n < len(nums)+1:
                nums_dic[n] = True
        for i in range(1, len(nums)+1):
            if not nums_dic[i]:
                return i
        return len(nums)+1


if __name__ == '__main__':
    print(Solution().firstMissingPositive([3, 4, -1, 1]))
