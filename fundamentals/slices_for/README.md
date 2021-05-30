### Slices


* In Go we can hold list of elements in `Array or Slice`
  * `Array` : Fixed Length list of elements
  * `Slice` : An `Array` that can grow or shrink
  * ## Note : Every element in a slice must be of same type


#### Declaring slice
    ```go
        var_name := []var_type{elements}

        Ex:
        cards := []string{"Five Of Diamonds", "Ace of Diamonds"}
    ```
    * `var_name` : Defining slice name (in Ex : cards)
    * `[]var_type`: specifing that it's an slice (in ex :  it is slice of strings)
    * `{...}` : Population slice at the time of declaration


#### Appending element to slice
    ```go

        new_slice := append(existing_slice, element_to_add)

        Ex: newSetOfCards := append(cards, "Six of Diamonds")
            cards = append(cards, "Six of Diamonds")
    ```
    * `append` : The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself.


#### Iterating through Slice

    ```go

        for index, element := range elements {
            //body
        }

        Ex:
        for index, card := range cards {
		    fmt.Println(index, card)
	    }
    ```
    * `index` : index of this element in the array
    * `element` : Current element we're iterating over
    * `range elements` : Take this slice of 'elements' and loop over it
    * `:=` : Every time define new index&element variables with new set of values and throw away past variable values.

#### SubSets of Slice
* Slices are `zero-indexed`
* Syntax to extract subset from slice
  * `slice[startIndex:endIndex]`
    * `startIndex` : Inclusive, if we don't specify it will subset elements from index 0
    * `endIndex`: Exclusive, if we don't specify it will subset elements till the last index

    ```go
        cardValues := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	    index2Value := cardValues[2] // 2

        first5Elements := cardValues[0:5] // [Ace 1 2 3 4]
        first5Elements2 := cardValues[:5] // [Ace 1 2 3 4]
        after5thElement := cardValues[5:len(cardValues)] // [5 6 7 8 9 10 Jack Queen King]
        after5thElement2 := cardValues[5:] //  [5 6 7 8 9 10 Jack Queen King]
    ```

### Questions
1. How will you define a slice where each element in it is of type int?
   *   `[]int{}`
2. Is the following code valid?
    ```go
        colors := []strings{"Red", "Yellow", "Blue"}
    ```
    * `No`. This is probably a typo - the author probably meant to write `string` instead of `strings`

3. Would the following code compile successfully?  Try it out yourself!
    ```go
        for index, card := range cards {
           fmt.Println(card)
        }
    ```
    * `No`, because every variable we declare must be used in our code. In this case, `index` is not being used. To deal the unused `index` varaible we should define as `_ instead of index`
    ```go
        for _, card := range cards {
           fmt.Println(card)
        }
    ```
4. Can a slice have both values of type 'int' and of type 'string' in it?
   *   `No`, because a slice can only have one type of value in it.