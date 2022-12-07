class Solution:
    def numDecodings(self, s: str) -> int:
        prev_c = s[0]
        if prev_c == '0':
            return 0
        count1 = 1
        count2 = 0
        for cur_c in s[1:]:
            if prev_c == '1' and cur_c != '0':
                count1, count2 = count1+count2, count1
            elif prev_c == '2' and cur_c in ('1', '2', '3', '4', '5', '6'):
                count1, count2 = count1+count2, count1
            elif cur_c == '0' and prev_c in ('2', '1'):
                count1, count2 = 0, count1
            elif cur_c == '0' and prev_c not in ('2', '1'):
                return 0
            else:
                count1, count2 = count1+count2, 0

            prev_c = cur_c
        return count1+count2


if __name__ == '__main__':
    s = "230"
    print(Solution().numDecodings(s))
