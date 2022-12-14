
from math import floor


class Solution:
    def getAverages(self, nums: list[int], k: int) -> list[int]:
        averages = [-1]*len(nums)
        suma = sum(nums[0:k*2+1])
        nums.append(0)
        n = k*2+1
        for i in range(0, len(nums)-k*2-1):
            averages[k+i] = floor(suma/n)
            suma = suma - nums[i] + nums[n+1]
        return averages


if __name__ == "__main__":
    nums = [7, 4, 3, 9, 1, 8, 5, 2, 6]
    k = 3
    print(Solution().getAverages(nums, k))
