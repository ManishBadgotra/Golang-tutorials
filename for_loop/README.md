For-loop is an only loop that is available in Golang. There is no While or Do-While loop but we have functionality to use for-loop in a way While loop works.

Loop is helpful where we have to repeat our specific scenario again and again continously.

There are different ways to use for-loop in Go.
    -Iteration
    -for-loop with Range
    -for-loop as while loop



Working in for-loop:

    Package Requirements:
        -github.com/briandowns/spinner
        -bufio
        -fmt
        -os
        -strings
        -time

    Files:
        -iteration_in_for_loop.go
        -loop_over_string.go
        -reverse_iteration_in_for_loop.go
        -while.go

    Working of files individually:
        -main.go:
            Tells what this program is for.
            And asks name input from user in loop then appends it in slice
            after appending the given name it asks user if they want to continue Y/n and if user enters Y
            it repeats the program or else ends for-loop and calls multiGreetings function.

        -multiGreetings.go:
            Takes slice as an parameters and loops through it for Greeting individual names that is given by user.

        -getChar.go:
            This function takes two parameters str and index
            str that is given by user and
            index which is set to 0 in main file.
            Because we only need 0th index value from str(string)
            then slice of "Rune" helps in changing that Unicode code points to Characters.(Rune is from UTF-8 encoding)
            Because Unicode code points is a superset of ASCII which only contains 128 characters from 0 to 127. 


Resources:
    Theory:
        -Website JavaTPoint:
            --https://www.javatpoint.com/go-for
    Video:
        -Channel name Hitesh Choudhary (YouTube.com):
            --https://www.youtube.com/watch?v=0A5fReZUdRk