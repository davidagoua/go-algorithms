package main

/*
 * Counting sort - https://en.wikipedia.org/wiki/Counting_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func getK(arr []int) int {
    if len(arr) == 0 {
        return 1
    }

    k := arr[0]
    
    for _, v := range arr {
        if v > k {
            k = v
        }
    }
    
    return k+1
}

func main(){
    arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")

    k := getK(arr)
    array_of_counts := make([]int, k)
    
    for i:= 0; i < len(arr); i++ {
        array_of_counts[arr[i]] += 1
    }

    for i, j := 0, 0; i < k; i++ {
        for {
            if array_of_counts[i] > 0 {
                arr[j] = i
                j += 1
                array_of_counts[i] -= 1
                continue
            }
            break
        }
    }
    
    fmt.Println("Sorted array is: ", arr)
}
