package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		return (float64(nums1[(len(nums1)-1)/2]) + float64(nums1[((len(nums1)-1)+1)/2])) / 2
	}
	return float64(nums1[(len(nums1)-1)/2])

}

func main() {

}
