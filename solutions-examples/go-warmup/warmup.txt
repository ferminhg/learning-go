# Warmup

1. Create a program that prints out "Hello, world"

2. Store a `name` and an `age` into variables. Use a `uint8` for storing the age, and print them,
to look like this "John is years 32 years old"

2.1 Create a function named getAge that returns the age, and use it in the print statement

2.2 Rename the function to getPersonalInfo and have it return both the name and age

2.3 Print out just the name without the year.

3. Create a struct named `person` that contains the fields `name` and `age`.
Initialize it into a variable and print the text similar to item 2 ("John is years 32 years old")
Also try printing the struct instance directly using "%v".

3.1. Create a receiver method for the struct that prints out the name and age.

3.2. Create a receiver method that sets the age to an arbitrary value.
Check if the changes are applied with a both a value and pointer receiver.
If you're having trouble here, jump to number 4.

4. We have the following code

  x := 40

Print the address of the "x" variable (in other words the pointer of x).

4.1 Store the address of x to a variable named y. Use y to print the value 40 (basically dereference y).

5. Create a slice of strings containing the first 5 week days.
Iterate over the entries and print them ex. "Current day is Monday", etc.

5.1. On a new line, `append` onto the slice the weekend days.

5.2 If the current element index you're iterating over is greater than 5, print "It's weekend."