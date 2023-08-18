A pointer is a variable that stores the memory address of another variable. For getting an address of other variable we adds & (address-of operator) symbol before the variable's name whom's memry address we want.

example: var a int = 7
        var b *int = &a // in this we used * symbol to tell the compiler that this is an pointer variable.


With pointers, we can pass a reference to a variable (for example, as a parameter to a function), instead of passing a copy of the variable which can reduce memory usage and increase efficiency.

Pointer example in Go:
        package main  
        import (  
                "fmt"  
            )  
        func main() {  
            x:=10  
            changeX(&x)  
            fmt.Println(x)  // outpur is x = 0
        }  
        func changeX(x *int){  
            *x=0  
        }  