package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
	"os/signal"
)

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
// fmt.Stringer: An interface that defines a method String() string, allowing types to define their own string representation.
type stringer interface{
	String() string
}


// fmt.Error: An interface that defines a method Error() string, allowing types to represent errors in a custom way.
// type error interface{ // commented out as it was colliding with the built in error interface in the functions used here
// 	error() string
// }


// io.Reader: An interface that defines a method Read(p []byte) (n int, err error), allowing types to implement custom reading behavior.
// io.Writer: An interface that defines a method Write(p []byte) (n int, err error), allowing types to implement custom writing behavior.
type reader interface{
	Read(p []byte) (int, error)
}
type Writer interface{
	Write(p []byte) (int, error)
}
// io.Closer: An interface that defines a method Close() error, allowing types to implement custom cleanup behavior.
type closer interface{
	Close() error
}

// io.Seeker: An interface that defines a method Seek(offset int64, whence int) (int64, error), allowing types to implement custom seeking behavior.
type seeker interface{ // used a lot in on-disk file handling
	Seek(offset int64, whence int) (int64, error)
}

// sort.Interface: An interface that defines methods for sorting collections, allowing types to implement custom sorting behavior.
// similar to how we would implement index sorting in a database with a complex index ( cost , category preference , units_sold)
type sorter interface{
	Len() int // to detect the length of the collection ovbiously
	Less(i, j int) bool // comparision method passed in <>
	Swap(i, j int) // to swap the elements at index i and j
}

// Interface Composition
// we can also define our own custom interfaces that satisfy these built in interfaces or other custom interfaces.
type myStringer interface{
	stringer
	error
	reader
	Writer
}

type person struct { // to satisfy these built in interfaces.
	Name string
	Age  int
}
func (p person) String() string{
	return fmt.Sprintf("Person output : %s and his age is %d", p.Name, p.Age) // this would be automatically called when fmt.Println() is used around this type
}

func (p person) error() string {
	return fmt.Sprintf("Person error : %s is not his name and his age is not %d", p.Name, p.Age) // this would be automatically called when fmt.Errorf() is used around this type
}

//



// fmt.Scanner: An interface that defines methods for scanning input, allowing types to implement custom input parsing.


// time.Time: Represents a point in time, with methods for formatting and parsing.


// net.Conn: An interface that represents a network connection, providing methods for reading and writing data over a network.

// other essentials :
/*
io.Reader - Reading data from any source
io.Writer - Writing data to any destination
io.Closer - Closing resources
fmt.Stringer - Custom string representation
error - Error handling
sort.Interface - Custom sorting
http.Handler - HTTP request handling
context.Context - Cancellation and deadlines
json.Marshaler/Unmarshaler - Custom JSON serialization
sql.Scanner/driver.Valuer - Database type conversion
sync.WaitGroup: A synchronization primitive that allows you to wait for a collection of goroutines to finish.
*/




// Interfaces
// Suppose I want to implement a program that outsources reading and writing operations to a file or a network connection.
// I would define an interface that specifies the methods for reading and writing data.
// do checkout design patterns for using interfaces in Go



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


// INTERFACE Good todos:
/*
Keep them compact and focused on a single responsibility(basically small).
Accept Interfaces, Return Structs
Avoid using interfaces for types that are not expected to change or be extended.
*/

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


	done := make(chan struct{}) // gd would not work as chan is not a type written in a .go file hence lsp cannot trace it
	// rather it is a keyword built into the languge so to refer to https://github.com/golang/go/blob/master/src/runtime/chan.go
	// Channels are implemented in Go’s runtime, which is written in a mix of Go and assembly. go to the link above to checkout runtime implementations of all other built-in types
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


shashank := person{ "Shashank", 25}
fmt.Println(shashank) // would call the String() method automatically


