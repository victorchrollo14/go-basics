// data types, conditionals, loops, type casting and inference
// arrays, slices, maps, make, structs

package main

import (
	"fmt"
	"math"
	"strings"
)

func var_declaration() {
	// variable declaration
	var b int       // uninitialized values use a default value here it is 0
	var c float32   // uninitialized values use a default value, using 0 here
	var bolion bool // default value is false
	b = 25          // variable initialization

	// declaring and initializing in the same line;
	var a = "string"

	// multi variable declarations
	var x, y int = 4, 5

	// declaring and initializing without var
	f := "new string"
	d := 34

	const DATABASE_URL = "mongosh://someurl:db"
	fmt.Println(a, b, c, x, y, f, d, bolion, DATABASE_URL)

}

func data_types_primitive() {
	// the main primitive datatypes are int, string, float, bool, byte and rune;
	// non primitive datatypes are array, slice, map, struct, pointer, channel, interface;

	// there are a lot of datatype variants like int8, int16, int32, etc. we can use int most of the time, unless we need
	// the other ones specifically for some reason.
	var a int = -23 // signed integers store both positive and negative values
	var b uint = 23 // unsigned integers store only positive values;

	c := 34.32 // defaults to float64
	var isOk bool = true

	// we can also represent complex numbers
	realNum := 34.0
	imaginary := 65.0
	complexNum := complex(realNum, imaginary)
	complexNum2 := 3 + 4i

	// it's an alias for uint8
	var rbyte byte = 'A'

	// this is to represent unicode characters
  // enclosed with single quotes all the time, double quotes is string
	var rPound rune = '£'
	fmt.Println(a, b, c, isOk, complexNum, complexNum2, rbyte, rPound)
	fmt.Printf("Unicode CodePoint: %U\n", rPound)

}

func formatting() {
	// formatting;
	var name string = "howl"
	var age int = 200

	fmt.Printf("the name is %v, and age is %v \n", name, age)
	fmt.Printf("the typeof(name) is %T, and typeof(age) is %T \n", name, age)
	fmt.Printf("the binary format of age %b \n", age)
	fmt.Printf("the base 16 format of age %X", age)
}

func arrays_and_slices() {
	// non primitive datatypes
	// arrays - contain fixed size, you cannot append or remove items after initialization;
	var fruits = [4]string{"apples", "bananas", "pineapple"}
	var vegetables = [...]string{"beans", "carrot", "beetroot", "tomato"} // inferring array length based on values provided
	var items = [5]int{}
	fmt.Println(fruits, "\n", vegetables, items)

	// slices - dynamic arrays, dynamically add, remove items from an array
	var devilFruits = []string{"gum-gum-fruit", "hito-hito-nomi", "magma-fruit"}
	fmt.Println("initial list", devilFruits)
	devilFruits = append(devilFruits, "chop-chop-fruit", "light-light-fruit", "sand-sand-fruit")
	fmt.Println("after appending an item", devilFruits)

	fmt.Println("the length of slice = ", len(devilFruits))
	fmt.Println("the capacity of slice = ", cap(devilFruits))

	// accessing elements in a slice;
	var first_fruit string = devilFruits[0]
	fmt.Println("the first devil fruit = ", first_fruit)

	devilFruits[1] = "fish-fish-fruit"
	fmt.Println("the modified slice = ", devilFruits)

	// iterating over a slice
	for i := 0; i < len(devilFruits); i++ {
		fmt.Printf("index = %v\nvalue = %v \n\n", i, devilFruits[i])
	}
	for i, val := range devilFruits {
		fmt.Println(val, i)
	}

	// comparing arrays
	if vegetables == fruits {
		fmt.Println("fruits and vegetables are the same")
	} else {
		fmt.Println("nope they are different")
	}

	// copying arrays, but array2 is not an instance of array1
	array1 := [3]int{3, 4, 3}
	var array2 [3]int = array1
	array2[0] = 39
	fmt.Println(array1, array2)

	// slicing an array
	var devilFruitsSlice = devilFruits[0:2]
	fmt.Println("the slice from array = ", devilFruitsSlice)

	var makeSlice = make([]int, 5, 10)
	makeSlice = []int{3, 4, 5}
	makeSlice = append(makeSlice, 6, 7, 4)
	fmt.Println("the slice created with make keyword = ", makeSlice)

	// append slices
	// appendedSlice := append(devilFruitsSlice, makeSlice)
	// fmt.Println(appendedSlice)
}

