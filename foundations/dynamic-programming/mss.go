package main

import "fmt"

func max( a, b int ) ( int ) {

     if a > b {

     	return a

     } else {

     	return b
     }
}

func mss(v []int, n int) (maxsum, start, end int){

     var possibleStart int

     possibleStart = 0

     maxsum = v[ 0 ]

     start = 0

     end = -1

     currentSum := 0


     for i := 0; i < n; i++ {

     	 if currentSum < 0 {

     	 	possibleStart = i
     	 }

     	 currentSum = max( currentSum + v[ i ], v[ i ])

     	 if currentSum > maxsum {

     	 	maxsum = currentSum

     	 	start = possibleStart

     	 	end = i
     	 }
     }

     start++

     end++
 
 return
}

func main() {

     v := []int{2, -3, 5, 7, -1, 4}

     n := len(v)

     fmt.Println( mss(v, n) )    
}  
