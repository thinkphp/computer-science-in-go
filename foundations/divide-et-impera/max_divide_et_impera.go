package main

import "fmt"

func getMax(vec []int, lo,hi int) (int) {

     if lo == hi {

     	return vec[ lo ]

     } else {

       var a, b, mid int

       mid = (lo + hi)>>1

       a = getMax(vec, lo, mid)

       b = getMax(vec, mid + 1, hi)

       if a > b {

       	 return a

       } else {

       	 return b
       }

     }     
}

func main() {

     vec := []int{16,25,34,43,52,16,711,81,91}

     var max int

     max = getMax(vec, 0, len(vec) - 1)

     fmt.Println(vec)

     fmt.Printf("Max = %d ", max)

}
