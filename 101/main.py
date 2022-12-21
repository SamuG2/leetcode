class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSymmetric(self, root):
        if root is None:
            return True
        else:
            return self.isMirror(root.left, root.right)

    def isMirror(self, left, right):
        if left is None and right is None:
            return True
        if left is None or right is None:
            return False

        if left.val == right.val:
            outPair = self.isMirror(left.left, right.right)
            inPiar = self.isMirror(left.right, right.left)
            return outPair and inPiar
        else:
            return False


if __name__ == "__main__":
    root = TreeNode(val=1, left=TreeNode(val=2, left=TreeNode(val=3), right=TreeNode(
        val=4)), right=TreeNode(val=2, left=TreeNode(val=4), right=TreeNode(val=3)))
    Solution().isSymmetric(root)
