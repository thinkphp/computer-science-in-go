package main

import ("fmt"
        "io"
        "os"
)

func outputFile(arr []int, n int) {

    var i int 

    file, err := os.Create("algsort2.out") // creating...

    if err != nil {

        fmt.Printf("error creating file: %v", err)

        return
    }

    defer file.Close()

    for i = 0; i < n; i++ { // loop through array

        _, err = fmt.Fprintf(file, "%d ", arr[i])

        if err != nil {

            fmt.Printf("error writing integers: %v", err)

        }
    }
}

func bubblesort(arr []int, n int) {

     var finished bool
     var swapped bool
     var i int 
     var size int 

     finished = false

     size = n 

     for !finished {
    
         swapped = false

         for i = 0; i < size - 1; i++ {

             if arr[i] > arr[i+1] {

                arr[i], arr[i+1] = arr[i+1], arr[i] 

                swapped = true 
             }
         }        
         
         if swapped {

            size--

         } else { 

                finished = true 
         }
     }  
}

func main() {

     file, err := os.Open("algsort.in")

     if err!= nil {

        fmt.Println(err)

        os.Exit(1) 
     }

     var n int
     var item int
     var arr []int

     fmt.Fscanf(file, "%d\n", &n) //give a pattern to scan

     for {
         _, err := fmt.Fscanf(file, "%d", &item) //give a pattern to scan

         if err != nil {

            if err == io.EOF {
                   break //stop reading the file
            }
            fmt.Println(err)
            os.Exit(1)
         }

         arr = append(arr, item)
     }

     fmt.Print("Input Array -> ")
     fmt.Println(arr)

     bubblesort(arr, n)

     outputFile(arr, n)
     fmt.Print("Output Array -> ")
     fmt.Println(arr) 
}
       