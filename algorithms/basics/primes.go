//
//  Description: Display n Primes with n scanf from std
//
//  Author     : Adrian Statescu <http://t...content-available-to-author-only...p.ch>        
//
//  License    : MIT
// 
 
package main
 
import "fmt"
 
func isPrime(n int)(bool) {
 
	 if n <= 1 {
 
	 	return false
	 }
 
     if n == 2 || n == 3 {
 
     	return true
     }  
 
     prime := true    
 
     for k := 2 ;k * k <= n && prime; {         
 
         prime = n % k != 0
 
         k++
     }
 
     return prime
}
 
func main() {
 
     var n int    
 
     fmt.Scan(&n)
 
     var i int
 
     counter := 0
 
     i = 2
 
     for {
 
         if isPrime( i ) {
 
         	fmt.Printf("%d ", i)
 
         	counter++
         }
 
     	 if counter == n {
 
     	 	break
 
     	 } else {
 
     	 	i++
     	 }
     }
}
