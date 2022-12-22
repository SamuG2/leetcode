class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        mapa = {}
        for i, n in enumerate(nums):
            if n in mapa:
                return [mapa[n], i]
            mapa[target-n] = i
        return []



if __name__ == '__main__':
    print(Solution().twoSum([2,7,11,15], 9))