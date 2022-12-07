class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        i = 0
        intervals.sort(key=lambda x: x[0])

        while i < len(intervals)-1:
            if intervals[i][1]>=intervals[i+1][0] and intervals[i][0]<=intervals[i+1][1]:
                intervals[i][0] = min(intervals[i][0], intervals[i+1][0])
                intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
                intervals.pop(i+1)
            else:
                i += 1
        return intervals
