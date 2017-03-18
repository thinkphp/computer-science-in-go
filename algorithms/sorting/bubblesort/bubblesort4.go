/**
 *
 * Author     : Adrian Statescu mergesortv@gmail.com http://adrianstatescu.com
 *
 * Description: BubbleSort is a simple sorting algorithm that repeatedly steps through the elements
 * in an array to be sorted, compares each pair of adjiacent elements and swap thems 
 * if they are in the wrong order.  The pass through the list is repeated until no more swaps is needed.
 *
 * Complexity : Worst-case O(n^2)
 *              Best-case O(n)  
 *              Average-case O(n^2) 
 *
 *
 * MIT License 
 *
 */

package main

import ("fmt"
        "math/rand")

func bubblesort(arr []int) {

     var i, j int

     for i = 0; i < len(arr) - 1; i++ {

         for j = 0; j < len(arr) - i - 1 ; j++ {

             if arr[ j ] > arr[ j + 1 ] {

                arr[ j ], arr[ j + 1 ] = arr[ j + 1 ], arr[ j ]
             }
         }
     }   
}

func main() {

     size := 10

     var arr = make([]int, size)

     index := 0

     for index < size {

         arr[ index ] = rand.Intn(size)

         index++ 
     } 

     bubblesort( arr )

     for _, v := range arr {

         fmt.Printf("%d ", v)
     }   
}