class Solution:
    def maxNumber(self, nums1: list[int], nums2: list[int], k: int) -> list[int]:
        res = [0]*k
        for i in range(k):
            largest1, idx1 = self.findLargest(nums1, len(nums1)-(k-len(nums2)-1))
            largest2, idx2 = self.findLargest(nums2, len(nums2)-(k-len(nums1)-1))
            if largest1 > largest2:
                res[i] = largest1
                nums1 = nums1[idx1+1:]
            elif largest1 < largest2:
                res[i] = largest2
                nums2 = nums2[idx2+1:]
            else:
                if self.compare(nums1[:idx1], nums2[:idx2]):
                    res[i] = largest2
                    nums2 = nums2[idx2+1:]
                    
                else:
                    res[i] = largest1
                    nums1 = nums1[idx1+1:]

            k -= 1

        return res

    def compare(self, nums1, nums2):
        if not nums1 and not nums2:
            return True
        largest1, idx1 = self.findLargest(nums1, len(nums1))
        largest2, idx2 = self.findLargest(nums2, len(nums2))

        if largest1 > largest2:
            return True
        elif largest2 > largest1:
            return False
        else:
            return self.compare(nums1[:idx1], nums2[:idx2])

    def findLargest(self, nums: list, max_len: int):
        largest = 0
        idx = 0
        for i in range(min(max_len, len(nums))):
            if nums[i] > largest:
                largest = nums[i]
                idx = i
                if largest == 9:
                    break
        return largest, idx


if __name__ == "__main__":
    nums1 = [6,7]
    nums2 = [6,0,4]
    k = 5

    Solution().maxNumber(nums1, nums2, k)
