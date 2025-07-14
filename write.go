package main

import "time"
import "fmt"
import "unsafe"





// functions in go are first class citizens i.e. they could be treated like any other value or variable
// passing as arguments , assigning to a variable or reassiging the values
// since it's a statically typed language, function variables have fixed types both parameters and return(obviously)
// there is no concept of void , just don't mention the return type

    func add(a int , b int) int{
        return a+b
    }

    func addmul ( a int , b int ) (int , int){
        return a+b, a*b
    }

    func recover() bool{
        return true
    }

    func divide(a int , b int) (int,error) {
        if b==0{
            return 0, fmt.Errorf("division by %d",0)
        }
        return a/b,nil
    }

    func altdivide(a int , b int) (int , string) { // same functionality without using fmt.Errorf()
        if b==0{
            return 0, "division by 0"

        }
        return a/b , ""
    }



// use case for this first class citizenship

//     var operation func(int, int) int
// if mode == "add" {
//     operation = func(a, b int) int { return a + b }
// } else {
//     operation = func(a, b int) int { return a * b }
// }



// closures , variables that are closed over are stored in heap not stack , hence they persist
// it is a normal function but uses variables from outside its scope
func outer() func(int) int {
    count:=5
    return func (a int) int{ // has to be unnamedd as there is no point in giving name , if a name is to be
        // given , use first class citizenship and store it in a variable
        count+=a
        return count
    }

}

type writer interface{
	write(byte) (string , error)

}

// type readwrite struct{
//     write() int // this throws some syntax error
// }
// error is because you cannot define a method inside a struct, it has to be defined outside with
// a receiver argument, which is the struct itself

type getter interface{
    get() string
    relocate (string)
}
// At runtime, an interface value in Go is implemented as a two-word pair:
//Type Information Pointer: Points to information about the  dynamic type stored in the interface.
//Data Pointer: Points to the actual data of the dynamic value. what we would later study as |var a getter|





// structs
// A struct is a composite data type that groups together variables (fields) under a single name.
// It allows you to create complex data types that can hold multiple values of different types.
// Structs are value types, meaning they are copied when assigned to a new variable or passed to a function.
// They are defined using the type keyword followed by the struct keyword and a set of fi
// Under the hood, structs are implemented as a contiguous block of memory,
// where each field is stored in the order they are defined

type alignment struct{
	a int8
	b int32
}

type address struct{
    City string
}

type user struct{
    name string
    address address
}

// Essential inbuilt structs that I'd revise

// to be completed





// Interfaces
// Suppose I want to implement a program that outsources reading and writing operations to a file or a network connection.
// I would define an interface that specifies the methods for reading and writing data.

type readwriter interface{
	 write([]byte) (string , error)
	 read() (string , error)
}

type readwrite struct{
	rdwr readwriter
}

// sample readwriter

type iohandler struct{}

func (io iohandler) write(data []byte) (string, error) {
	// Implementation for writing data
	return "Data written successfully", nil
}
func (io iohandler) read() (string, error) {
	// Implementation for reading data
	return "Data read successfully", nil
}
// now we can use this iohandler as a readwriter
// and pass it to the readwrite struct or any other function that expects a readwriter interface

// var ioHandler inter = iohandler{}
   // rw := readwriter{writer: ioHandler}

// You are making a new readwriter object and telling it, "Here’s your writer — use this ioHandler to do the
// reading and writing.

// Unlike many languages, Go uses implicit interface satisfaction. You don't declare that a type implements an interface
// - if it has the right methods, it automatically satisfies the interface.



//  EMPTY INTERFACES
// An empty interface is a special type in Go that can hold values of any type.
// It is defined as `interface{}` and is different from nil, which is a specific value representing the absence of a value.
// The empty interface is often used when you want to write functions or data structures that can accept any type of value.
// This take 0 bytes in memory, as it does not store any type information.



// methods
// a method is a function with special reviever argument
// The receiver appears in its own argument list between the func keyword and the method name.

