package main
import (
	"fmt"
	"sync"
	"time"
	"unsafe"
	"context"
)

func add(a int,b int) int{
	return a+b
}

func addmul(a int , b int) (int , int){
	return a+b, a*b;
}

func divide(a int, b int) (int , error){
	if b!=0{
		return a/b,nil
	}
	return 0, fmt.Errorf("Division by zero")
}


func outer() func(int32) int32{
	var count int32 = 0

	return func(a int32) int32{
		count++
		return count+a
	}
}

func somefunc(n int){
	fmt.Println("the number was : ",n)
}


type country struct{
	capital string
	population int64
}
func (c country) cap() string{
		return c.capital
}
func(c country) inc(n int64) {
	c.population += n;
}



type Person struct{
	name string
	address string
	number string
}



type T struct{
	a int32
	b bool
	c int32
}



func dowork(done <-chan int){
	for{
		select{
		case <-done:
			return
		default:
			fmt.Print("Working")
		}
	}
}



func stage1(data []int) <-chan int{

	out := make(chan int)

	go func(){
		for i:= range data{
			out<- i*i
		}
		close(out)
	}()
	return out
}

func stage2(in <-chan int) <-chan int{
	out:= make(chan int)

	go func(){
		for i:= range in{
			out<- i+5
		}
		close(out)
	}()

	return out
}




func main(){

	fmt.Println("Running")

	var string1 = "writer"
	var strup = []byte(string1)
	strup[2] = 'I'
	fmt.Println(string1, "\n",string(strup));

	// strup[3] = '₹'   this would fail as '₹' is a rune Its value is 8377 That overflows byte

	for i:=range 6{
		switch(i%2){
		case 1:
			fmt.Println("odd")
		case 0:
			fmt.Println("even")
		default:
		    fmt.Println("done")
		}

	}

	var arr = []string{"banana", "apple", "papaya"}

    s:= ""
	for _,fruit := range arr{
	  	s= fmt.Sprint(s," ",fruit)
	  	s= fmt.Sprint(s,' ',fruit) // this '' is a rune(just a placeholder for int32) for its UTF-8 value
		// '\n' newline rune (10)
		// "\n" newline string (length 1)
	}

	fmt.Println(s)
	k, l := addmul(5,4)
	fmt.Println(k,l)


	fun := outer()

	fmt.Print(fun(1))
	fmt.Print(fun(1))

	jod:= func(a int , b int) { fmt.Println(a+b)}
	jod(1,2)


    defer fmt.Println("first")
	defer fmt.Println("Second")
	go somefunc(5)
	go somefunc(7)

	mychan := make(chan string)


	go func(){
		mychan <-"string"
	}()

	msg:= <- mychan
	fmt.Println(msg)


	time.Sleep(time.Second*1)

	slice1 := []int{1,2,3}

	arr1:= [5]int{1,2,3,4,5}
	slice2 := arr1[1:4]

	fmt.Print(slice1, slice2)

	source:= []int{1,2,3}
	dest := make([]int, len(source))

	copy(dest, source) // else we would have a copy value

	source = append(source, dest...)

	row,column := 4,5
	matrix:= make([][]int,row)

	for i:= range matrix{
		matrix[i] = make([]int,column)
	}


	// Note: rows can have different lengths (jagged array)
	jagged := [][]int{
		    {1, 2},
		    {3, 4, 5, 6},
		    {7},
	}
	fmt.Println(jagged)



	pehlachan := make(chan int)
	dusrachan := make(chan int)


	go func(){
		pehlachan<- 7
	}()

	go func(){
		dusrachan<- 7
	}()

	select{
	case msg:= <-pehlachan:
        		fmt.Println("Pehlachan", msg);
	case msg:= <-dusrachan:
				fmt.Println("Dusrachan", msg)

	}

	 charchannel := make(chan int,3)

	go func(){
		for i:=range 10{
			charchannel<-i
			// fmt.Println(i)
		}
		close(charchannel)
	}()


	go func(){
		for{
		 <-charchannel
			// fmt.Println(i)
		}
	}()

time.Sleep(time.Second*1)


	var p Person
	p.name = "shankeleven"
	p.number = "09"

	fmt.Println(p.name)

	p1:= new(Person)
	p1.name="shankshank"

	fmt.Println(*p1)

	fmt.Println("\n\n\n")


    fmt.Println(unsafe.Sizeof(T{}))
	fmt.Println(unsafe.Alignof(T{}))
	fmt.Println(unsafe.Offsetof(T{}.a))
   // Group fields by decreasing alignment , this usually keeps the padding and cache-friendliness optimal
	//
	// go func(){
	// 	for{
	// 		select{
	// 		default:
	// 			fmt.Println("writing")
	// 		}
	// 	}
	// }()
	//

	// time.Sleep(time.Second*2)


	done:= make(chan int)

	go dowork(done)

	time.Sleep(time.Millisecond*5)
	done<-1

	fmt.Println("Pipeline")


	data:= []int{1,2,3,4,5,6}

	channel := stage1(data)
	final:= stage2(channel)

	for i:= range final{
		fmt.Println(i)
	}


 var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		defer wg.Done()
	}()

	wg.Wait()



	ctx, cancel := context.WithCancel(context.Background())


	go func(){

		for{
			select{
			case <- ctx.Done():
				return
			default:
			time.Sleep(time.Millisecond*2)
			fmt.Print("working")
		}}
	}()

	time.Sleep(time.Millisecond*20)

	cancel()


	var a any

	a=90

	val,sahi:= a.(int)
	if sahi{
       fmt.Println(val)
	}
	a = "ui"
	val,sahi= a.(int)
	if sahi{
       fmt.Println(val)
	} else{
      fmt.Println("Not a number")
	}



	// contexts






}
