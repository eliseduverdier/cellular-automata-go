# Cellular Automata Documentation

This page describes how the project works, mainly what are the maths behind the generation or 1st and 2nd order elementary cellular automata.

## What's that ?

We call the image below a **elementary cellular automata** :

![s2-o1-r73](https://user-images.githubusercontent.com/4543769/131231053-411bac12-cb31-4ca4-acbd-ecf27a4764ad.png)

The first row of pixels at the top is chosen by the software (here, random dots of certain colors, black or white), and the rows below are calculated from the one above.

## Nice! How can I make them ?
In the case of a automata with 2 states and of the first order (only the preceding row is used), we can represent each case like this:

| cells above  | □□□ | □□■ | □■□ | □■■ | ■□□ | ■□■ | ■■□ | ■■■ |
| ------------ | --- | --- | --- | --- | --- | --- | --- | --- | 
| cell below   |  □  |  □  |  □  |  ■  |  ■  |  ■  |  ■  |  □  |

In this pattern, it means that "if the three cells above are all white, then create a white cell", etc for each pattern encountered. Now that see see the logic, we can replace the white/blacks cells by zeros and one, an only use these binaries number by their decimal representation. See, each number from 0 to 7 are written in the "cells above" row.

The "cell below" lines describes the number `00011110`, translated as `30` in decimal. So here, we have our rule **30**.

For 3 cells, we have (2^3) **8** combinaisons of cells possible. And as there is 8 combinaisons, we have (2^8) **256** possibles rules, each corresponding to an different pattern.

---

I presented here the *elementary* cellular automaton principle, but there are many more. We can complexify things by having more states ; we can also depends on more cells, not from the row before only, but the one before (they are called 2nd order CA, and some are reversible !).

They can also start not as a line but as a plan, and each cell changing itself according to their 8 surronding cells, thus creating the most famous of automata, the [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

 > aside note: Philosophically, the game of life is an endless and fascinating subject. Starting from zeros, ones, a simple rule, they can simulate whole turing-complete systems. That's only a little grid with moving dots, and they can create detailed and complex patterns, compute stuff they didn't even were made for!
 > 
 > And what happens outside the grid? What if an infinite game of life was a whole universe in itself? Like our neurons are just cells, firing or not at each other in a big globby-cluster, making a brain, creating complex and psychologically subtle behaviors in their host. 
 > 
 > Which in turn, waste their time endlessly pondering the nature of those automata and the universe, thanks to some little dots.

---

## Yeah right so were's the code ?

We we'll want to encode the array above in our code like this: `[ 0, 0, 0, 1, 1, 1, 1, 0 ]`.

And here is a solution to convert any rule number into its corresponding array  (I'm gonna use pseudo-go-code for clarity):
```go
function RuleToArray(states, order, ruleNumber) -> []int
	ruleInBaseN := ruleNumber->convertToBase(states)  // 30 becomes 11110
	reversedRule := Reverse(ruleInBaseN)              // which becomes [0,1,1,1,1]

	maxRule := states ^ (order + 2)                   // states ^ nb of preceding cells
	reversedRule.appendWith(0).until(maxRule)         // and we have [0,1,1,1,1,0,0,0]

	return ruleArray
```

Now we can work on creating our automata! So we'll need a few steps, summarized here:

```go
function (c CellularAutomata) GetMatrix() -> [][]int
	matrix := [][]                                   // initialize and empty matrix
	matrix[0] = c.getFirstLine()                     // populate the first line (randomly or not)

	for i:= range matrix {
		matrix[i+1] = c.getNextLine(matrix, i)   // populate each next row
	}
	
	return matrix
```

and in turn, `getNextLine()` will parse each pixel of the row and call `newCell()` on each. This is where our `ruleArray` gonna be useful.

```go
func (c CellularAutomata) newCell(position, currentLineIndex, matrix) -> int
	sumOfBaseCells := matrix[currentLineIndex][position] * 10               // _1_

	// handle diagonals cells on the sides: loop to the other side           // 111
	if position == 0  // first
		sumOfBaseCells += matrix[currentLineIndex][c.Columns - 1] * 100
		               + matrix[currentLineIndex][1]
	else if position == c.Columns-1  // last
		sumOfBaseCells += matrix[currentLineIndex][position - 1] * 100
		               + matrix[currentLineIndex][0]
	else
		sumOfBaseCells += matrix[currentLineIndex][position - 1] * 100
		               + matrix[currentLineIndex][position + 1]

	if c.Order > 1 && currentLineIndex > 1 { // 2nd order: also add center cell from n-2 line
		sumOfBaseCells += matrix[currentLineIndex - 1][position] * 1000

	index := sumOfBaseCells->convertFromBase(c.States)->toBase(10)         // converting our sum as 1100 to 12

	ruleArray := RuleToArray(c.States, c.Order, c.RuleNumber)              // fetching our mapping ruleArray

	return ruleArray[index]                                                // and returning the new cell state
}
```


You can see in details the all functions [here for the rule calculation](https://github.com/eliseduverdier/cellular-automata-go/blob/main/src/automata/rule_generator.go) and [there for the automata generation](https://github.com/eliseduverdier/cellular-automata-go/blob/main/src/automata/cellular_automata.go#L64).