func iterables() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		// fmt.Printf("sum = %v in iteration = %v\n", sum, (i+1))
	}
	fmt.Printf("sum = %v", sum)

	// works like a while from other languages.
	for sum < 100 {
		sum += 50
	}
	fmt.Printf("the final sum = %v\n", sum)

	if x := "hello"; x != "hi!" {
		fmt.Println("true")
	}

	// this statement runs after the last statement is completed
	defer fmt.Println("the statement returned a value")

	fmt.Println("returning value")

}

func pointer_struct_slices() {
	// more types structs, pointers, maps etc

	// declaring a pointer variable which is a pointer to an integer variable, whose initial value is nil
	var p *int
	fmt.Println(p)

	i := 23
	// getting the variable memory address using &i and assiging it to the pointer p, now p is the
	// memory address of variable i
	p = &i
	q := &i

	// *p is the value at memory address p
	fmt.Println(i, *p, q)

	// structs
	type Person struct {
		name   string
		age    int
		gender string
	}

	victor := Person{"victor", 22, "male"}
	fmt.Println(victor)

	// we can also use pointers on struct and store the memory address of a pointer
	var sp *Person
	sp = &victor
	fmt.Printf("value stored at memory address of pointer sp = %v and memory address = %v \n", *sp, sp)

	// we can also access the struct values from the pointer variable
	fmt.Println("name of person = ", sp.name) // syntactic sugar, no need to use *sp.name

	type Car struct {
		name, model string
		price       int
	}

	maruti := Car{name: "maruti", model: "n32"} // we can declare only few values of a struct using named syntax
	tesla := Car{name: "tesla", model: "musk-melon12"}
	nothing := Car{}
	// a pointer variable can be declared using var and also using this syntax which is short hand
	short_hand_pointer := &Car{"some car", "m18", 120000} // has type *Car

	fmt.Println(maruti, tesla, nothing)
	fmt.Println(short_hand_pointer)

	// arrays, very limited. we can only declare them with specific length, read values, change values via index
	var nums [4]int
	nums[2] = 10
	nums[3] = 12

	fmt.Println(nums)
	fmt.Println("printing between a certain range =", nums[1:3])

	// slices, we are gonna use this for most array stuff
	names := []string{"victor", "matt", "michael", "manja", "seth-rollings"}
	fmt.Println(names)

	cars := []Car{
		{"maruti", "x23", 1233},
		{"tesla", "1hdd", 123023},
	}
	fmt.Println(cars, nothing)

	// extracting a range of values from a slice

	first_names := names[0:3]
	// fist_name_alt := names[:3] // same as above 0 - 3
	fmt.Println("first 3 names", first_names)

	last_names := names[3:]
	fmt.Println("last names = ", last_names)

	// length and capacity
	total_names := len(names)
	names_capacity := cap(names)
	fmt.Printf("total number of names = %v and total capacity = %v", total_names, names_capacity)

	// more stuff in slices
	s := []int{12, 12, 43, 2, 5, 3}

	fmt.Println("original slice = ", s)
	// slice the slice and make it 0 or empty
	s = s[:0]
	fmt.Println("emptied slice = ", s)

	// extend slice values
	s = s[:4]
	fmt.Println("extended slice = ", s)

	// drop first 2 values
	s = s[2:]
	fmt.Println("slice without first 2 elements = ", s)

	var bike []string

	// the type of uninitialized slice is nil
	if bike == nil {
		fmt.Println("bike is nil")
	}

	// using make to create slices
	schools := make([]string, 5)
	schools = append(schools, "auxilium", "citizen", "new baldwin")
	fmt.Println(schools)

	// slice of slices
	matrix := [][]string{
		[]string{"-_-", "-_-", "-_-"},
		[]string{"__", "__", "__"},
		[]string{"-_-", "-_-", "-_-"},
	}
	fmt.Println(matrix)

	// tick tack toe game
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// players take return
	board[0][1] = "X"
	board[0][2] = "O"
	board[2][2] = "X"
	board[1][1] = "X"
	board[0][0] = "O"

	for i = 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], "\t"))
	}

	// range in iteration works on maps and slice
	for i, v := range cars {
		fmt.Println(i, v)
	}
}

