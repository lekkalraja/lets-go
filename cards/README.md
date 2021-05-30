## Project Overview

#### Project Name : Cards
##### Behaviour (functions)
1. `newDeck` -> Create a list of playing cards, Essentially an array of strings
2. `print` -> Log out the contents of a deck of cards
3. `shuffle` -> Shuffles all the cards in a deck
4. `deal` -> Create a "hand" of cards.
5. `saveToFile` -> Save a list of cards to a file on the local machine
6. `newDeckFromFile` -> Load a list of cards from the local machine


### Questions

1. Think back to our current "cards" program, where we have the following code.

    ```go
        func main() {
            cards := newDeck()

            hand, remainingCards := deal(cards, 5)

            hand.print()
            remainingCards.print()
        }
    ```
    After calling `deal` and passing in `cards`, does the list of strings that the `cards` variable point at change?  In other words, did we modify the `cards` slice by calling `deal`?

    Ans) `No`, `cards` will be the same before and after calling `deal`. We created two new references that point at subsections of the `cards` slice. We never directly modified the slice that `cards` is pointing at.