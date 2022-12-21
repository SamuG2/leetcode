
class Solution:
    def minCost(self, nums: list[int], cost: list[int]) -> int:
        arr = sorted(zip(nums, cost))
        total, cnt = sum(cost), 0
        for num, c in arr:
            cnt += c
            if cnt > total // 2:
                target = num
                break
        return sum(c * abs(num - target) for num, c in arr)
    
if __name__ == "__main__":
    nums = [1,3,5,2]
    cost = [2,3,1,14]
    Solution().minCost(nums, cost)