func maps() {
	// maps
	type Vector3 struct {
		x, y, z float64
	}

	var points map[string]Vector3
	fmt.Println(points)
	// can't pass value to a nil map, we need to use make to declare and initialize map in order to use it
	// points["one"] = Vector3{
	//   x: 2,
	//   y:-2.3,
	//   z: 0.23,
	// }

	// initialize map with make function, which makes it ready to use
	points3d := make(map[string]Vector3)
	points3d["one"] = Vector3{
		12.2,
		10,
		21,
	}
	fmt.Println(points3d["one"])

	// we can omit the type name since the definition has it
	newPoints := map[string]Vector3{
		"one": {x: 0.12},
		"two": {z: -12.2},
	}

	// can't omit here since go needs to know what's the type we are assigning here
	newPoints["three"] = Vector3{23, 0.12, 21}
	fmt.Println(newPoints)

	// mutation in maps
	// adding a new key value pair
	newPoints["four"] = Vector3{0.21, 0.23, 12}

	// getting a value
	point1 := newPoints["one"]
	fmt.Println("the first point in 3d space = ", point1)

	// deleteing a key
	delete(newPoints, "one")
	fmt.Println("the 3d points after deleting element 1", newPoints)

	// checking if an elem is present in a map
	point2, ok := newPoints["one"]
	fmt.Println(point2, "isPresent?", ok)

}

// it does not allow you to define a function within a function, you can assign a variable to a function though
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func func_and_closures() {
	// functions are like values too
	// you can declare a variable and pass a function as value

	hypo := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// fmt.Println(hypo(3, 4))
	fmt.Println(compute(hypo))
	fmt.Println(compute(math.Pow))

	// closures - functions which have access to variables outside their scope even after the outer function ends;
	adder := func() func(x int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos := adder()
	fmt.Println("sum = ", pos(1))
	// pos function has access to sum variable, which is not inside the scope of the function returned by adder func
	fmt.Println("sum = ", pos(3))

}

/*
*  Methods - these are likes classses in other languages. you can add functions to a type
*
 */

// methods - there are no classes in go, but we can define methods on type
type Vertex struct {
	x, y float64
}

// adding a method called abs() to type Vertex
func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// adding a second method to Vertex type called scale
func (v Vertex) Scale(s int) Vertex {
	return Vertex{v.x * float64(s), v.y * float64(s)}
}

// we can also pass a pointer reciever in cases where we need to write to the vertex rather than reading from it
func (v *Vertex) modify(x, y float64) {
	v.x, v.y = x, y
}

// methods can be defined on all other types as well
type FloatNum float64

func (f FloatNum) abs() float64 {
	if f < 0 {
		return float64(-f) // go doesn't auto infer custom types like FloatNum as float64 hence need explicit convertion
	}
	return float64(f)
}

func methods() {

	v := Vertex{3, 4}
	fmt.Println(v.abs())
	fmt.Println(v.Scale(2))
	v.modify(4, 8)
	fmt.Println("vertex after modification = ", v)

	f := FloatNum(-10.23)

	fmt.Println("float abs value = ", f.abs())

}

type Shape interface {
	perimeter() int
	area() int
	modify(l, b int)
}

type Rectangle struct {
	length  int
	breadth int
}

// this means type Rectangle implements interface Shape, but you don't need to specify that explicitly
func (r Rectangle) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) modify(l, b int) {
	if r == nil {
		fmt.Println("<nil> the rectangle value is nil")
		return
	}
	r.length = l
	r.breadth = b
}

func main() {
	fmt.Println("hello world")

	// interfaces - these are a set of method signatures it's like a custom type for a class / methods in this case i.e in go.
	var shape Shape
	rect := Rectangle{3, 4}

	// throws error if type Rectangle only implements one method signature i.e perimeter and not area
	shape = &rect // type Rectangle implements Shape
	fmt.Println("the value of shape = ", shape)
	fmt.Println("the area of shape = ", shape.area())
	fmt.Println("the perimeter of shape = ", shape.perimeter())

	// interface with nil underlying value
	var shape2 Shape
	var rect2 *Rectangle

	shape2 = rect2
	// even if rect2 is nil, shape2 is not nil which allows us to access the shape2.modify method,
	// then inside the modify function we can handle the nil error that prevents code from crashing
	fmt.Printf("shape type = %T and value = %v", shape2, shape2)
	fmt.Printf("rect2 type = %T and value = %v", rect2, rect2)
	shape2.modify(4, 7)

	// empty interface - holds value of any type
	var unknown interface{}
	unknown = rect
	fmt.Println(unknown)

	// type assertions
	var i interface{} = map[string]int{"x": 3, "y": 4}
	s, ok := i.(map[string]int) // asserts if interface i contains a value of type map of string key and int value 
	fmt.Println("containes type of map[string]int?", ok, s)

	f, ok := i.(float64) // returns false : panic if ok is not passed here
	fmt.Println("contains float value", ok, f)

	// type switches - it allows to assert multiple types for an interface at once;
	switch v := i.(type) {
	case int:
		fmt.Println("the interface i was of type int")
	case float64:
		fmt.Println("the interface i was of type float64")
	default:
		fmt.Printf("I don't know the type %T", v)
	}



}
