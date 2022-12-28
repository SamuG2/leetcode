import heapq


class Solution:
    def minStoneSum(self, piles: list[int], k: int) -> int:
        heap = [-n for n in piles]
        heapq.heapify(heap)
        for _ in range(k):
            heapq.heapreplace(heap, heap[0]//2)

        return -sum(heap)


if __name__ == "__main__":
    piles = [4, 3, 6, 7]
    k = 3
    Solution().minStoneSum(piles, k)