// concurrency

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

	// go uses a fork - join model for concurrency
	// but right now we are not using any synchronization mechanism to wait for the goroutines to finish before exiting the main function.
	// so there is no coordination between the goroutines and the main function. everything is running independently and concurrently.


	// channels
	// channels are used to communicate between goroutines and synchronize their execution.
	// sort of referring to the same place in memory to communicate , but with go it is fundamental to use channels for communication between goroutines
	// one goroutine can send data to a channel, and another goroutine can receive data from the same channel.basically a fifo queue
	// main function is also a goroutine, so it can also send and receive data from channels.
	// main funcction would block until the data is sent or received from the channel, so it is a blocking operation.

	// so if there is a channel that the main reads from and other goroutines write to , this could be a good way to achieve synchronization

	pehlachan := make(chan string) // create a channel of type string
	dusrachan := make(chan int)
    teesrachan := make(chan int)
	go func(){
		pehlachan <- "namaste ji" // send data to the channel

	}()
	sandesh := <- pehlachan // receiving data is a blocking operation, so the main function will wait until data is received from the channel
	// this acted as the join for the fork-join model
	fmt.Println("namaste yahase lekar")
	fmt.Println(sandesh)

	for i:=0;i<10;i++{ // to showcase for select loop
	go func(){
		dusrachan <- 69;
	}()

	go func(){
		teesrachan <- 6969;
	}()
	// just to notice the difference in timing as the signals are recieved
	time.Sleep(1*time.Second) // this is just to give some time for the goroutines to finish before the main function exits
	}
	// select statements

	select {
	case msg := <-dusrachan:
		fmt.Println("Received from dusrachan: ", msg)
	case msg := <-teesrachan:
		fmt.Println("Received from teesrachan: ", msg)
	}



	/* Vague internals:
The goroutine executing this select gets blocked if neither channel is ready.
if both channels are ready, it will choose one of them randomly to proceed with.
The Go scheduler parks this goroutine and resumes it only when any of the channels becomes
ready (i.e., someone sends a value on dusrachan or teesrachan).
It uses semacquire/semrelease operations in the runtime to handle this blocking efficiently (not a busy wait).

Each Go channel is implemented as a hchan struct internally in the runtime, which manages:
A queue of senders and queue of receivers
A buffer (if it's a buffered channel)
A mutex for synchronization
When you do <-dusrachan:
If the channel has data, it’s read and returned.
If it’s empty, your goroutine is queued as a receiver and goes to sleep.
The select statement is compiled into a call to runtime.selectgo().

Under the hood, it uses a randomized polling algorithm (to ensure fairness).
Here's what happens:
It checks the readiness of each channel.
If multiple are ready: one is picked at random.
If none are ready: goroutine blocks and is registered as a receiver on both channels.
When any channel receives data, the goroutine is woken up and the corresponding case is executed


The Go runtime attempts to avoid locking by using atomic instructions (e.g., CAS – compare and swap) for fast paths:
If a channel is ready immediately, data is pulled without locks.
Only the slow path (when blocking is needed) uses full synchronization.

Any references in msg are tracked by the garbage collector.
While the goroutine is parked, it’s still reachable and safe — Go's GC is goroutine-safe and concurrent.

The select gets translated to a jump table or switch-like structure in compiled code.
Branch prediction may play a role in performance if one channel is more frequently used.



	*/

// majorly used concurrency patterns in go

	// for select loop
	// Worker Pool: A fixed number of goroutines that process tasks from a shared channel.
	// Fan-out: Multiple goroutines reading from the same channel to distribute workload.
	// Fan-in: Merging multiple channels into a single channel to simplify communication.
	// Pipeline: A series of stages where each stage is a goroutine that processes data and passes it to the next stage.
	// Selective Receive: Using select to handle multiple channels and perform actions based on which channel is ready.

	// channels can either be buffered or unbuffered
	// Buffered channels allow sending and receiving without blocking until the buffer is full or empty.
	// an unbuffered channel provides a guarantee that the exchange takes place synchronously, exactly when the send and receive operations are executed.
	// Buffered channels are useful when you want to decouple the sender and receiver, allowing them to operate independently upto a certain limit.as islam puts it , send 2-3 years and forget , this 2-3 is the buffer capacity
	// a clever implementation : use this to limit the number of goroutines that can run concurrently by using a buffered channel as a semaphore.
	// for example, if you want to limit the number of goroutines to 3, you can create a buffered channel with a capacity of 3 and use it to signal when a goroutine is done.(using defer)
	// or you can use a buffered channel to limit the number of concurrent requests to a server by using it as a semaphore.
	/*

Internally, works like a FIFO circular queue.
Channels transfer copies of objects, rather than the actual objects.


So, when we talk about implementing channels, we can picture a simple way with a queue protected by locks.
Essentially, Go takes a similar approach, and this is embodied in the hchan structure. This structure plays a central role in channel implementation.
Go’s channel implementation is centered around three structures: hchan, waitq, and sudog


    */


	// bufchan := make(chan string, 3) // gd and read

	khelkhatam := make(chan struct{}) // used to signal that we have gracefully shut down
	// because if the main thread(go routine) exits we'd be done for good as this signal handling goroutine would stop as well
	sigchan := make(chan os.Signal, 1) // used to handle signals like ctrl+c(interrupt) or kill
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM) // notify the channel when an interrupt signal is received

	go func(){ // but this would in effect only after the main thread has gone past the signal handling code
		<-sigchan // this would block until a signal is received
		fmt.Println("\nReceived interrupt signal, exiting gracefully...")
		os.Exit(0) // exit the program gracefully
		close(khelkhatam) // close the channel to signal that we are done
	}() // this would run in a separate goroutine so that the main function can continue executing