// func ( u user) get() string{
//     return u.name
// }

func (u user) get() string{
    return u.name
}

func (u user) relocate( newadd string ) {
    u.address.City = newadd
}

func (u *user) vrelocate( newadd string ) {
    u.address.City = newadd
}

//

func printnum(a int){
	fmt.Println(a)
}



func main(){
    fmt.Println("Hello world");

    //
    // var rit getter


    //

    //variables
    var a int32;
    var b int;
    a=0 // compiler would start crying if these are not used
    b=9

    fmt.Println(a,b)

    // In Go, strings are immutable;
    // once created, their contents cannot be changed.
    //  Any modification results in the creation of a new string,
    //  which can be inefficient in scenarios involving frequent changes.
    // In contrast, []byte slices are mutable, allowing in-place modifications
    // without additional memory allocations.
    // Modifying a string involves creating a new copy,
    // which can be costly in terms of memory and processing time.\
    //  []byte allows for modifications without such overhead
    // it is also used for standard io operations in packages like net, io


    str1 := "hello"
    str1 = "Hello" // new string allocated

    fmt.Print(str1)

    str:= "hello"
    strb := []byte(str)
    strb[0] = 'H'
    fmt.Println(string(strb)) // same string but with modified H
    // good for string where frequent changes are occuring
    // In Go, []byte is a slice of bytes, where each byte is an
    // uint8 (an unsigned 8-bit integer, range: 0–255).
    // each representing a UTF-8 notation letter



    by := []byte{65, 66, 67}
fmt.Println(by)             // [65 66 67]
fmt.Println(string(by))     // "ABC"



    // variations of switch-case
    // switch (x) {
    // case value1:
    //}

    // switch{ case condition1:   }

    //basic loop , here for is used in all settings , it is the only looping construct
//     for i:=0;i<b;i++{
//         for j:=0;j<i;j++{
//         fmt.Print("*")
//     }
//     fmt.Println()
// }

    // as a while loop
    i:=0
    for i<3 {
        fmt.Println(i)
         i++;
    }

    // infinite loop but with break condition
    i=0
    for{
        i++;
        if i>10{
            break
        }}
    fmt.Println(i)

    // for range
    fruits := []string {"banana","apple","papaya"}
    for index, fruit := range(fruits) {
        fmt.Printf("Index: %d , Fruit: :%s\n", index, fruit)

    }


    // go read the functions now
    addition := add(3,4)
    k,l := addmul(3,4)
    fmt.Println(addition, k,l )

    jod := func(a int ,  b int) int { return a+b}
    fmt.Println(jod(1,2))



    resp:=outer()
    fmt.Println(resp(5))
    fmt.Println(resp(5))

    // defer executes when a function is about to end or return
    // either through return or panic
    defer fmt.Println("function has ended")
    fmt.Println("function still going on"    )
    // So defers are stacked, and executed in reverse order.

    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")

// Each function call has its own defer stack maintained by the Go runtime, hence they execute even in panic cases.
// When you call defer f(), Go runtime would act as follows:
// Captures the function f and its arguments (evaluated immediately).
// Pushes them onto the function's defer stack.
// At the end of the function's execution, the Go runtime:
// Pops each deferred call off the stack.
// Executes them in LIFO order(obviously).


    x:= 10
    defer fmt.Println(x)// would still print 10 at the end of the function , because the parameters and everything are copied like a snapshot
    x=20
    // Arguments Are Evaluated Immediately

// Go stores the function pointer and
// evaluated arguments in a data structure called a defer record.
// These are allocated on the heap or stack frame, depending on escape analysis.

if r := recover(); r { // inline initialisation in if statement , here r is boolean
    fmt.Println("Recovered:", r)}

// try-catch does not exist because of philosophical differences
// go believes in Errors are values, explicit error handling using return values
// but we would see similar behaviour with recover() for unexpected errors

// Errorf() Used for recoverable errors in normal program flow.
// Panic is unexpexted and everything layer by layer in the call stack stops and executes their respective defer statements
// if we want everything else to work normally even after panic , we need to use recover()

result , err := divide(4,0)
if err!=nil{
    fmt.Println("Error:", err)
}   else { // \n  else would give error , because of some stupid syntax quirk
    fmt.Println("Result: " , result)
}

res, dikkat := altdivide(4,2)
if dikkat != ""{
    fmt.Println("DIkkat: ", dikkat)
} else{
    fmt.Println(res)
}




// pointers
// pointer arithmetic is not allowed , as it would confuse out stupid garbage collector
// only use this for pass by value and pass by reference
// pointer arithmetic could be used using hte package "unsafe" , but the name speaks for itself

m,n:= 34, 45
fmt.Println(m,n)
madd , nadd := &m , &n
fmt.Println(madd, nadd)

*madd = 54
fmt.Println(m,n)


// reference types in go : slices, pointers , function, channel and map
// maps are reference type because they store address of the hashtable that is located in the heap memory
// m2:=m1 and thereafter when changes are made to the either , both would be affected , hence reference type

// structs
//




// methods
// a method is a function with special reviever argument
// The receiver appears in its own argument list between the func keyword and the method name.

// func ( u user) get() string{
//     return u.name
// }

st:= user{
    "india" ,
     address{"dehradun"}}

     fmt.Println(st)


    // go interfaces : a behaviour specification
    // An interface in Go is a type that defines a set of method signatures
    // that a concrete type must implement to satisfy the interface.
	// think of an interface as a plug socket "interface" similar to India vs Dubai
	// assume there is a code that makes use of a database , now it does not need to know how the database optimises something
	// or how it stores data , it just needs to support insertion with "INSERT INTO" and retrieval with "SELECT FROM"
	// so the code would be written in such a way that it would use the interface and not the concrete type
	// this allows for flexibility and extensibility in the codebase, as new types can be added without modifying existing code.
	// so it is  is a contract that defines what methods a type must have.


     var rit getter = st // variable of an interface , can store any variable satisfying the interface
    fmt.Println("rit :", rit)


fmt.Println(st.get())
st.relocate("delhi") // this would not change
st.vrelocate("delhi") // this would change
fmt.Println(st)



// maps

grades:= map[string]int{ // best for compile-time known data , no relocation hence performant
    "india" : 98,
    "pakistan" : 2,
}

marks:= make(map[string]int) // worst as highest chances of reallocation as well as rehashing
marks["india"] = 98
marks["pakistan"] = 2

numbers:= make(map[string]int , 2) // best for dynamic as the hint helps approximate
numbers["india"] = 98
numbers["pakistan"] = 2
fmt.Println(grades)
fmt.Println(marks)
fmt.Println(numbers)

shamarks , exist := marks["pakistan"]

if exist {
    fmt.Println(shamarks)
}else{
    fmt.Println("not found , probably deleted by india")
}

// if they continue the war
delete(marks , "pakistan") // used with many structures just like make()

shamarks , exist = marks["pakistan"]

if exist {
    fmt.Println(shamarks)
    }else{
        fmt.Println("not found , probably deleted by india")
    }

    markscpy := marks // both would point to the same underlying map and changes to one would affect both
    fmt.Println(markscpy)

	// empty structs {}

    var xem struct {} // defined it directly as there were no items inside so, could do it otherwise
	xem = struct{}{} // empty struct , no fields , size 0 bytes
	fmt.Println("the size of empty struct is ", unsafe.Sizeof(xem)) // 0 bytes, used for signalling or as a placeholder
	// it is in unsafe because it violates the go's abstractions that keep you gay and hides low level details

	// go optimises empty structs to save memory
	opt_sample := []struct{}{{}, {}, {}} // slice of empty structs
	// they take 0 space in memory
	fmt.Println("the size of slice of empty struct is ", unsafe.Sizeof(opt_sample)) // 24 bytes in slice header ; 0 bytes for the values, as they are optimised to take no space


	// it can be used as methods recievers with no state
	/* type MyStruct struct{}

	func (MyStruct) DoSomething() {
    	fmt.Println("Doing something")
	}
*/


	done := make(chan struct{})
	go func() { //  no name becuase wanted to use inside main function
    // do something
    done <- struct{}{} // signal completion
	}()
	<-done



	// there is this clever implementation of sets using empty structs as the language go somehow does not come with it

	mockset := map[string]struct{}{
		"apple" : {},
		"banana" : {},
	}

	_, ok := mockset["appple"]
	if !ok{
	fmt.Println("ok, it does exist")
	}




// in interesting detail about padding and alignment
/* In memory, most CPUs require that data is aligned on certain byte boundaries. For example:
int32 (4 bytes) should start at addresses divisible by 4 & int64 (8 bytes) at addresses divisible by 8
This means Go might insert padding bytes between fields to satisfy these constraints.

*/

var padding_sample alignment;
	fmt.Println("the size of the alignment struct is : ", unsafe.Sizeof(padding_sample)) // this comes out to be 8
	// as there are 3 padding bytes after int8 , asit takes one byte




printnum(0) // this would run synchrounously , this means the main function call will block until it completes before moving to the next line of code.

go printnum(1) // this would run asynchronously , this means the main function call will not block and will move to the next line of code immediately
// this just forks off the function call to a new goroutine and the main function will continue executing without caring about it.
// this also means that the main function will not wait for the goroutine to finish before exiting.

	for i :=0;i<10;i++{
		go printnum(i)
		// this will create 10 goroutines that will print the numbers from 0 to 9 not necessarily in order
	}
	// some goroutines might finish before others, and the main function will not wait for them to finish before exiting.
	// so you would see out of 10 , some numbers are not printed sometimes, because the main function exits before the goroutines finish executing.
	// comment out the time.Sleep() line to see this behaviour

	time.Sleep(2 * time.Second) // now all would be printed



}


