package main

import "fmt"



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


type getter interface{
    get() string
    relocate (string)
}
// At runtime, an interface value in Go is implemented as a two-word pair:
//Type Information Pointer: Points to information about the  dynamic type stored in the interface.
//Data Pointer: Points to the actual data of the dynamic value. what we would later study as |var a getter|







type address struct{
    City string
}

type user struct{
    name string
    address address
}
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

     //
     // ""
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

    // go interfaces : a behaviour specification
    // An interface in Go is a type that defines a set of method signatures
    // that a concrete type must implement to satisfy the interface.

    


}


// Language Quirks
// Capitalization Controls Visibility (Exported(name starts with Capital letter)  vs Unexported)
// no implicit type conversion even from uint4 to uint16
// multiple return values
// the compiler inserts semicolon but is sensitive to \n
