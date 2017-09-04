package main 

import "fmt"

func BinarySearch(arr []int, search int) (result int, searchCount int){

     middle := len(arr) / 2

     switch {

         case len(arr) == 0: result = -1//not found

         case arr[middle] > search:

              result, searchCount = BinarySearch(arr[:middle], search)               

         case arr[middle] < search:

              result, searchCount = BinarySearch(arr[middle+1:], search)
              result = middle + 1

         default: result = middle  
     }

     searchCount++
 return
}

func main() {

     searchField := []int{2, 5, 7, 9, 11, 13, 15, 17, 23}

     for _, searchNumber := range searchField {

         fmt.Printf("Searching for %d in list %v\n", searchNumber, searchField)

         result, searchCount := BinarySearch(searchField, searchNumber) 

         fmt.Printf("Your number was found in position %d after %d steps.\n\n", result, searchCount)
     }
}