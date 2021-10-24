package main

import "fmt"

func Receiver(v interface{})  {
	switch v.(type){
	case int :
		fmt.Println("这个事int类型")
	case string:
		fmt.Println("这个事string类型")
	case bool:
		fmt.Println("这个事bool类型")
	}
}
/*func main(){

}*/
