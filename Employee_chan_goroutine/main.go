package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	empChan := make(chan Employee)

	for i := 0; i < 10; i++ {
		go func(id int) {
			empChan <- Employee{
				id:   id,
				name: fmt.Sprintf("Employee%d", id),
			}
		}(i)
	}
	employee := make([]Employee, 0, 10)
	for i := 0; i < 10; i++ {
		emp := <-empChan
		employee = append(employee, emp)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%s for id %d\n", employee[i].name, employee[i].id)
	}

}
