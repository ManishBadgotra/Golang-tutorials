Switch case is similar to if-else-if chain in working it executes 1 out of many cases.

syntax of Switch case is 

    switch variable {
        case 1:
            // what you want to execute when case 1 matches
        case 2:
            // executes case 2 when it gets matched.
        .
        .
        .// as many cases as you want
        .
        .
        default:
            // default case will get execute when no case get matched.
    }

-We do not need break statement after every case to exit after execution of matched case like we are required to enter in other languages.

-Number after each case keyword checks for number, you can even match for individual characters like 'a' or string like "String".