Namaste Multi-Greetings Program:
    -This Program takes name from user and Greets them "Namaste $UserName" to everyone individually.


Working of Multi Greetings Program:

    Package Requirements:
        -github.com/briandowns/spinner
        -bufio
        -fmt
        -os
        -strings
        -time

    files:
        -main.go
        -multiGreetings.go
        -getChar.go

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
    theory:
        -Website Geeksforgeeks:
            --https://www.geeksforgeeks.org/how-to-map-a-rune-to-uppercase-in-golang/?ref=ml_lbp
    video:
        -Channel name Golang Dojo (YouTube.com):
            --https://www.youtube.com/watch?v=7isCXLWPTqI