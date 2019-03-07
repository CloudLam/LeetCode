/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (45.91%)
 * Total Accepted:    516.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 * 
 * Example:
 * 
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }lists
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    temp := result
    for l1 != nil || l2 != nil {
        if l1 == nil {
            temp.Next = l2
            l2 = nil
            continue
        }
        if l2 == nil {
            temp.Next = l1
            l1 = nil
            continue
        }
        if l1.Val > l2.Val {
            temp.Next = l2
            temp = temp.Next
            l2 = l2.Next
        } else {
            temp.Next = l1
            temp = temp.Next
            l1 = l1.Next
        }
    }
    return result.Next
}
