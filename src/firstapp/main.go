package main

import "fmt"
import "strconv"
import "reflect"

// must be declared this way
// will not be 'exported'
var globalNumber int = 122

// will be exported because the name of the variable starts with a capital letter
var GlobalNumber int = 1000

// conventions
// acronyms should be in block letters
var theURL string = "http://google.com"

func main () {
  fmt.Println("Hello world")

  // variables
  // primitives are passed by value

  fmt.Println("\n\nVariables\n")

  var number int = 11
  fmt.Println(number)

  number_2 := 12
  fmt.Println(number_2)
  
  number_3 := 3.142
  fmt.Printf("%v %T\n", number_3, number_3)

  fmt.Println(globalNumber)
  
  // shadowing

  fmt.Println("\n\nShadowing\n")

  var globalNumber = 100

  fmt.Println(globalNumber)

  number_4 := int(number_3)
  fmt.Printf("%v\n", number_4)

  fmt.Println(strconv.Itoa(number_4))

  var state bool = true

  fmt.Println(state)

  // arrays
  // passed by value

  fmt.Println("\n\nArrays\n")

  var matrix [3][3]int
  fmt.Printf("matrix: %v\n", matrix)

  var nums = [...]int { 3, 2, 1 }
  fmt.Printf("nums: %v\n", nums)
  fmt.Printf("len(nums): %v\n", len(nums))
  fmt.Printf("cap(nums): %v\n", cap(nums))

  nums2 := [...]int{ 10, 11, 12 }
  fmt.Printf("nums2: %v\n", nums2)

  numsc1 := [...]int { 1, 2, 3 }
  numsc2 := numsc1 // creates a copy. So numsc2 and numsc1 do not point to the same location in memory

  fmt.Printf("numsc1: %v\n", numsc1)
  fmt.Printf("numsc2: %v\n", numsc2)

  fmt.Printf("numsc1[0] = 10\n")
  numsc1[0] = 10

  fmt.Printf("numsc1: %v\n", numsc1)
  fmt.Printf("numsc2: %v\n", numsc2)

  numsr1 := [...]int { 1, 2, 3 }
  numsr2 := &numsr1 // numsr2 and numsr1 point to the same location in memory

  fmt.Printf("numsr1: %v\n", numsr1)
  fmt.Printf("numsr2: %v\n", numsr2)

  fmt.Printf("numsr1[0] = 10\n")
  numsr1[0] = 10

  fmt.Printf("numsr1: %v\n", numsr1)
  fmt.Printf("numsr2: %v\n", numsr2)

  // slices
  // passed by reference
  // works for arrays or other slices

  fmt.Println("\n\nSlices\n")

  var nums3 []int = []int { 100, 200, 300, 400 }
  fmt.Printf("nums3: %v\n", nums3)
  
  nums4 := nums3[1:] // works just like it does in python nums3[[inclusive]:[exclusive]]
  fmt.Println("nums4 = nums3[1:]")
  fmt.Printf("nums4: %v\n", nums4)

  var nums5 = make([]int, 5, 5) // make([]type, len, [cap])
  fmt.Println("nums5 = make([]int, 5, 5)")
  fmt.Printf("nums5: %v\n", nums5)

  var nums6 = []int{ 101, 102, 103 }
  fmt.Printf("nums6: %v\n", nums6)
  fmt.Printf("nums6 + nums5: %v\n", append(nums6, nums5...))

  // emulating a stack
  // shift (add an element to the beginning)
  var nums7 = nums6[1:]
  fmt.Println("nums7 = nums6[1:]")
  fmt.Printf("nums7: %v\n", nums7)

  // removing an element from the middle
  fmt.Println("nums8 = remove(nums6, index = 1)")
  var nums8 = append(nums6[:1], nums7[1:]...)
  fmt.Printf("nums8: %v\n", nums8)

  // suppose we have a slice a = { 1, 2, 3, 4, 5 }
  // and we remove an element from the midle of a slice (even if its in another variable)
  // b = append(a[:1], a[1:]...)
  // when you try to access a, you'll notice the following:
  // a = { 1, 3, 4, 5, 5 }
  // if you look closely, you'll notice that the array a has had its element at index 1
  // replaced with the (former) elements ar indices 2...4 in order.
  // PS: the element at index 4 remains untouched
  // think of this like an overflow occurring at index 1
  
  
  // maps
  // maps are passed by reference

  fmt.Println("\n\nMaps\n")

  personDetails := make(map[string]int)
  // or
  // var personDetails = map[string]int
  fmt.Printf("personDetails: %v\n", personDetails)

  personDetails["age"] = 18
  personDetails["birthday"] = 19
  personDetails["unknown"] = 100

  fmt.Println("personDetails['age'] = 18")
  fmt.Println("personDetails['birthday'] = 19")
  fmt.Println("personDetails['unknown'] = 100")
  fmt.Printf("personDetails: %v\n", personDetails)

  fmt.Println("delete(personDetails, unknown)")
  delete(personDetails, "unknown")
  fmt.Printf("personDetails: %v\n", personDetails)

  fmt.Printf("personDetails['unknown']: %v\n", personDetails["unknown"])
  // but you'll notice that the 'deleted' key -- 'unknown' -- still has a value of 0
  // here's what we can do to check if the key actiually exists
  val, ok := personDetails["unknown"] // in the place of 'val' you can also use the write only variable -- '_'
  fmt.Printf("val, ok = %v, %v\n", val, ok)
  // so we can use the 'ok' variable to check if the key exists in the map
  fmt.Printf("len(personDetails): %v\n", len(personDetails))


  // struct
  // structs are passed by value

  fmt.Println("\n\nStructs\n")


  type Person struct {
    firstName string
    lastName string
    age uint8
    children []Person
  }

  fmt.Println("creating struct uchenna...")
  var uchenna = Person {
    firstName: "Uchenna",
    lastName: "Ofoma",
    age: 18,
  }

  fmt.Println("uchenna:", uchenna)
  fmt.Printf("uchenna.firstName: %v\n", uchenna.firstName)

  // anonymous structs
  fmt.Println("creating anonymous struct randomPerson...")
  var randomPerson = struct { name string } { name: "Stranger" }
  fmt.Println("randomPerson:", randomPerson)

  // embedding
  // Go doesn't use an OO model but rather it uses a Composition model

  type Animal struct {
    name string
    age int
  }

  // NOTE: that Bird has Animal like characteristics but it is not an Animal
  type Bird struct {
    Animal // notice that we did't do 'Animal animal'. This is because we wan't to embed the Animal struct into the bird (and not have an animal property on the Bird struct)
    canFly bool
  }

  fmt.Println("creating a penguin...")
  penguin := Bird { canFly: true } // super penguin ;D
  // or penguin := Bird { Animal: Animal { name: "penguin", age: 2 }, canFly: true }
  // for promoted fields, they have to be set this way. Oterwise you'd run into an error
  // saying cannot use promoted field Animal.[field] in struct literal of type Bird
  // where [field] is the name of the field you are trying to set
  penguin.name = "penguin"
  penguin.age = 2

  fmt.Println("penguin:", penguin)


  // tags in structs
  // tags are used to describe some specific info on the struct fields
  type Book struct {
    title string `required max:"255"` // the tag is defined inside the `` (backticks)
  }

  fmt.Println("creating a book...")
  // to get the tag you have to use the reflection API in the 'reflect' package
  t := reflect.TypeOf(Book {})
  field, fieldExists := t.FieldByName("title")
  fmt.Println("title.tag:", field.Tag) // prints out the tags of field 'title'
  fmt.Println("does the 'title' field exists on book?", fieldExists)

  // Conditionals
  fmt.Println("\n\nConditionals\n")
  
  // if statements
  
  if true {
    fmt.Println("the test is true")
  }

  // basic conditional operators apply in go as well
  // == != > >= <= <
  // basic logical operators also apply in go
  // || && !
  if unknown, ok := personDetails["unknown"]; ok { // variables used during the definition of a conditional block are scoped to that block
    fmt.Println("The 'unknown' field exists in personDetails")
    fmt.Println(unknown)
  } else { // has to be defined like this and not like this
  // }
  // else {
    fmt.Println("The 'unknown' field does not exist on personDetails")
  }

  // switch statemenets
  // unlike in other languages, in go the 'break' is implied when dealing with switch statments

  // you can't have overlapping cases in a switch case
  switch 2 { // here the '2' is called a 'tag'
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3, 4, 6:
      fmt.Println("three, four or six")
    default:
      fmt.Println("not one not two")
  }

  i := 2 + 3;
  // tagless switch
  switch {
    case i <= 7:
      fmt.Println("i is within the range [ -infinity, 7 ]")
      fallthrough // use thsi keyword to make the case statement execute the following block
      // fallthrough is logicless, so it does not obey the conition set in the next case block
    case i <= 10: // can overlap when using tagless syntax
      fmt.Println("i is within the range [ -infinity, 10 ]")
    default:
      fmt.Println("i is not within the required range")
  }

  // type switch
  var j interface {} = 1
  switch j.(type) {
    case int:
      fmt.Println("j is an integer")
    case string:
      fmt.Println("j is a string")
    default:
      fmt.Println("j's type is unknown")
  }


  // Loops
  fmt.Println("\n\nLoops\n")
  // there is only one loop statement in go and that is the for loop 
  // see fizzbuzz for more examples
  
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  for i,j := 0, 0; i < 10; i, j = i + 1, j + 2 {
    fmt.Println(i, j)
  }
}
