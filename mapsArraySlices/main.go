package main
import (
	"fmt"
)

//A way to maps into a struct 
type Saiyan struct {
	Name string
	Friends map[string]*Saiyan
}
func main() {
	 // var scores [10]int 
	 // scores[0] = 339

	// scores := [4]int{345, 234, 123, 54}

	// for index, value := range scores {
	// 	fmt.Printf("Index %d Value %d\n", index, value)
	// }


	//Slices a lightweight data structure that wrpas an array

	// var scores = []int{3, 5, 4, 3}

	// scores := make([]int, 10)

	//CRASHES no allocation in previous indexes
	// scores := make([]int, 0, 10) 
	// scores[7] = 9033 
	// fmt.Println(scores)

	// scores := make([]int, 0, 10) 
	// scores = append(scores, 5) 
	// fmt.Println(scores) // prints [5]


	// scores := make([]int, 0, 10) 
	// scores = scores[0:8] 
	// scores[7] = 9033
	// fmt.Println(scores)


	// scores := make([]int, 0, 5) 
	// c := cap(scores) 
	// fmt.Println(c)
	// for i := 0; i < 25; i++ { 
	// 	scores = append(scores, i)
	// 	// if our capacity has changed,
	// 	// Go had to grow our array to accommodate the new data if cap(scores) != c {
 //      c = cap(scores)
 //      fmt.Println(c)
 //   }
   //We can see from the above example, that [X:] 
   //is shorthand for from X to the end while [:X] 
   //is shorthand for from the start to X.
   // Unlike other languages, Go doesn’t support 
   //negative values. If we want all of the values of 
   //a slice except the last, we do

    // scores := []int{1, 2, 3, 4, 5} 
    // scores = scores[:len(scores)-1]

		// scores := []int{1, 2, 3, 4, 5} 
		// scores = removeAtIndex(scores, 2) 
		// fmt.Println(scores)

		// scores := make([]int, 100)

		// for i := 0; i < 100; i++ {
		// 	scores[i] = int(rand.Int31n(1000))
  // 	}
  // 	sort.Ints(scores)
		// worst := make([]int, 5) 
		// copy(worst, scores[:3]) 
		// fmt.Println(worst)


	//Maps

	lookup := make(map[string]int) 
	lookup["goku"] = 9001
	power, exists:= lookup["vegeta"]
  // prints 0, false
  // 0 is the default value for an integer
  fmt.Println(power, exists)

  // returns 1
	//total := len(lookup)
	// has no return, can be called on a non-existing key
	delete(lookup, "goku")

	//mapsgo dynamically but you canset up a size
	//lookup := make(map[string]int, 100)
	goku := &Saiyan{
		Name: "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan {
		Name: "Krillin",
		Friends: make(map[string]*Saiyan),
	}
	fmt.Println(goku.Friends["krillin"])

	characters := map[string]int{
		"goku": 9000,
		"gohan": 4500,
	}

	for key, value := range characters {
		fmt.Printf("Power %s Name %d\n", key, value)
	}
}

 //So when is :X is 0 to X - 1
// func removeAtIndex(source []int, index int) []int {
// 	lastIndex := len(source) - 1
// 	//swap the last value and the value we want to remove 
// 	fmt.Printf("Before index %d. Before last index %d\n", source[index], source[lastIndex])
// 	source[index], source[lastIndex] = source[lastIndex], source[index] 
// 	fmt.Printf("After index %d. After last index %d\n", source[index], source[lastIndex])
// 	fmt.Println(source)
// 	return source[:lastIndex]
// }