package main

import "fmt"

func _reverse(v []int, lo, hi int) {

	  
      if lo < hi {

	  	  v[lo], v[hi] = v[hi], v[lo]

          _reverse(v, lo + 1, hi - 1)	  	    
	  }
}

func main() {

	 arr :=[]int{1,2,3,4,5,6,7,8,9}

	 var n int

	 n = len(arr)

	 for i:=0; i < n; i++ {

	 	 fmt.Printf("%d ", arr[i])
	 }

	 fmt.Printf("\n")

     _reverse(arr, 0, n - 1)

     for i:=0; i < n; i++ {

	 	 fmt.Printf("%d ", arr[i])
	 }

	 fmt.Printf("\n")     
}