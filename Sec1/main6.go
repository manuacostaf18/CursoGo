package main
import "fmt"

// Control structures 2

func main(){
	fmt.Println("vim.go")

	// Continue
	/*for i := 0; i < 10; i ++ {
		if i%2 == 0{ // If it satisfies the condition, it continues with the loop above
			continue
		}
		fmt.Println(i)
	}*/

	// Break 
	/*flag := true
	for i := 1; i < 10; i ++ {
		if i%2 == 0{
			flag = false 
			break
		}else if i == 1{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(flag)*/

	// Switch 
	day := "Fri"
	/*switch day{
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("Default")
	}*/

	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}



}