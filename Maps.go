// a map is a data structure that provides n unordered key-value pairs where each key is unique.

/*
Real-life example :
Think of a phonebook:
Name → Phone number
"Rahul" → 9876543210
"Anita" → 9123456780

Here:
Key = Name
Value = Phone number
That’s exactly what a map does in Go.

map[keyType]valueType
map[string]int

*/

package main

import "fmt"

func main() {
    // Create a map
    marks := map[string]int{
        "Math":    90,
        "Science": 85,
        "English": 88,
    }

	// ANOTHER 
	students := make(map[string]int)
	students["Rahul"] = 100
	students["Anita"] = 200
	students["Suresh"] = 350

    // Print the map
    fmt.Println(marks)

	// Access a value using its key
	fmt.Println(marks["Math"])

	fmt.Println("Marks of Anita :" , students["Anita"]) // default value 0
	students["Anita"] = 400
	fmt.Println("Updated Marks of Anita :" , students["Anita"])


	// DELETE A KEY VALUE PAIR
	delete(students, "Suresh")
	fmt.Println("Students map after deleting Suresh:", students)

	//CHECK IF A KEY EXISTS
	marksValue, exists := marks["History"]
	if exists {
		fmt.Println("Marks in History:", marksValue)
	} else {
		fmt.Println("No entry found for History.")
	}



	// ITERATE OVER A MAP
	for subject, score := range marks {
		fmt.Printf("Subject: %s, Score: %d\n", subject, score)
	}
}

/* We use make function to create a map Because a map must be initialized before you can use it.
make() creates and initializes the map so Go can store data in it.
*/