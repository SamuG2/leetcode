class Solution:
    def canJump(self, nums: list[int]) -> bool:
        return True if len(nums) < 2 else self.checkJump(nums, len(nums)-1, [])[0]

    def checkJump(self, nums: list[int], idx: int, checked: list[int]) -> list[bool, list]:
        if idx == 0:
            return True, checked

        for i in checked:
            if i == idx:
                return False, checked
        for i in range(idx-1, -1, -1):
            if i+nums[i] >= idx:
                res = self.checkJump(nums, i, checked)
                if res:
                    return True, checked
        checked.append(idx)
        return False, checked


nums = [1, 2]
Solution().canJump(nums)
