package main

type Human struct{
	Name string
	Age  int
}

type HumanSlice []Human

func (hs HumanSlice) Len()int{
	return len(hs)
}

func (hs HumanSlice) Less(i ,j int)bool{
	return hs[i].Age>hs[j].Age
}

func (hs HumanSlice) Swap(i ,j int){
	temp:=hs[i]
	hs[i]=hs[j]
	hs[j]=temp
}
/*func main(){

var humen HumanSlice
for i:=0;i<114;i++{
	human :=Human{
		Name: fmt.Sprintf("仙贝%d号", rand.Intn(514)),
		Age:  rand.Intn(1919810),
	}
	humen=append(humen,human)
}
	sort.Sort(humen)

	for _ , i:= range humen{
		fmt.Println(i)
	}

}*/