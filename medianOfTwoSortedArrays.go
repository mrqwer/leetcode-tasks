package main

import "fmt"


// just used the the law of medians
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    low1 := 0
    high1 := len(nums1) - 1
    var med1 float64 = (float64(nums1[low1]) + float64(nums1[high1])) / 2

    low2 := 0
    high2 := len(nums2) - 1
    var med2 float64 =  (float64(nums2[low2]) + float64(nums2[high2])) / 2
    return (med1 + med2) / 2
}

func main(){
    nums1 := []int{1, 3}
    nums2 := []int{2}

    fmt.Println(findMedianSortedArrays(nums1, nums2))
}

/*
[1,3]
[2]

2.0

[1,2]
[3,4]


[1 2 3 4 5 6]

[1 2 3]
[4 5 6]

2 5


*/
