class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        if target < nums[0]:
            return 0
        elif target > nums[-1]:
            return len(nums)
        
        
        