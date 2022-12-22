import re
class Solution:
    def winnerOfGame(self, colors: str) -> bool:
        return re.findall('(?=AAA)', colors) > re.findall('(?=BBB)', colors) 
        # con_A, con_B = 0, 0
        # for i in range(1, len(colors)-1):
        #     if colors[i] == colors[i-1] == colors[i+1] == 'A':
        #         con_A += 1
        #     elif colors[i] == colors[i-1] == colors[i+1] == 'B':
        #         con_B += 1
        # return con_A > con_B
