class Solution:
    # def minKBitFlips(self, nums: list[int], k: int) -> int:
    #     if k == 1:
    #         return nums.count(0)
    #     nums = [n == 1 for n in nums]
    #     transformations = 0
    #     i = 0
    #     while True:
    #         found = False
    #         for t, v in enumerate(nums):
    #             if not v:
    #                 i = t
    #                 found = True
    #                 break
    #         if not found:
    #             return transformations
    #         if i+k > len(nums):
    #             return -1
    #         nums = list(map(lambda x: not x, nums[i:i+k]))+nums[i+k:]
    #         transformations += 1
    
    def minKBitFlips(self, nums: list[int], k: int) -> int:
        


if __name__ == "__main__":
    nums = [0,0,0,1,0,1,1,0]
    k = 3
    print(Solution().minKBitFlips(nums, k))
