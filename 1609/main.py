class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isEvenOddTree(self, root: TreeNode) -> bool:
        return self.bfs([root], True)

    def bfs(self, nodes: list[TreeNode], even: bool) -> bool:
        if not nodes:
            return True
        next_level = []
        for i, node in enumerate(nodes):
            if even:
                if node.val % 2 == 0 or (i != 0 and node.val <= nodes[i-1].val):
                    return False
            else:
                if node.val % 2 != 0 or (i != 0 and node.val >= nodes[i-1].val):
                    return False
            if node.left:
                next_level.append(node.left)
            if node.right:
                next_level.append(node.right)
        return self.bfs(next_level, not even)


if __name__ == "__main__":
    root = TreeNode(val=1, left=TreeNode(val=10, left=TreeNode(val=3, left=TreeNode(
        val=12), right=TreeNode(val=8))), right=TreeNode(val=4, left=TreeNode(val=7, left=TreeNode(val=6)), right=(TreeNode(val=9, right=TreeNode(val=2)))))

    Solution().isEvenOddTree(root)
