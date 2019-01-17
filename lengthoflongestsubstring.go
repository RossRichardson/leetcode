/*
  lengthOfLongestSubstring
  ref: https://leetcode.com/problems/longest-substring-without-repeating-characters
  
  A simple implementation that beats 47% of submissions.
  
  Complexity: 
    - Time: Not yet analyzed
    - Space: Not yet analyzed
*/


import "fmt"

func main() {
    testLengthOfLongestSubstring("abcabcbb","abc")
    testLengthOfLongestSubstring("bbbbb","b")
    testLengthOfLongestSubstring("pwwkew","wke")
    testLengthOfLongestSubstring("bb","b")    
    testLengthOfLongestSubstring("aab","ab")    
    testLengthOfLongestSubstring("dvdf","vdf")
    testLengthOfLongestSubstring("fddvdf","vdf")    
    testLengthOfLongestSubstring("anviaj","nviaj")
    testLengthOfLongestSubstring("efgabcd","efgabcd")
    testLengthOfLongestSubstring("aabaab!bb","ab!")    
    testLengthOfLongestSubstring("eeydgwdykpv","gwdykpv")
}

func testLengthOfLongestSubstring(res string, exp string) {    
    var result = lengthOfLongestSubstring(res)
    fmt.Println("input:",res)
    fmt.Println("result: ", result, "pass:", result==len(exp))
    fmt.Println("")
}
    
func lengthOfLongestSubstring(s string) int {
        
  if len(s) == 1 {
        return 1
    }
    
    li := 0
    lj := 0
    x:=len(s)
    
    for i := 0; i <= x; i++ {
        
        if lj - li > x-1-i {            
            break;
        }
                
        for j := i+1; j < x; j++ {
            
            if j < lj {
                j=lj-1
                continue
            }
            
            l := strings.IndexByte(s[i:j], s[j])
            if l > -1 {
                // need to terminate early           

                if lj - li < j-i {
                    // store longest candidate
                    li = i
                    lj = j
                }
                
                if l >= i {
                    // start from the point we detected the dupe, if that point is after i
                    i = l
                }
                break
            }
            
            if j == x-1 {
                // at the end but don't need to terminate early
                if lj - li < x-i {
                    // store longest candidate
                    li = i
                    lj = x
                }
            }
            
        }
        
    }
    
    return lj-li
}
