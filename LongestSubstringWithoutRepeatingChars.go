package main

import (
    "fmt"
    "strings"
)

// func lengthOfLongestSubstring(s string) int {
//     var maxLength int
//     var currLength int = 1
//     i := 0
//     j := i+1
//     for j < len(s){
//         if strings.Contains(s[i:j], string(s[j])){ // check if current i:j substring contains j char 
//             i++                                    // if contains then increment i 
//             currLength = 1                         // and refresh the currLength value
//         }else{
//             currLength++                           // if not contains, then current substring has no duplicates, increment currLength
//         }
//         maxLength = max(maxLength, currLength)     // update here max length so far
//         j++                                        // increment j to the next char
//     }
//     return maxLength                            

// }

// best solution, put here as an example
func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    maxLen := 0
    start := 0

    for i, ch := range s {
        if prevIndex, found := charIndex[ch]; found && prevIndex >= start {
            start = prevIndex + 1
        }
        charIndex[ch] = i
        maxLen = max(maxLen, i-start+1)
    }

    return maxLen
}


func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func main(){
    s := "abcabcbb"
    fmt.Println(lengthOfLongestSubstring(s))
}

/*

a b c a b c b b

abc
bca
cab

3 length of longest substring





*/
