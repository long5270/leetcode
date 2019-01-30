/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var nodeList [] *ListNode
    current := head
    for {
        if current == nil{
            break
        }
        nodeList = append(nodeList, current)
        current = current.Next
    }
    length := len(nodeList)
    if n == length{
        if length > 1{
            return nodeList[1]
        } else {
            return nil
        }
    }
    if n == 1{
        nodeList[length - n - 1].Next = nil
    } else{
        nodeList[length - n - 1].Next = nodeList[length - n + 1] 
    }
    return head
}
