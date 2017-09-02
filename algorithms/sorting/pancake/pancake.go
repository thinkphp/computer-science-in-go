package main

import "fmt"

func findMax(vec []int, index int) int {

     var iMax int
     var i int
 
     iMax = 0

     for i = 1; i < index; i+=1 {

         if vec[i] > vec[iMax] {

            iMax = i
         }  
     }
  
     return iMax
}

func _flip(vec []int, index int) {

     var start int

     start = 0

     for start < index {

         vec[start],vec[index] = vec[index],vec[start]

         start+=1

         index-=1
     }
}

func pancake(vec []int) {

     var curr int

     var iMax int

     size := len(vec)

     for curr = size; curr > 1; curr-=1 {

         iMax = findMax(vec, curr)

         _flip(vec, iMax)
         _flip(vec, curr-1)
     }
}

func main() {

     arr := []int{9,8,7,6,5,4,3,2,1,0}

     fmt.Println(arr)

     pancake(arr)

     fmt.Println(arr)
}