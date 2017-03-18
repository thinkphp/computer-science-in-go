/**
 *
 * Author     : Adrian Statescu mergesortv@gmail.com http://adrianstatescu.com
 *
 * Description: BubbleSort is a simple sorting algorithm that repeatedly steps through the elements
 * in an array to be sorted, compares each pair of adjiacent elements and swap thems 
 * if they are in the wrong order. 
 * The pass through the list is repeated until no more changes swaps are needed.
 * 
 * Complexity : Worst-case O(n^2)
 *              Best-case O(n)  
 *              Average-case O(n^2) 
 *
 * MIT License 
 *
 */
package main

import "fmt"

func bubblesort(arr []int) {

     done := true
 
     for done {

         changes := 0
         
         for i := 0; i < len(arr) - 1; i++ {

             if arr[ i ] > arr[ i + 1 ] {

                arr[ i ], arr[ i + 1 ] = arr[ i + 1 ], arr[ i ]

                changes++ 
             }   
         } 

         if changes == 0 {

            done = false
         }
     }
}

func main() {

     arr := []int{9,8,7,6,5,4,3,2,1}

     bubblesort( arr )

     for _, v := range arr {

         fmt.Printf("%d ", v)
     }   
}