## Variable Declarations

### 2 types of declaring a variable

1. `var var_name var_type = var_value`
   * `var card string = "Ace of Cards"`
     1. `var` : It's a keyword to create a new variable
     2. `var_name` (card) : The name of the variable (in ex: it's card)
     3. `var_type` (string) : The data type of a variable (in ex: Only a "string" will ever be assigned to variable)
     4. `var_value` : Value of the variable (in ex: Assigning the value "Ace of Cards"to variable)

2. `var_name := var_value` (Shorthand sytax)
    * Can create a new variable by omitting `var_type` and `var`(keyword) and using special character `:=`
    * `:=` -> Can infer type by using the right side value of `=` called `type inference`

###### Note1 : Should not use `:= (It's for only declaration)` for variable re-assignment, should use only `=`
###### Note2 : `:=` can use only to local variable declarations not for global variable declarations

##### Basic Go Types
- bool
- string
- int
- float64

###### Questions
1. Is the following a valid way of initializing and assigning a value to a variable?
   * `var bookTitle string = "Harry Potter"`
      * Ans) `Yes`
2. Is the following a valid way of initializing an assigning a value to a variable?
   * fruitCount := 5
     * Ans) `Yes`
3. After running the following code, Go will assume that the variable quizQuestionCount is of what type?
   * `quizQuestionCount := 10`
     * Ans) `Integer`
4. Will the following code compile?  Why or why not?
    ```go
      paperColor := "Green"
      paperColor := "Blue"
    ```
    * `No`, because a variable can only be initialized one time. In this case, the `:=` operator is being used to initialize `paperColor` two times

5. Are the two lines following ways of initializing the variable 'pi' equivalent?
    ```go
      pi := 3.14
      var pi float = 3.14
    ```
    * `Yes`
6. Is the following code valid?

    ```go
      package main

      import "fmt"

      deckSize := 20

      func main() {
        fmt.Println(deckSize)
      }
    ```
    * `No`
7. We might be able to initialize a variable and then later assign a variable to it.  Is the following code valid?
    ```go
      package main

      import "fmt"

      func main() {
        var deckSize int
        deckSize = 52
        fmt.Println(deckSize)
      }
    ```
    * `Yes`
8. Is the following code valid?

    ```go
      package main

      import "fmt"

      var deckSize int

      func main() {
        deckSize = 50
        fmt.Println(deckSize)
      }
    ```
    * `Yes`
9. Is the following code valid?  Why or why not?

    ```go
      package main

      import "fmt"

      func main() {
        deckSize = 52
        fmt.Println(deckSize)
      }
    ```
    * `No`, because `deckSize` is assigned a value before it is initialized.