The Struct is known as Structures, Struct are like Class in Object-Oriented based language.
There is no such OOP like concepts but we can access these concepts in a Go way.

syntax:
    /*
        Below is a way to declaring struct.
    */

    type Employee struct {
	FirstName string
	LastName  string
	Age       int
    }

    /*
        Below is a function made with Struct.
    */
    func (emp Employee) Fullname() {
	fmt.Printf("Hello, %v\n", emp.FirstName+" "+emp.LastName)
	fmt.Printf("Wow, you are %d years old.\n", emp.Age)
    }

    /*
        Below is a function made with Struct.
    */
        

In this syntax some rules we have to follow which are:
    Capitalize first letter of struct name and variables inside it.
    Struct name must justify what it will be used for.
    If struct is inside the main file than it has nothing to do with package name but if we made an different file to save struct in it then. We will have to follow above rules strictly and then import our struct package in our main.go 
