package main

import ("fmt"
	    "math")

func generate_subsets(vec string, dim uint) {

     var i, j uint

     var total uint

     total = uint(math.Pow(2,float64(dim)) )

     for i = 1; i < total; i++ {

     	 for j = 0; j < dim; j++ {
     	 	 
     	 	if i & (1<<j) > 0 {

     	 	 	fmt.Printf("%c ", vec[j])
     	 	 } 
     	 }

         fmt.Println("")     	 
     }
}

func main() {


	 var myStr string

	 myStr = "abcd"

     var dim uint

	 dim = uint(len(myStr))

	 generate_subsets(myStr, dim)
}
