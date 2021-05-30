### Function

* Function can be declarated as below
  ```go
    func func_name(args) return_types {
        // function body
        return (returning values)
    }

    Ex:
        func newCard() string {
        	card := "Five of Diamonds"
	        return card
        }
  ```
  * `func` : Is a keyword to define a function
  * `func_name`: Define a function with the name. In ex: `newCard`
  * `return_types`: When executed, function will return values of the specified return types. In ex: `string`

#### Questions

1. If a function returns a value, do we have to annotate, or mark, the function with the type of value that is being returned?
    * Ans) `Yes`. Every function that returns a value must indicate what type of value it is returning.
2. The Go compiler is presenting an error message about the following function.  What should we do to fix the error?
    ```go
        func getSize() {
           return "30 meters"
        }
    ```
    * Ans) Add a return type of `string` to the function, like `func getSize() string {..}`
3. Is the following function valid?

    ```go
        func estimatePi() float64 {
            return 3.14
        }
    ```
    * Ans) `Yes`

4. Is the following code valid?  Why or why not?

    ```go
        package main

        import "fmt"

        func main() {
            fmt.Println(getTitle())
        }

        func getTitle() string {
            return "Harry Potter"
        }
    ```
    * Ans) `Yes`

5. Here's a tough one!  Imagine we have two files in the same folder with the same package name.  Will the following code successfully compile?  This might take a little experimentation on your side.  If you do try this out yourself, run your code with the command go run main.go state.go .
    ```go
        In `main.go`:

        package main

        func main() {
            printState()
        }


        In a separate file called `state.go`:

        package main

        import "fmt"

        func printState() {
            fmt.Println("California")
        }

    ```
    * `Yes`, Files in the same package can freely call functions defined in other files.
