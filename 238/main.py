class Solution:
    def productExceptSelf(self, nums: list[int]) -> list[int]:
        left, right = [1]*len(nums), [1]*len(nums)
        for i in range(1, len(nums)):
            left[i] = left[i-1]*nums[i-1]
            right[len(nums)-i-1] = right[len(nums)-i]*nums[len(nums)-i]
        for i in range(0, len(nums)):
            nums[i] = left[i]*right[i]

        return nums