// Language Quirks:
// Capitalization Controls Visibility (Exported(name starts with Capital letter)  vs Unexported)
// Names starting with a capital letter are exported, meaning they can be accessed from other packages.
// Names starting with a lowercase letter are unexported, meaning they are only accessible within the same package.
// no implicit type conversion even from uint4 to uint16
// multiple return values
// the compiler inserts semicolon but is sensitive to \n






/* what values can be nil and what cannot be nil:

Every function call in Go passes arguments by value. That means:
The function receives a copy of the argument.
But if the argument is a reference type, the copy still points to the same underlying data.


┌────────────────────────┐
│     Reference Types    │
│ (Can be nil, 0 bytes)  │
├────────────────────────┤
│ *T      => Pointer     │
│ interface{} / error    │
│ []T     => Slice       │
│ map[K]V => Map         │
│ chan T  => Channel     │
│ func()  => Function    │
└────────────────────────┘

┌────────────────────────┐
│      Value Types       │
│ (Never nil, have data) │
├────────────────────────┤
│ int, float, bool       │
│ string                 │
│ struct{}               │
│ [N]T => Array          │
└────────────────────────┘

The term "reference type" in Go refers to types that internally hold a pointer to data.

_________________________________________________________________________________
| Reference Type | What’s copied?                | Can modify underlying data? 	|
| -------------- | ----------------------------- | -----------------------------|
| `*T`           | The pointer (address)         | ✅ Yes                       |
| `[]T`          | Slice header (ptr, len, cap)  | ✅ Yes                       |
| `map[K]V`      | Map header (internal pointer) | ✅ Yes                       |
| `chan T`       | Channel handle                | ✅ Yes                       |
| `func`         | Function pointer              | ✅ Yes                       |
| `interface{}`  | Interface header              | ✅ Yes (depends on content)  |
|_______________________________________________________________________________|
*/


