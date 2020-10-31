// Description: Friends Numbers in Golang.
//
// Author     : Adrian Statescu <http://thinkphp.ch>
//
// License    : MIT

package main

import "fmt"

const MAX = 3000

func isFriend(a, b int, v [ MAX ]int) (out bool) {

     if v[ a - 1 ] == b && v[ b - 1 ] == a {

     	out = true  

     } else { 

        out = false;
     }   
          
    return      
}

func main() {     

     var i, j, sd int

     var v[ MAX ]int

     for i = 1; i <= MAX; i++ {

     	 sd = 1

         for j = 2; j <= i / 2; j++ {

         	 if i % j == 0 {

         	 	sd = sd + j
         	 }
         }

         v[i - 1] = sd
     }

     for i = 1; i < MAX; i++ {

     	for j = i + 1; j <= MAX; j++ {
       
            if isFriend(i, j, v)  {

         	      fmt.Println(i, j, "Friends")         
            }  	      
        }    
     }
     
}
