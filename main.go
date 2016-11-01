package main 

import (
	"fmt"
)

//Structs
type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

//Function declaration inside a strcut 
func (s *Saiyan) Super() {
 s.Power += 10000
}
func main() {
	// if len(os.Args) != 2 {
	// 	os.Exit(1)
	// }

	// fmt.Println("It is over", os.Args[1])

	// var power int
	// power = 9000
	//or
	// var power int = 9000
	//or
	// power := 9000
	// fmt.Println("It is over", power)

	// power := 1000
	// fmt.Printf("default power is %d\n", power)
 //  name, power := "Goku", 9000
 //  fmt.Printf("%s's power is over %d\n", name, power)

  //This Won't complie 
  // name, power := "Goku", 1000 
  // fmt.Printf("default power is %d\n", power)

	//Declaring structs
  // Goku := Saiyan{
  // 	Name: "Goku",
  // 	Power: 9000,
  // }
  // goku := Saiyan{"Goku", 9000}
  // fmt.Printf("default power is %d\n", Goku.Power)
  // fmt.Printf("default power is %d\n", goku.Power)
  // //Another ways to declare
 //  goku := Saiyan{}
	// // or
	// goku := Saiyan{Name: "Goku"}
	// goku.Power = 9000



	goku := &Saiyan{"Goku", 9000, nil} 
	goku.Super()
	fmt.Println(goku.Power)

	gohan := &Saiyan{
  Name: "Gohan",
  Power: 1000,
  Father: &Saiyan {
		Name: "Goku", 
		Power: 9001, 
		Father: nil,
		}, 
	}
	gohan.Power += 234
	fmt.Println(gohan)
	// goku := new(Saiyan) 
	// goku.name = "goku" 
	// goku.power = 9001
	// //vs
	// goku := &Saiyan {
	//   name: "goku",
	//   power: 9000,
	// }
}

// func Super(s *Saiyan) {
// 	s.Power += 10000
// }

