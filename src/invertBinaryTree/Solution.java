package invertBinaryTree;

public class Solution {

	public TreeNode invertTree(TreeNode root) {
		if(root == null){
			return root;
		}else{
			root.left = invertTree(root.left);
			root.right = invertTree(root.right);
			TreeNode tn = root.left;
			root.left=root.right;
			root.right=tn;
			return root;
		}
    }
	
}
