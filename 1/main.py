class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        vistos = {}
        for i, v in enumerate(nums):
            if target-v in vistos:
                return [i, vistos[target-v]]
            else:
                vistos[v] = i
        print(vistos)
        return []



if __name__ == '__main__':
    print(Solution().twoSum([2,7,11,15], 9))