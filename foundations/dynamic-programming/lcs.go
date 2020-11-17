package main

import "fmt"

func max(a,b int) (int) {

     if a > b {

     	return a

     } else { 

     	return b
     }	
}

func createSubseq(lcs [][]int, x []int, y []int, i int, j int) {
 
     if i == 0 || j == 0 {
     	return
     }

     if x[i] == y[j] {

     	createSubseq(lcs, x, y, i-1, j-1)

     	fmt.Printf("%d ", x[i])

     } else {

        if(lcs[i][j-1] < lcs[i-1][j]) {

        	createSubseq(lcs, x, y, i - 1, j)

        } else {

        	createSubseq(lcs, x, y, i, j - 1)

        }
     }
}

func main() {
 
     x := []int{0,1,7,3,5,8}

     y := []int{0,7,5,8,1}

     var n, m int;

     n = len(x)

     m = len(y)
     
     lcs := make([][]int, n)

         for i:= range lcs {

         	  lcs[i] = make([]int, m)
         }
     
     for i := 1; i <= n - 1; i++ {

     	  for j := 1; j <= m - 1; j++ {

             if x[i] == y[j]  {

                lcs[ i ][ j ] = 1 + lcs[i - 1][j - 1]

             } else {

             	lcs[ i ][ j ] = max(lcs[i][j - 1], lcs[ i -1 ][ j ])

             }
     	 	 
     	 } 	
     	 
     }

     fmt.Printf("%d\n", lcs[n-1][m-1]) 

     createSubseq(lcs, x, y, n - 1, m - 1) 

     i := n - 1
     j := m - 1

     var ans []int

     for i != 0 && j != 0 {

         if x[i] == y[j] {

            ans = append(ans, x[i]) //push

            i = i - 1
            j = j - 1    

         }  else {

         	if lcs[i][j-1] < lcs[i-1][j] {

         		i = i - 1; 

         	} else {

         		j = j - 1;
         	}

         }
     }

     fmt.Println("")

     for len(ans) > 0 {

     	 n := len(ans) - 1

     	 fmt.Printf("%d ", ans[n]) //top element

     	 ans = ans[:n] //pop
     }

     fmt.Println("")
}
