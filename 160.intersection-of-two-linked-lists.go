/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pA, pB := headA, headB
    redirectA, redirectB := false, false
    for pA != nil && pB != nil {
        if pA == pB {
            return pA
        }
        pA, pB = pA.Next, pB.Next
        if pA == nil && !redirectA {
            pA = headB
            redirectA = true
        }
        if pB == nil && !redirectB {
            pB = headA
            redirectB = true
        }
    }
    return nil
}
