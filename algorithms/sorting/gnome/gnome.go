package main

import ("fmt")

func gnome(arr []int) {

     n := len(arr)

     var pos int

     pos = 0

     for pos < n {

         if pos == 0 || arr[pos] >= arr[pos-1] {

            pos += 1

         } else {

            arr[pos], arr[pos-1] = arr[pos-1], arr[pos]

            pos -=1 
         }
         
     }  
}

func main() {

     arr := []int{9,8,7,6,5,4,3,2,1,0}

     fmt.Println(arr)

     gnome(arr)

     fmt.Println(arr)
}