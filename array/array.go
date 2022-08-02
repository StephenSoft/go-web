package main
import(
	"fmt"
)
func add(array [5]string)(newArray [5]string){
	
	newArray = array
	newArray[4] = "aaaa"
	return
}

func suqare(len, width int) (suqare int){
	suqare = len * width
	return
}

func two_sum(nums []int, target int ) [2]int{
	m := make(map[int]int)
//   a := 
	for  index, value := range nums {

		// fmt.Println("Index:",index, "Value:", value)
		aaa_index, ok := m[target-value]; 
		// fmt.Println(aaa_index)
		fmt.Println(ok)

		if ok {
			return [2]int{aaa_index, index}
		}else{
			m[value] = index
		}
	}
	return [2]int{}
}


func main(){
	var ary = [5]string{}

	ary[0] = "lennon"
	ary[1] = "paul"
	ary[2] = "harison"
	ary[3] = "ringo"

	// ca := add(ary)

	// fmt.Println(ca)

	// aaa := suqare(2,5)
	// fmt.Println(aaa)
	// sss := []int{1,3,2,3,5,}
	// two_sum(sss, 5)

	// ccs := map[int]int{
	// 	1:1,
	// 	2:2,
	// }

	var a int = 5
	var b int = 10

	c := a - b
	d := b - a

	fmt.Println("c:", c)
	fmt.Println("d:", d)

}