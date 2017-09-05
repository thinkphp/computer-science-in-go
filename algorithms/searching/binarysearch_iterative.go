package main

import "fmt"

func BinarySearch(arr []int, search int) (middle int) {

     lo := 0

     hi := len(arr) - 1

     for lo <= hi {

         middle = (lo+hi)>>1
         
         if search < arr[middle] {

            hi = middle - 1 

         } else if search > arr[middle] {

            lo = middle + 1

         } else  { return middle }
     }

     return -1
}

func main() {
    
     searchField := []int{1,3,5,7,9,13,15,17}

     fmt.Println(searchField)

     for _, searchNumber := range searchField {

            fmt.Printf("Your number %d was found in position: %d\n", searchNumber, BinarySearch(searchField, searchNumber))
     }

     fmt.Printf("Your number %d was found in position: %d\n", 117, BinarySearch(searchField, 117))
}