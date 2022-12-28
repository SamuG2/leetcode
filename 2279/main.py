class Solution:
    def maximumBags(self, capacity: list[int], rocks: list[int], additionalRocks: int) -> int:
        #both = [capacity[i]-rocks[i] for i in range(len(capacity))]
        both = [cap-rock for cap, rock in zip(capacity, rocks)]
        both.sort()
        res = 0
        for n in both:
            if n == 0 or additionalRocks >= n:
                additionalRocks -= n
                res += 1
            else:
                break
        return res
        # both = list(zip(capacity, rocks))
        # both.sort(key= lambda x: x[0]-x[1])
        # res = 0
        # i=0
        # while additionalRocks>0:
        #     if both[i][0]-both[i][1] == 0:
        #         res+=1
        #     else:
        #         additionalRocks -= both[i][0]-both[i][1]
        #         if additionalRocks >=0:
        #             res+=1
        #             both[i][1] = both[i][0]


if __name__ == "__main__":
    capacity = [2, 3, 4, 5]
    rocks = [1, 2, 4, 4]
    additionalRocks = 2
    Solution().maximumBags(capacity, rocks, additionalRocks)
