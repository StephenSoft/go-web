package main

import(
	"fmt"
)

type Sleeper interface{
	sleep()
}

type User struct{
	Name string
	Age int
	Gender int
}

func (user *User) sleep() (sleepIno string){
	sleepIno = user.Name + "Sleepping Now"
	// fmt.Println(sleepIno)
	return
}

type Admin struct{
	User
	title string
}

func (admin *Admin) sleep()(sleepInfo string){
	sleepInfo = admin.title + ":" + admin.Name + "Sleepping Now"
	return
}


func testInterface(){
	user := User{"小明", 18, 1}
	info := user.sleep()

	admin := Admin{User: User{Name: "长贵", Age: 56, Gender: 1}, title: "象牙山村长" }

	adminInfo := admin.sleep()
	
	fmt.Println(info)
	fmt.Println(adminInfo)
}
func main(){

	sum := 0
	n := 100

	sum = (1 +n) * n / 2
	fmt.Println(sum)
}