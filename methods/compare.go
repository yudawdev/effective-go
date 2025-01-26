package methods

import "fmt"

type User struct {
	Name string
	Age  int
}

// u1: {U1 10}
// u2: {U2 11}
// u1==u2 : false
// -----------------------------------------------------
// u2: {U2 11}
// u3: {U2 11}
// u2==u3 : true
// -----------------------------------------------------
// u3: {U2 11}
// u4: &{U2 11}
// &u3==u4 : false
func compare() {

	u1 := User{
		Name: "U1",
		Age:  10,
	}

	u2 := User{
		Name: "U2",
		Age:  11,
	}

	u3 := User{
		Name: "U2",
		Age:  11,
	}

	u4 := &User{
		Name: "U2",
		Age:  11,
	}

	fmt.Printf("u1: %v\nu2: %v\nu1==u2 : %v\n", u1, u2, u1 == u2)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("u2: %v\nu3: %v\nu2==u3 : %v\n", u2, u3, u2 == u3)
	fmt.Println("-----------------------------------------------------")
	//fmt.Printf("u3: %v\n, u4: %v\n, u3==u4 : %v\n", u3, u4, u3 == u4)  比较时，类型不同，编译出错
	fmt.Printf("u3: %v\nu4: %v\n&u3==u4 : %v\n", u3, u4, &u3 == u4)
	fmt.Printf("u3: %v\nu4: %v\nu3==*u4 : %v\n", u3, *u4, u3 == *u4)
}
