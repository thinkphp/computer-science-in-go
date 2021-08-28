/**
 *
 * Author:
 * Adrian Statescu http://adrianstatescu.com
 *
 * Description:
 * Heap Sort Algorithm
 *
 * Time Complexity with Big O Notation:
 *
 *              Worst-case  : O(n log n)
 *              Best-case   : O(n log n)
 *              Average-case: O(n log n)
 *
 * License:
 *  MIT
 *
 */
package main

import "fmt"

var Heap [ 500500 ]int

var sizeH int

func up(child int) {

     var parent = child / 2

     for parent != 0 {

         if Heap[ parent ] > Heap[ child ] {

            aux := Heap[ parent ]

            Heap[ parent ] = Heap[ child ]

            Heap[ child ] = aux

            child = parent

            parent = child / 2

         } else {

           break
        }
     }

}

func down(parent int) {

    for 2 * parent <= sizeH {

        var child = 2 * parent

        if (2 * parent + 1) <= sizeH && Heap[ 2 * parent + 1 ] < Heap[ 2 * parent ] {

                child += 1
        }

        if Heap[ parent ] <= Heap[ child ] {

           break
        }

        aux := Heap[parent]

        Heap[parent] = Heap[child]

        Heap[child] = aux

        parent = child
    }

}

func removeHeap() int {

     var ret = Heap[ 1 ]

         Heap[ 1 ] = Heap[ sizeH ]

         sizeH -= 1

         down( 1 )

         return ret
}

func insertHeap(value int) {

    sizeH += 1

    Heap[ sizeH ] = value

    up( sizeH )
}

func main() {

     arr := []int{ -50, 9, 8, 7, 6, 5, -4, -3, 2, 1, -20, 10, 20 , 30 , 40, 50}

     n := len( arr )

     for i := 0 ; i < n; i++ {

       fmt.Printf("%d ", arr[i] )

     }

     fmt.Printf("\n")

     sizeH = 0

     for _, v := range arr {

        insertHeap( v )
     }

     for i := 1 ; i <= n; i++ {

         fmt.Printf("%d ", removeHeap())
     }
}
