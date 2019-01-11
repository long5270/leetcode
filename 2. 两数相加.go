/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    list1 := ListNode2Number(l1)
    list2 := ListNode2Number(l2)
    minLen, minList := len(list1), list1
    maxLen, maxList := len(list2), list2
    if maxLen < minLen{
        minList, maxList = list2, list1
        minLen, maxLen = maxLen, minLen
    }
    var before int
    var afterNode *ListNode
    startNode := ListNode{}
    afterNode = &startNode
    l3 := &startNode
    for i,v := range maxList{
        tmp := afterNode 
        var sum int
        if i < minLen{
            sum = v.Val + minList[i].Val + before
            
        } else{
            sum = v.Val + before
        }
        if sum >= 10{
                tmp.Val = sum % 10
                before = 1
        }else{
            tmp.Val = sum
            before = 0  
        }
        if i + 1 < maxLen{
            afterNode = &ListNode{}
            tmp.Next = afterNode
        }

    }
    if before > 0{
        afterNode.Next = &ListNode{Val:before}
    }
    return l3
}

func ListNode2Number(l *ListNode) []*ListNode{
    var lastNode *ListNode 
    var list [] *ListNode
    lastNode = l
    for lastNode != nil {
        list = append(list, lastNode)
        lastNode = lastNode.Next
    }
    return list
}
