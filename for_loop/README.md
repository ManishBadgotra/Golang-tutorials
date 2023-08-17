For-loop is an only loop that is available in Golang. There is no While or Do-While loop but we have functionality to use for-loop in a way While loop works.

Loop is helpful where we have to repeat our specific scenario again and again continously.

There are different ways to use for-loop in Go.
    -Counter-controlled iteration
    -Condition-controlled iteration


Working in for-loop:

    Files:
        -iteration_in_for_loop.go
        -loop_over_string.go
        -reverse_iteration_in_for_loop.go
        -while.go

    Working of files individually:
        -iteration_in_for_loop.go:
            For-loop follows a condition as follow (declare; condition; increment) and in this we counter on i until i equals to 10, when i is not equal to 10 add comma after printing i or only i in fmt.Println().

        -loop_over_string.go:
            We can also loop over string by taking length of string. Because string is like slice in Golang. And we can do same counter-controlled like we did in iteration_in_for_loop.go with a change instead of printing i we have to print value on string's index in string[i] 

        -reverse_iteration_in_for_loop.go:
            We can reverse iterate too just instead of incrementing i, we have to decrement i. Remember to check for declaration value of i must be a value that can decrement.


        -while.go:
            in for-loop, if we want working of while loop than instead of declaring, condition, and increment|decrement in a single line we declare variable before loop, enter condition in for condition place without semicolon, and increment|decrement inside loop block.


Resources:
    Theory:
        -Website JavaTPoint:
            --https://www.javatpoint.com/go-for
    Video:
        -Channel name Hitesh Choudhary (YouTube.com):
            --https://www.youtube.com/watch?v=0A5fReZUdRk