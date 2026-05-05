package main

import "fmt"



func dup2(s []int ) []int{
var res []int
	m:= make(map[int]bool)
	for i:= range s{
		_,ok:=m[s[i]]
		if !ok{
			continue
		}
		res = append(res,s[i])
	}
return res
}
/*
func dup1(s []int ) []int{

}
*/

type rect struct{
	len int
	wid int
}

func (r rect) area() int{
	return r.len*r.wid
}

func (r *rect) longen(n int) {
 r.len *= n
}

// use use pointer recievers
// when copying is expensive or ofc we want to update the struct


func main(){


	fmt.Println("started")

	var a [3]int
	a[2]=9;
	fmt.Println(a);

	nums:= [3]int{1,2,3}; // it's just that here compiler would enforce the size
	alt:= []int{1,2,3} // this is a slice
	alt1:= [...]int{1,2,3} // in these compiler would count the size
	fmt.Println(nums);
	fmt.Println(alt, len(alt1));

	fmt.Println(&nums[0]);
	fmt.Println(&nums); // in go, these are different than C , this is the address of the array [3]int and not nums[0]

/*
In Go, arrays and pointers to arrays are distinct types, and &nums is of type *[3]int32, not *int32
When you print &nums, fmt recognizes it as a pointer to an array and displays the array's contents, not the address.
If you want the address of the first element, you use &nums[0], which is of type *int32.
	*/

/*
but there are certain limitations that come with them , first is that the size shall be known at the compile time
and once declared the size cannot be changed.
Despite their rigidity, arrays have a few niche but important use cases in Go:
Fixed-size data like IP addresses
Low-level data structures
Interop with C or system calls
*/

/*
Because arrays are fixed-size, Go introduced slices: flexible, dynamic sequences built on top of arrays. Think of slices as views into arrays. A slice keeps three things:
Pointer: A reference to the underlying array.
Length: The number of elements in the slice.
Capacity: The maximum number of elements the slice can hold (which is always greater than or equal to the length).
	*/


	var s []int
	fmt.Println(cap(s));

	/*
With var s []int you are declaring a slice.
That means you’ve introduced a variable s of type “slice of int” ([]int), but you haven’t yet given it any backing array.
At this point, s is nil – it doesn’t point to any actual storage.
That’s why its length and capacity are both zero, until you allocate or append to it.
	*/

	var s1 = make([]int,3,5) // 3 is length and 5 is its capacity

// this is roughly equivalent to:

	/*
If there’s enough capacity, append just writes into the existing array. If not, Go automatically allocates a new larger array, copies the old elements over, and adds the new value. That’s why a slice can grow even though arrays themselves are fixed-size.
On one hand, this provides flexibility, but it can also lead to performance overhead due to the need for memory allocation and copying.
To mitigate this, it's a good practice to preallocate slices with an appropriate capacity when you know the size in advance.
	*/
	fmt.Println(len(s1),cap(s1));
	s1= append(s1, 2,3,4,5);
	fmt.Println(len(s1),cap(s1)); // len goes from 3 to 7 as first three were already set to zero values as make would create an underlying array for s1

	// so from what i undertand, slice is just a viewing window into and underlying array , storing the pointer to the array, so if we slice into a slice,
	// any changes whatsoever would affect both as the underlying array is same ,
	// this could come with certain performance trade-offs as memory reallocation could be required at times,
	// slices just come with an additional functionality of dynamic resizing .

	s2:= s1[2:4]
	s2[0]=9

	fmt.Println(s1,s2)
// when we pass a slice to a function, it is pass by value, we are passing the reference but the value of the slice struct is copied by value


	// so whenever i write [] in front of a datatype i define a slice and [2] defines an array
	// []int is an int slice. [2]int is an int array, [][]int is a slice of slices of type int, and [3][]int is an array of slices of type int

	// slice of slices mean each element is an independent slice and

	a2 := [2][3]int{{1,2,3},{3,4,5}}
	sa2 := a2[:]
	fmt.Println(a2,sa2)


	for i:=0;i<len(a2);i++{

	}

	// to be explored tomorrow

	fmt.Println("MAP_________")

	m:= make(map[string]int)
	m["shashank"]=90
	m["non-shasahnk"] = 100

	fmt.Println(m)

	/*
Go computes the hash of the key ("alice") to find which bucket in the hash table to look in.
A bucket is a small container within the hash table that holds one or more key-value pairs. When multiple keys hash to the same bucket, they are stored together inside it.
It searches the bucket for the key.
If the key exists, Go returns the associated value (23 in this case).
If the key doesn’t exist, Go returns the zero value of the map’s value type (0 for int, "" for string, nil for a pointer or slice, and so on).
	*/


	res:= m["re"]
	fmt.Println(res)
	// this returns zero value of the type
	// but the problem with this is to distinguish between values that do not exist or are zero
	//hence

	re, ok:= m["re"]
	if ok{
		fmt.Println(re)
	}

	for k,v := range m{
		fmt.Printf("%s: %d\n",k,v);
	}

	/*
Map iteration order in Go is randomized: each loop may produce keys in a different order.
This prevents you from relying on insertion order. If you need a deterministic order,
you can collect the keys into a slice, sort them, and iterate over the sorted keys.


Keys are hashed to decide which bucket they go into.
Each bucket holds multiple key-value pairs.
When a bucket gets too full, Go splits it into two (similar to dynamic resizing).
That's why map operations are usually O(1), but not guaranteed constant time.
Just keep in mind that maps are not safe for concurrent writes.
If multiple goroutines write to a map at the same time, you’ll get a runtime panic.
Use sync.Mutex or sync.RWMutex to protect map access in concurrent scenarios.
	*/






	type person struct{
		name string
		age int
	};
// three ways to initialise
	p1:= person{
		name:"shashank",
		age:40,
	}

	p2:= person{"shasha",47	} // ofc not recommended, as relies on order

	var p3 person
	p3.name="sha"
	p3.	age=89

	fmt.Println(p1,p2,p3)
	p4 := &p1
	p1.name= "shashanked"

	fmt.Println(p4,p2,p3)

	type address struct{
		city string
		lane string
	}

	type entry struct{
		id int
		person
		address
	}

	// we can directly access city as entry.city and this is a promoted field from the embedded struct,
	// including one struct in other is called composition



}
