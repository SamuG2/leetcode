class Solution:
    def largestDivisibleSubset(self, nums: list[int]) -> list[int]:
        if len(nums) == 1:
            return nums
        nums.sort()
        one = False
        if nums[0] == 1:
            one = True
            nums.pop(0)
        max_subset = []
        checked = {}
        for i in range(len(nums)):
            if checked.get(i, True):
                candidate = [1]
                for t in range(i, len(nums)):
                    if nums[t] % candidate[-1] == 0:
                        candidate.append(nums[t])
                        checked[t] = False
                if len(candidate) > len(max_subset):
                    max_subset = candidate
        return max_subset if one else max_subset[1:]


if __name__ == "__main__":
    nums = [5, 9, 18, 54, 108, 540, 90, 180, 360, 720]

    Solution().largestDivisibleSubset(nums)
