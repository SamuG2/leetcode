class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        if(not head):
            return head
        odd_cur = head
        even_cur = head.next
        evenHead = even_cur
        while(even_cur and odd_cur and even_cur.next and odd_cur.next):
            odd_cur.next = even_cur.next
            odd_cur = odd_cur.next
            even_cur.next = odd_cur.next
            even_cur = even_cur.next
        odd_cur.next = evenHead
        return head


def createNodes(values: list)-> ListNode:
    head = ListNode()
    cur = head
    for i in range(len(values)):
        cur.val = values[i]
        if i < len(values) -1:
            cur.next = ListNode()
            cur = cur.next
    return head

if __name__ == "__main__":
    head = createNodes([1,2,3,4,5])
    
    res = Solution().oddEvenList(head)
    
    while res:
        print(res.val)
        res = res.next