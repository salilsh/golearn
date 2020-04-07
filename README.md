`Following is taken from Todd's doc. This is a list of exercises you'd complete if you complete all the sections of the Todd's course. Solutions linked below are Todd's. For my solutions, look at the ninja level and the exercise number, and navigate through the folders in the repo.`

`####################################`

# **Exercises - Ninja Level 1 (variables, values, &amp; type)**

Contribute your code

As you work through these hands on exercises, if you create code which you would like to share with the rest of the class, tweet me the link ( [https://twitter.com/Todd\_McLeod](https://twitter.com/Todd_McLeod) ) and I will add it to our course outline.

**video: 016b**

Hands-on exercise #1

1. Using the short declaration operator, **ASSIGN** these **VALUES** to **VARIABLES** with the IDENTIFIERS &quot;x&quot; and &quot;y&quot; and &quot;z&quot;
  1. 42
  2. &quot;James Bond&quot;
  3. true
2. Now print the values stored in those variables using
  1. a single print statement
  2. multiple print statements

**code: here&#39;s the solution:** [**https://play.golang.org/p/yYXnWXIQNa**](https://play.golang.org/p/yYXnWXIQNa)

**video: 017**

Hands-on exercise #2

1. Use var to **DECLARE** three **VARIABLES**. The variables should have package level scope. Do not assign **VALUES** to the variables. Use the following **IDENTIFIERS** for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that **TYPE** ).
  1. identifier &quot;x&quot; type int
  2. identifier &quot;y&quot; type string
  3. identifier &quot;z&quot; type bool
2. in func main
  1. print out the values for each identifier
  2. The compiler assigned values to the variables. What are these values called?

**code: here&#39;s the solution:** [**https://play.golang.org/p/jzHwSlles9**](https://play.golang.org/p/jzHwSlles9)

**video: 018**

Hands-on exercise #3

Using the code from the previous exercise,

1. At the package level scope, assign the following values to the three variables
  1. for x assign 42
  2. for y assign &quot;James Bond&quot;
  3. for z assign true
2. in func main
  1. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER &quot;s&quot;
  2. print out the value stored by variable &quot;s&quot;

**code: here&#39;s the solution:** [**https://play.golang.org/p/QFctSQB\_h3**](https://play.golang.org/p/QFctSQB_h3)

**video: 019**

Hands-on exercise #4

- FYI - nice documentation and new terminology &quot; **underlying type**&quot;
  - [https://golang.org/ref/spec#Types](https://golang.org/ref/spec#Types)

For this exercise

1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER &quot;x&quot; using the &quot;VAR&quot; keyword
3. in func main
  1. print out the value of the variable &quot;x&quot;
  2. print out the type of the variable &quot;x&quot;
  3. assign 42 to the VARIABLE &quot;x&quot; using the &quot;=&quot; OPERATOR
  4. print out the value of the variable &quot;x&quot;

**code: here&#39;s the solution:** [**https://play.golang.org/p/snm4WuuYmG**](https://play.golang.org/p/snm4WuuYmG)

**video: 020**

Hands-on exercise #5

Building on the code from the previous example

1. at the package level scope, using the &quot;var&quot; keyword, create a VARIABLE with the IDENTIFIER &quot;y&quot;. The variable should be of the UNDERLYING TYPE of your custom TYPE &quot;x&quot;

1. in func main
  1. this should already be done
    1. print out the value of the variable &quot;x&quot;
    2. print out the type of the variable &quot;x&quot;
    3. assign your own VALUE to the VARIABLE &quot;x&quot; using the &quot;=&quot; OPERATOR
    4. print out the value of the variable &quot;x&quot;
  2. now do this
    1. now use CONVERSION to convert the TYPE of the VALUE stored in &quot;x&quot; to the UNDERLYING TYPE
      1. then use the &quot;=&quot; operator to ASSIGN that value to &quot;y&quot;
      2. print out the value stored in &quot;y&quot;
      3. print out the type of &quot;y&quot;

**code: here&#39;s the solution:** [**https://play.golang.org/p/cj8RrYgBOD**](https://play.golang.org/p/cj8RrYgBOD)

**video: 021**

Hands-on exercise #6

[Take this quiz.](https://goo.gl/forms/dfwmTuYTe5ox8nyt1)

**video: 022**

# **Exercises - Ninja Level 2 (programming fundamentals)**

Hands-on exercise #1

Write a program that prints a number in **decimal, binary, and hex**

solution: [https://play.golang.org/p/bAQxcEuK8O](https://play.golang.org/p/bAQxcEuK8O)

**video: 031**

Hands-on exercise #2

Using the following operators, write expressions and assign their values to variables:

1.
  1. **a.**** ==**
  2. **b.**** \&lt;=**
  3. **c.**** \&gt;=**
  4. **d.****!=**
  5. **e.**** \&lt;**
  6. **f.**** \&gt;**

Now print each of the variables.

solution: [https://play.golang.org/p/76R-poSzaY](https://play.golang.org/p/76R-poSzaY)

**video: 032**

Hands-on exercise #3

Create **TYPED and UNTYPED constants**. Print the values of the constants.

solution: [https://play.golang.org/p/NutvJXWUx2](https://play.golang.org/p/NutvJXWUx2)

**video: 033**

Hands-on exercise #4

Write a program that

- assigns an int to a variable
- prints that int in decimal, binary, and hex
- **shifts the bits** of that int over 1 position to the left, and assigns that to a variable
- prints that variable in decimal, binary, and hex

solution: [https://play.golang.org/p/Ms964T8SbH](https://play.golang.org/p/Ms964T8SbH)

**video: 034**

Hands-on exercise #5

Create a variable of type string using a **raw string literal**. Print it.

solution: [https://play.golang.org/p/dLy36A-V-w](https://play.golang.org/p/dLy36A-V-w)

**video: 035**

Hands-on exercise #6

Using **iota** , create 4 constants for the **NEXT** 4 years. Print the constant values.

solution: [https://play.golang.org/p/MDLF3v5EGT](https://play.golang.org/p/MDLF3v5EGT)

**video: 036**

Hands-on exercise #7

[**take this quiz**](https://goo.gl/forms/xoDXOWAHK0A4WN2V2)

**video: 037**

# **Exercises - Ninja Level 3 (Control flow)**

Hands-on exercise #1

Print every number from 1 to 10,000

solution: [https://play.golang.org/p/voDiuiDGGw](https://play.golang.org/p/voDiuiDGGw)

**video: 050**

Hands-on exercise #2

Print every rune code point of the uppercase alphabet three times. Your output should look like this:

65

 U+0041 &#39;A&#39;

 U+0041 &#39;A&#39;

U+0041 &#39;A&#39;

66

 U+0042 &#39;B&#39;

 U+0042 &#39;B&#39;

 U+0042 &#39;B&#39;

 … through the rest of the alphabet characters

solution: [https://play.golang.org/p/1OjnCX1D5H](https://play.golang.org/p/1OjnCX1D5H)

**video: 051**

Hands-on exercise #3

Create a for loop using this syntax

- for condition { }

Have it print out the years you have been alive.

solution: [https://play.golang.org/p/tnyqBPJ-i5](https://play.golang.org/p/tnyqBPJ-i5)

**video: 052**

Hands-on exercise #4

Create a for loop using this syntax

- for { }

Have it print out the years you have been alive.

solution: [https://play.golang.org/p/9VpnB-I1Pz](https://play.golang.org/p/9VpnB-I1Pz)

**video: 053**

Hands-on exercise #5

Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

solution: [https://play.golang.org/p/ohfJOW9euy](https://play.golang.org/p/ohfJOW9euy)

**video: 054**

Hands-on exercise #6

Create a program that shows the &quot;if statement&quot; in action.

solution: [https://play.golang.org/p/DpZ\_FLfn5s](https://play.golang.org/p/DpZ_FLfn5s)

**video: 055**

Hands-on exercise #7

Building on the previous hands-on exercise, create a program that uses &quot;else if&quot; and &quot;else&quot;.

solution: [https://play.golang.org/p/IDnrJpE7vT](https://play.golang.org/p/IDnrJpE7vT)

**video: 056**

Hands-on exercise #8

Create a program that uses a switch statement with no switch expression specified.

solution: [https://play.golang.org/p/YpPgkWGqKY](https://play.golang.org/p/YpPgkWGqKY)

**video: 057**

Hands-on exercise #9

Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER &quot;favSport&quot;.

solution: [https://play.golang.org/p/h-8FnjpzWB](https://play.golang.org/p/h-8FnjpzWB)

**video: 058**

Hands-on exercise #10

Write down what these print:

- fmt.Println( **true &amp;&amp; true** )
- fmt.Println( **true &amp;&amp; false** )
- fmt.Println( **true || true** )
- fmt.Println( **true || false** )
- fmt.Println( **!true** )

solution:

**video: 059**

# **Exercises - Ninja Level 4 (Grouping data)**

Hands-on exercise #1

- Using a COMPOSITE LITERAL:

- create an ARRAY which holds 5 VALUES of TYPE int
- assign VALUES to each index position.

- Range over the array and print the values out.
- Using format printing
  - print out the TYPE of the array

**solution:** [**https://play.golang.org/p/tD0SzV3hdf**](https://play.golang.org/p/tD0SzV3hdf)

**video: 071**

Hands-on exercise #2

- Using a COMPOSITE LITERAL:

- create a SLICE of TYPE int
- assign 10 VALUES

- Range over the slice and print the values out.
- Using format printing
  - print out the TYPE of the slice

**solution:** [**https://play.golang.org/p/sAQeFB7DIs**](https://play.golang.org/p/sAQeFB7DIs)

**video: 072**

Hands-on exercise #3

Using the code from the previous example, use SLICING to create the following new slices which are then printed:

- [42 43 44 45 46]
- [47 48 49 50 51]
- [44 45 46 47 48]
- [43 44 45 46 47]

**solution:** [**https://play.golang.org/p/SGfiULXzAB**](https://play.golang.org/p/SGfiULXzAB)

**video: 073**

Hands-on exercise #4

Follow these steps:

- start with this slice
  - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- append to that slice this value
  - 52
- print out the slice
- in ONE STATEMENT append to that slice these values
  - 53
  - 54
  - 55
- print out the slice
- append to the slice this slice
  - y := []int{56, 57, 58, 59, 60}
- print out the slice

**solution:** [**https://play.golang.org/p/QUYhtSBaDS**](https://play.golang.org/p/QUYhtSBaDS)

**video: 074**

Hands-on exercise #5

To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:

- start with this slice
  - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- use APPEND &amp; SLICING to get these values here which you should ASSIGN to a variable &quot;y&quot; and then print:

-
  - [42, 43, 44, 48, 49, 50, 51]

**solution:** [**https://play.golang.org/p/u8zpHLfgSE**](https://play.golang.org/p/u8zpHLfgSE)

**just for fun:**

- **●●** [**https://goo.gl/frz786**](https://goo.gl/frz786)
- **●●** [**https://www.youtube.com/channel/UCJ8YIwWQCO7hMiqpOw2ZLFw**](https://www.youtube.com/channel/UCJ8YIwWQCO7hMiqpOw2ZLFw)

**video: 075**

Hands-on exercise #6

Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

**solution:**

- **●●**** the solution **** shown in the video was incorrect **** - this one is incorrect**
  - **○○** [**https://play.golang.org/p/tRKQDQuQCE**](https://play.golang.org/p/tRKQDQuQCE)
  - **○○**** I used assignment and a composite literal instead of append**
- **●●**** here is **** the correct solution**
  - **○○** [**https://play.golang.org/p/dzxZh4lhEpT**](https://play.golang.org/p/dzxZh4lhEpT)

**video: 076**

Hands-on exercise #7

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

&quot;James&quot;, &quot;Bond&quot;, &quot;Shaken, not stirred&quot;

&quot;Miss&quot;, &quot;Moneypenny&quot;, &quot;Helloooooo, James.&quot;

Range over the records, then range over the data in each record.

**solution:** [**https://play.golang.org/p/FMM4c2PhZg**](https://play.golang.org/p/FMM4c2PhZg)

**video:  077**

Hands-on exercise #8

Create a map with a key of TYPE string which is a person&#39;s &quot;last\_first&quot; name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

 `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`

`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`

`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

**solution:** [**https://play.golang.org/p/nTzSlRa9\_A**](https://play.golang.org/p/nTzSlRa9_A)

**video: 078**

Hands-on exercise #9

Using the code from the previous example, add a record to your map. Now print the map out using the &quot;range&quot; loop

**solution:** [**https://play.golang.org/p/\_CkxAhRrDJ**](https://play.golang.org/p/_CkxAhRrDJ)

**video: 079**

Hands-on exercise #10

Using the code from the previous example, delete a record from your map. Now print the map out using the &quot;range&quot; loop

**solution:** [**https://play.golang.org/p/TYl5EbjoeC**](https://play.golang.org/p/TYl5EbjoeC)

**video: 080**

# **Exercises - Ninja Level 5 (Structs)**

Hands-on exercise #1

Create your own type &quot;person&quot; which will have an underlying type of &quot;struct&quot; so that it can store the following data:

- first name
- last name
- favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

**solution:**

- **●●** [**https://play.golang.org/p/ouRHmH\_POg**](https://play.golang.org/p/ouRHmH_POg)

**video: 086**

Hands-on exercise #2

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

**solution:** [**https://play.golang.org/p/RvvLyAMvGo**](https://play.golang.org/p/RvvLyAMvGo)

**video: 087**

Hands-on exercise #3

- Create a new type: vehicle.
  - The underlying type is a struct.
  - The fields:
    - ■■doors
    - ■■color
- Create two new types: truck &amp; sedan.
  - The underlying type of each of these new types is a struct.
  - Embed the &quot;vehicle&quot; type in both truck &amp; sedan.
  - Give truck the field &quot;fourWheel&quot; which will be set to bool.
  - Give sedan the field &quot;luxury&quot; which will be set to bool. solution
- Using the vehicle, truck, and sedan structs:
  - using a composite literal, create a value of type truck and assign values to the fields;
  - using a composite literal, create a value of type sedan and assign values to the fields.
- Print out each of these values.
- Print out a single field from each of these values.

**solution:** [**https://play.golang.org/p/PrTtTv\_vVO**](https://play.golang.org/p/PrTtTv_vVO)

**video: 088**

Hands-on exercise #4

Create and use an anonymous struct.

**solution:** [**https://play.golang.org/p/otBHFs-lPp**](https://play.golang.org/p/otBHFs-lPp)

**video: 089**

# **Exercises - Ninja Level 6 (functions)**

Hands-on exercise #1

- Review
  - functions
  - purpose of functions
    - ■■abstract code
    - ■■code reusability
      - DRY - Don&#39;t Repeat Yourself
  - func, receiver, identifier, params, returns
  - parameters vs arguments
  - variadic funcs
    - ■■multiple &quot;variadic&quot; params
    - ■■multiple &quot;variadic&quot; args
  - returns
    - ■■multiple returns
    - ■■named returns - yuck!
  - func expressions
    - ■■assigning a func to a variable
  - callbacks
    - ■■passing a func into another func as an argument
  - closure
    - ■■one scope enclosing another
    - ■■variables declared in the outer scope are accessible in inner scopes
    - ■■closure helps us limit the scope of variables
  - recursion
    - ■■a func that calls itself
    - ■■factorial
- life philosophy
  - focus on what&#39;s important; not upon what&#39;s urgent
- Hands on exercise
  - create a func with the identifier foo that returns an int
  - create a func with the identifier bar that returns an int and a string
  - call both funcs
  - print out their results

**code:** [**https://play.golang.org/p/owEJNF5WAD**](https://play.golang.org/p/owEJNF5WAD)

**video: 102**

Hands-on exercise #2

- create a func with the identifier foo that
  - takes in a variadic parameter of type int
  - pass in a value of type []int into your func (unfurl the []int)
  - returns the sum of all values of type int passed in
- create a func with the identifier bar that
  - takes in a parameter of type []int
  - returns the sum of all values of type int passed in

**code:** [**https://play.golang.org/p/B0yRxtBQPD**](https://play.golang.org/p/B0yRxtBQPD)

**video: 103**

Hands-on exercise #3

Use the &quot;defer&quot; keyword to show that a deferred func runs after the func containing it exits.

**code:** [**https://play.golang.org/p/XluEuUD0Nw**](https://play.golang.org/p/XluEuUD0Nw)

**video: 104**

Hands-on exercise #4

- Create a user defined struct with
  - the identifier &quot;person&quot;
  - the fields:
    - ■■first
    - ■■last
    - ■■age
- attach a method to type person with
  - the identifier &quot;speak&quot;
  - the method should have the person say their name and age
- create a value of type person
- call the method from the value of type person

**code:** [**https://play.golang.org/p/NnXyWdqbbw**](https://play.golang.org/p/NnXyWdqbbw)

**video: 105**

Hands-on exercise #5

- create a type **SQUARE**
- create a type **CIRCLE**
- attach a method to each that calculates **AREA** and returns it
  - circle area= π r 2
  - square area = L \* W
- create a type **SHAPE** that defines an interface as anything that has the **AREA** method
- create a func **INFO** which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle

**code:** [**https://play.golang.org/p/NGGikHNvHv**](https://play.golang.org/p/NGGikHNvHv)

**video: 106**

Hands-on exercise #6

- Build and use an anonymous func

**code:** [**https://play.golang.org/p/DQX3xEIcRe**](https://play.golang.org/p/DQX3xEIcRe) ** **

**video: 107**

Hands-on exercise #7

- Assign a func to a variable, then call that func

**code:** [**https://play.golang.org/p/\_Qu7ZAyFDH**](https://play.golang.org/p/_Qu7ZAyFDH)

**video: 108**

Hands-on exercise #8

- Create a func which returns a func
- assign the returned func to a variable
- call the returned func

**code:** [**https://play.golang.org/p/c2HwqVE1Rd**](https://play.golang.org/p/c2HwqVE1Rd)

**video: 109**

Hands-on exercise #9

A &quot;callback&quot; is when we pass a func into a func as an argument. For this exercise,

- pass a func into a func as an argument

**code:** [**https://play.golang.org/p/0yGYPKh1y7**](https://play.golang.org/p/0yGYPKh1y7)

**video: 110**

Hands-on exercise #10

Closure is when we have &quot;enclosed&quot; the scope of a variable in some code block. For this hands-on exercise, create a func which &quot;encloses&quot; the scope of a variable:

**code:** [**https://play.golang.org/p/a56uWtoFSL**](https://play.golang.org/p/a56uWtoFSL)

**video: 111**

Hands-on exercise #11

The best way to learn is to teach. For this hands-on exercise,

- choose one of the above exercises, or use the recursion example of factorial
- download, install, and get it running
  - [https://obsproject.com/](https://obsproject.com/)
- record a video of YOU teaching the topic
- upload the video to youtube
- share the video on twitter and tag me in it ( [https://twitter.com/Todd\_McLeod](https://twitter.com/Todd_McLeod) ) so that I can see it!

**video: 112**

# **Exercises - Ninja Level 7 (Pointers)**

Hands-on exercise #1

- Create a value and assign it to a variable.
- Print the address of that value.

**code:** [**https://play.golang.org/p/57kW8xd0qT**](https://play.golang.org/p/57kW8xd0qT)

**video: 116**

Hands-on exercise #2

- create a person struct
- create a func called &quot;changeMe&quot; with a \*person as a parameter
  - change the value stored at the \*person address
    - **■■**** important**
      - to dereference a struct, use (\*value).field
        - p1.first
        - (\*p1).first
      - &quot;As an exception, if the type of x is a named pointer type and (\*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (\*x).f.&quot;
        - [https://golang.org/ref/spec#Selectors](https://golang.org/ref/spec#Selectors)
- create a value of type person
  - print out the value
- in func main
  - call &quot;changeMe&quot;
- in func main
  - print out the value

**code:** [**https://play.golang.org/p/AehM662HKS**](https://play.golang.org/p/AehM662HKS)

**video: 117**

# **Exercises - Ninja Level 8  (JSON)**

Hands-on exercise #1

Starting [with this code](https://play.golang.org/p/_fQUGm9Utvl), marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

**solution:** [**https://play.golang.org/p/8BK6PXj3aG**](https://play.golang.org/p/8BK6PXj3aG)

**video: 125**

Hands-on exercise #2

Starting [with this code](https://play.golang.org/p/b_UuCcZag9), unmarshal the JSON into a Go data structure. Hint: [https://mholt.github.io/json-to-go/](https://mholt.github.io/json-to-go/)

**code:**

- **●●**** code setup (just fyi, not needed for exercise):**
  - **○○** [**https://play.golang.org/p/nWPP5Z2Q4e**](https://play.golang.org/p/nWPP5Z2Q4e)
- **●●**** solution:**
  - **○○** [**https://play.golang.org/p/r8oSG8DIPR**](https://play.golang.org/p/r8oSG8DIPR)

**video: 126**

Hands-on exercise #3

Starting [with this code](https://play.golang.org/p/BVRZTdlUZ_), encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

**solution:** [**https://play.golang.org/p/ql90D1OQqd**](https://play.golang.org/p/ql90D1OQqd)

**video: 127**

Hands-on exercise #4

Starting [with this code](https://play.golang.org/p/H_q75mpmHW), sort the []int and []string for each person.

**solution:** [**https://play.golang.org/p/jz\_llY1aPp**](https://play.golang.org/p/jz_llY1aPp)

**video: 128**

Hands-on exercise #5

Starting [with this code,](https://play.golang.org/p/BVRZTdlUZ_) sort the []user by

- age
- last

Also sort each []string &quot;Sayings&quot; for each user

- print everything in a way that is pleasant

**solution:** [**https://play.golang.org/p/8RKkdtLl6w**](https://play.golang.org/p/8RKkdtLl6w)

# **Exercises - Ninja Level 9 (Concurrency)**

Hands-on exercise #1

- in addition to the main goroutine, launch two additional goroutines
  - each additional goroutine should print something out
- use waitgroups to make sure each goroutine finishes before your program exists

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 148**

Hands-on exercise #2

This exercise will reinforce our understanding of method sets:

- create a type person struct
- attach a method speak to type person using a pointer receiver
  - \*person
- create a type human interface
  - to implicitly implement the interface, a human must have the speak method
- create func &quot;saySomething&quot;
  - have it take in a human as a parameter
  - have it call the speak method
- show the following in your code
  - you CAN pass a value of type \*person into saySomething
  - you CANNOT pass a value of type person into saySomething
- here is a hint if you need some help
  - [https://play.golang.org/p/FAwcQbNtMG](https://play.golang.org/p/FAwcQbNtMG)

Receivers       Values

-----------------------------------------------

(t T)           T and \*T

(t \*T)          \*T

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 149**

Hands-on exercise #3

- Using goroutines, create an incrementer program
  - have a variable to hold the incrementer value
  - launch a bunch of goroutines
    - ■■each goroutine should
      - read the incrementer value
        - store it in a new variable
      - yield the processor with runtime.Gosched()
      - increment the new variable
      - write the value in the new variable back to the incrementer variable
- use waitgroups to wait for all of your goroutines to finish
- the above will create a race condition.
- Prove that it is a race condition by using the -race flag
- if you need help, here is a hint: [https://play.golang.org/p/FYGoflKQej](https://play.golang.org/p/FYGoflKQej)

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 150**

Hands-on exercise #4

Fix the race condition you created in the previous exercise by using a mutex

- it makes sense to remove runtime.Gosched()

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 151**

Hands-on exercise #5

Fix the race condition you created in exercise #4 by using package atomic

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 152**

Hands-on exercise #6

Create a program that prints out your OS and ARCH. Use the following commands to run it

- go run
- go build
- go install

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 153**

Hands-on exercise #7

Download [OBS](https://obsproject.com/). Create a screen recording teaching some aspect of the Go programming language. Some topics you might teach:

- [Why Go](https://docs.google.com/presentation/d/1x_mtMIBbxQU0zLmxgPSvZ0GFaRrfOOjneeNeahLQcto/edit?usp=sharing)
- installing go
- GOROOT &amp; GOPATH
  - environment variables
- hello world
- go commands
  - help
- variables
  - Short declaration operator
- constants
- loops
  - init, cond, post
  - break
  - continue
- functions
  - func (receiver) identifier(params) (returns) { code }
- methods
- interfaces
- method sets
- type
  - conversion
- concurrency vs parallelism
- goroutines
- waitgroups
- mutex
- atomic

After creating your recording, upload your video to youtube. Then tweet a link to your video and [tag me in the tweet](https://twitter.com/Todd_McLeod) so that I can see it.

**code:** [**https://github.com/GoesToEleven/go-programming**](https://github.com/GoesToEleven/go-programming)

**video: 154**

# **Exercises - Ninja Level 10 (Channels)**

Hands-on exercise #1

- get [this code](https://play.golang.org/p/j-EA6003P0) working:
  - **○○** using func literal, aka, anonymous self-executing func
    - **■■** solution: [https://play.golang.org/p/SHr3lpX4so](https://play.golang.org/p/SHr3lpX4so)
  - **○○** a buffered channel
    - **■■** solution: [https://play.golang.org/p/Y0Hx6IZc3U](https://play.golang.org/p/Y0Hx6IZc3U)

**video: 164**

Hands-on exercise #2

- Get this code running:
  - **○○** [**https://play.golang.org/p/oB-p3KMiH6**](https://play.golang.org/p/oB-p3KMiH6)
    - **■■**** solution:**[**https://play.golang.org/p/isnJ8hMMKg**](https://play.golang.org/p/isnJ8hMMKg)
  - **○○** [**https://play.golang.org/p/\_DBRueImEq**](https://play.golang.org/p/_DBRueImEq)
    - **■■**** solution:**[**https://play.golang.org/p/mgw750EPp4**](https://play.golang.org/p/mgw750EPp4)

**video: 165**

Hands-on exercise #3

- Starting with [this code](https://play.golang.org/p/sfyu4Is3FG), pull the values off the channel using a for range loop

**solution:** [**https://play.golang.org/p/D3N4Tq54SN**](https://play.golang.org/p/D3N4Tq54SN)

**video: 166**

Hands-on exercise #4

- Starting with [this code](https://play.golang.org/p/MvL6uamrJP), pull the values off the channel using a select statement

**solution:** [**https://play.golang.org/p/FulKBY5JNj**](https://play.golang.org/p/FulKBY5JNj) ** **

**video: 167**

Hands-on exercise #5

- Show the comma ok idiom starting with [this code.](https://play.golang.org/p/YHOMV9NYKK)

**solution:** [**https://play.golang.org/p/qh2ywLB5OG**](https://play.golang.org/p/qh2ywLB5OG)

**video: 168**

Hands-on exercise #6

- write a program that
  - **○○** puts 100 numbers to a channel
  - **○○** pull the numbers off the channel and print them

**solution:** [**https://play.golang.org/p/096Lk1BR7o**](https://play.golang.org/p/096Lk1BR7o)

**video: 169**

Hands-on exercise #7

- write a program that
  - **○○** launches 10 goroutines
    - **■■** each goroutine adds 10 numbers to a channel
  - **○○** pull the numbers off the channel and print them

**solutions:**

- **●●** [**https://play.golang.org/p/R-zqsKS03P**](https://play.golang.org/p/R-zqsKS03P)
- **●●** [**https://play.golang.org/p/quWnlwzs2z**](https://play.golang.org/p/quWnlwzs2z)

**student solutions:**

- **●●** [**https://twitter.com/mannion**](https://twitter.com/mannion)
  - **○○** [**https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f**](https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f)

**video: 170**

# **Exercises - Ninja Level 11 (Error handling)**

Hands-on exercise #1

Start with [this code](https://play.golang.org/p/3W69TH4nON). Instead of using the blank identifier, make sure the code is checking and handling the error.

**solution:**

- **●●** [**https://play.golang.org/p/tn8oiuL1Yn**](https://play.golang.org/p/tn8oiuL1Yn)

**video: 176**

Hands-on exercise #2

Start with [this code](https://play.golang.org/p/9a1IAWy5E6). Create a custom error message using &quot;fmt.Errorf&quot;.

**solution:**

- **●●** [**https://play.golang.org/p/HugU4HJEEO**](https://play.golang.org/p/HugU4HJEEO)
- **●●** [**https://play.golang.org/p/NII-lmGasj**](https://play.golang.org/p/NII-lmGasj)
- **●●** [**https://play.golang.org/p/Vo5kIoR-sG**](https://play.golang.org/p/Vo5kIoR-sG)

**video: 177**

Hands-on exercise #3

Create a struct &quot;customErr&quot; which implements the builtin.error interface. Create a func &quot;foo&quot; that has a value of type error as a parameter. Create a value of type &quot;customErr&quot; and pass it into &quot;foo&quot;. If you need a hint, [here is one](https://play.golang.org/p/L5QWV8-p11).

**solution:**

- **●●** [**https://play.golang.org/p/ixeowY2fd2**](https://play.golang.org/p/ixeowY2fd2)
- **●●**** assertion**
  - **○○** [**https://play.golang.org/p/pbl2kCYsM0**](https://play.golang.org/p/pbl2kCYsM0)
- **●●**** conversion**
  - **○○** [**https://play.golang.org/p/1ldiBdkdzA**](https://play.golang.org/p/1ldiBdkdzA)

**video: 178**

Hands-on exercise #4

Starting with [this code](https://play.golang.org/p/wlEM1tgfQD), use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your

- lat &quot;50.2289 N&quot;
- long &quot;99.4656 W&quot;

**solution:**

- **●●** [**https://play.golang.org/p/nsRxbDfkCh**](https://play.golang.org/p/nsRxbDfkCh)

**video: 179**