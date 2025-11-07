package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)               // initializing map to store element and its index
    for i := 0; i < len(nums); i++{      
        complement := target - nums[i]   // difference between target and current element
        if j, ok := m[complement]; ok{   // check if the map contains already complement value
            return []int{i, j}           // return index of current element and complements's index from the map 
        }
        m[nums[i]] = i                   // keep storing elements with indices

    }
    return []int{}
}

func main(){
    nums := []int{2, 7, 11, 15}
    target := 9

    fmt.Println(twoSum(nums, target))
}