/*
 if you handle the signal yourself (with signal.Notify) and do not exit immediately.
The Go runtime doesn’t automatically stop your app on SIGINT or SIGTERM.
It waits for your code to handle it — that's why your goroutine has time to run.
But, if you don’t handle the signal, or if your main goroutine exits early, the program ends immediately, and:
All goroutines stop.
Cleanup logic may be skipped.OS forcibly reclaims memory, file descriptors, etc.
*/





	go func(){

	for{  // this would mean that the main function would keep running and waiting for messages from the channels hence this is usually used in a server-like application
		select{
		case msg := <-dusrachan:
			fmt.Println("Received from dusrachan: ", msg)
		case msg := <-teesrachan:
			fmt.Println("Received from teesrachan: ", msg)
		 case <-time.After(500 * time.Millisecond): // this would block the select statement for 1 second
		// this acts as the heartbeat , if no message is received from either channel for 1 second, it would satisfy the case and execute the code
		// preventing the select statement from blocking indefinitely
		}
			// why does this print all at once rather than staggered? when the reciever is written later but staggered when the reciever is written earlier
			// quite an interesting problem , try figuring it out whenever you revisit.
			// could this be solved using buffered channels?
	}
 }() // for select loop is best suited for server-like applications where you want to keep the main function running and waiting for messages from the channels
	// keeping this in the main goroutine would block the main function and it would never exit




// Minimal definition for waitq and mutex to fix compile error , there's not any more sense to it
type waitq struct{}
type mutex struct{}

type hchan struct {  // the compiler is not crying because this is a type definition and not a variable declaration
	 qcount   uint           // Total data in the queue
	 dataqsiz uint           // Buffer size of the channel
	 buf      unsafe.Pointer // Pointer to an array of data elements
	 elemsize uint16         // Size of each element. Decide by type of elements
	 closed   uint32         // Flag indicating whether the channel is closed
	 elemtype any            // Type of the elements sent on the channel
	 sendx    uint           // Index of the next slot to send data
	 recvx    uint           // Index of the next slot to receive data
	 recvq    waitq          // Queue of waiting receivers
	 sendq    waitq          // Queue of waiting senders
	 lock     mutex          // Mutex for protecting the channel
}

// While making a channel, Go allocates hchan struct on a heap and returns a pointer to it. So, channel is just a pointer to a variable of type hchan.
/*
As explained above while we init a buffer channel it creates a buffer of channel length and waits for enqueue and deque the elements.
Let’s think there is a enque so it puts the elements in buffer and then deque poll the element from buffer.
and this is how vaguely the communication happens between goroutines.


When G1 (the sender) becomes scheduled for execution, there are two primary methods to resume the blocked receiver (G2):

Enqueue Approach: G1 enqueues the data into the channel’s buffer, dequeues the waiting G2 from the recvq, and signals the scheduler that G2 is ready to execute again.
Optimized Copying Approach: G1 directly copies the task object into the memory location reserved for G2’s stack, from the sudog.elem field.
Now question arises Why did G1 directly copy task0 into G2 stack instead enqueuing:

enque and deque are simple , you acquire the lock , make the read or write operation and thereafter release the lock.

The optimized copying approach is unconventional but efficient. Given that each goroutine has its separate stack space,
and goroutines do not access each other’s state directly, G1 can directly manipulate G2’s stack pointer.
This avoids the need for G2 to acquire a lock and modify the channel’s buffer. This optimization reduces memory copying overhead,
improving performance by reducing the overall synchronization overhead.



*/

tasks := []string{"task1", "task2", "task3"}
taskchan := make(chan string, 3) // buffered channel to hold tasks
	for _,task := range tasks{
		taskchan <- task

	}



	<-khelkhatam // this would block the main function until the channel is closed, which happens when the signal is received
	// right now ofcourse this is unreachable

}
// var a chan


// Language Quirks:
// Capitalization Controls Visibility (Exported(name starts with Capital letter)  vs Unexported)
// Only the Names starting with a capital letter are exported, meaning they can be accessed from other packages.
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
| `*T`           | The pointer (address)         | Yes 	                        |
| `[]T`          | Slice header (ptr, len, cap)  | Yes                          |
| `map[K]V`      | Map header (internal pointer) | Yes                          |
| `chan T`       | Channel handle                | Yes                          |
| `func`         | Function pointer              | Yes                          |
| `interface{}`  | Interface header              | Yes (depends on content)     |
|_______________________________________________________________________________|
*/



