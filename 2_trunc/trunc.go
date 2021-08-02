package main

import "fmt"

func main() {
    var fltNum float64;
    fmt.Println("Enter floating point number to be truncated : ");
    num, err := fmt.Scan(&fltNum);
    if(err!=nil){
        fmt.Println("Err : ", err, num);    
    }
    var intNum int64 = int64(fltNum);
    fmt.Println("Output : ", intNum);
}
