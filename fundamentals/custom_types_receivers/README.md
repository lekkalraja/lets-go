## Custom Types

### Person Object Oriented Approach

* In OO-Approach to represent real world objects we do have containers like `class` and define inside state and behaviours like below.

    ```java
        public class Person {
            private String name;

            public void printName() {
                System.out.println("Name of the Person " + name);
            }
        }

        public class App {
            public static void main(String... args) {
                Person person = new Person();
                person.printName();
            }
        }
    ```

* Whereas in `GoLang` we don't have oo features, so we do have some other alternative like extending types from basic go types and have receiver functions.

### Custom Types (Extended types)

    ```go
        Basic Go Types (string, int, float, array, map)
                    || we want to "extend" a base type
                    || and add some extra functionality to it
                    \/
             type person []string // Tell Go we want to create an array of strings and add a bunch of functions specifically made to work with it
                    ||
                    ||
                    \/
            Functions with "person" as a "receiver" // A function with a receiver is like a "method" - a function that belongs to an "instance"
    ```
* Above object oriented solution can be implemented in go like below

    ```go
        /* person.go */
        package main

        import "fmt"

        type Persons []string

        // print is a receiver function for Persons
        func (p Persons) print() {
            for index, person := range p {
                fmt.Println(index, person)
            }
        }

        /*main.go*/
        package main

        func main() {
            people := Persons{"Rama", "Seetha", "Radha", "Krishna"}

            people.print()
        }

    ```

### Receiver Functions

    ```go

        func (p Persons) print() {
            //....
        }

    ```
    * Any variable of type `Persons` now gets access to the `print` method
    * `p` : The actual copy of the `people` in the above example we're working with is available in the function as a variable called `p`
    * `Persons`: Every variable of type `Persons` can call this function on itself.

######     Note : It's a convention to have 1/2 letter varaible name as part of receiver (in ex: it's `p`)


### Questions

1. What would the following code print out?

    ```go
        package main

        import "fmt"

        type book string

        func (b book) printTitle() {
            fmt.Println(b)
        }

        func main() {
            var b book = "Harry Potter"
            b.printTitle()
        }
    ```
    * Ans) `"Harry Potter"`

2. Complete the sentence!
    By creating a new type with a function that has a receiver, we...

    * Ans) Are adding a `method` to any value of `that type`

3. In the following snippet, what does the variable 'ls' represent?

    ```go
        type laptopSize float

        func (ls laptopSize) getSizeOfLaptop() {
            return ls
        }
    ```
    * Ans) A `Value` of type `laptopSize`

4. Is the following code valid?

    ```go
        type laptopSize float64

        func (this laptopSize) getSizeOfLaptop() laptopSize {
            return this
        }
    ```
    * Ans) Yes, but it is breaking convention, Go avoids any mention of `this` or `self`. While the code is technically valid and will compile, we don't ever reference a receiver value as `this` or `self`.

5. What will the following program log out?

    ```go
        package main

        import "fmt"

        func main() {
        c := color("Red")

        fmt.Println(c.describe("is an awesome color"))
        }

        type color string

        func (c color) describe(description string) (string) {
        return string(c) + " " + description
        }
    ```
    * Ans) `Red is an awesome color`
    * Note : `describe` is a function with a receiver of type `color` that requires an argument of type `string`, then returns a value of type `string`