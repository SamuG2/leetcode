class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right



class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        nodes = self.exploreNodes(root, k)
        return nodes[k-1]
    
    def exploreNodes(self, root: TreeNode, max_depth: int)->list[TreeNode]:
        if not root:
            return []
        res = []
        res+=self.exploreNodes(root.left, max_depth)
        if len(res)>=max_depth:
            return res
        res.append(root.val)
        if len(res)>=max_depth:
            return res
        res+=self.exploreNodes(root.right, max_depth)
        if len(res)>=max_depth:
            return res
        return res
        
        
    


if __name__ == "__main__":
    root = TreeNode(val=5, left=TreeNode(val=3, left=TreeNode(
        val=2, left=TreeNode(val=1)), right=TreeNode(val=4)), right=TreeNode(val=6))
    k = 3
    Solution().kthSmallest()
