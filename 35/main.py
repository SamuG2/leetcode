class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        if target < nums[0]:
            return 0
        elif target > nums[-1]:
            return len(nums)

        left = 0
        right = len(nums)
        idx = 0
        while left < right:
            idx = left + (right-left)//2
            if nums[idx] > target:
                right = idx
            elif nums[idx] < target:
                left = idx + 1
            else:
                return idx
        return left


if __name__ == "__main__":
    pass
