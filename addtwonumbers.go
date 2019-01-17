/*
  addTwoNumbers
  ref: https://leetcode.com/problems/add-two-numbers
  
  A tuned implementation that beats 97% of submissions.
  It probably includes some unnesessary optimizations as well that
  just make it longer to read.
  
  Complexity: 
    - Time: Not yet analyzed
    - Space: Not yet analyzed
    
  Definition for singly-linked list:

  type ListNode struct {
      Val int
      Next *ListNode
  }
*/


import "fmt"

func main() {
    /*l1 := 
    l2 := 
    result := addTwoNumbers(l1,l2)*/
    fmt.Printf("not yet implemented")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    n1 := l1
    n2 := l2 
    
    carry := 0
    
    // add the first values
    val := n1.Val + n2.Val

    // carry logic
    if val > 9 {
        carry = 1
        val = val - 10
    }

    // create node
    a := ListNode {
        Val: val,
        Next: nil,
    }
    
    var answer *ListNode = &a
    var current *ListNode = &a
    
    if n1 != nil {
        if n1.Next == nil {
             n1 = nil
        }else{
            n1 = n1.Next
        }
    }

    if n2 != nil {
        if n2.Next == nil {
            n2 = nil
        }else{
            n2 = n2.Next
        }
    }

    if n1 == nil && n2 == nil {   

        if carry == 1 {
            // create last carry node
            a := ListNode {
                Val: carry,
                Next: nil,
            }

            current.Next = &a
        }
        
        return answer
    }
    
    
    for {
        
        val = carry
        
        if n1 != nil {           
            val = val + n1.Val
        }
        
        if n2 != nil {
            val = val + n2.Val
        }
        
        /// carry logic
        if val > 9 {
            carry = 1
            val = val - 10
        }else{             
            carry = 0
        }
        
        /// create node
        a := ListNode {
            Val: val,
            Next: nil,
        }
        
        current.Next = &a
        current = &a
        
        if n1 != nil {
            if n1.Next == nil {
                n1 = nil
            }else{
                n1 = n1.Next
            }
        }
        
        if n2 != nil {
            if n2.Next == nil {
                n2 = nil
            }else{
                n2 = n2.Next
            }
        }
        
        if n1 == nil && n2 == nil {   

            if carry == 1 {
                // create last carry node
                a := ListNode {
                    Val: carry,
                    Next: nil,
                }

                current.Next = &a
            }
            
            
            return answer
        }
    }
    
    return answer
}
