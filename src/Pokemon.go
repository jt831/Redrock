package main

import (
	"fmt"
	_ "math"
	"math/rand"
)

type Action interface {
	Hit(p1 *Traveller,p2 *Traveller)  //普通攻击，使对方hp-1，不需要消耗魔力
	SHit(p1 *Traveller,p2 *Traveller) //强化攻击，使对方hp-10，但需要消耗2点魔力
	Miss()
}

//Traveller 人物类
type Traveller struct {
	Name string
	ATK  int //攻击力
	Hp   int
	Mp   int

}

func (t *Traveller)Hit(p1 *Traveller,p2 *Traveller){
	i:=rand.Intn(4)
	if i!=2{
		fmt.Printf("%s使用了普通攻击！%s受到了%d点伤害\n",p2.Name,p1.Name,p2.ATK)
		p1.Hp-=p2.ATK
	}else{
		fmt.Printf("%s并没有打中\n",p2.Name)
	}
}

func (t *Traveller)SHit(p1 *Traveller,p2 *Traveller){
	i:=rand.Intn(4)
	if i==2{
		fmt.Printf("%s使用了强化攻击！%s受到了大量伤害\n",p2.Name,p1.Name)
		p1.Hp-=p2.ATK*3
		p2.Mp-=2
	}else {
		fmt.Printf("残念desu~~ %s并没有打中\n",p2.Name)
		p2.Mp-=2
	}
}

func (t Traveller)Miss(){}



func status (p1 Traveller){//显示状态函数
	fmt.Printf(" %s\n 攻击力:%d\n Hp:%d\n Mp:%d\n ",p1.Name,p1.ATK,p1.Hp,p1.Mp)
}

func die (p1 Traveller)(bool){ //判断die没die 函数
	if p1.Hp<0{
		return true
	}
	return false
}

func battle (p1 Traveller,p2 Traveller){//塔塔开函数
	fmt.Printf("野生的%s出现了！你选择：\n1,普通攻击，造成等同于攻击力的伤害，不需要消耗魔力\n2,强化攻击，造成3倍攻击力的伤害，但需要消耗2点魔力\n3,尼给路达哟∑三三三⊂ (っゝд・)っ\n",p2.Name)
	times:=0
	for ;p1.Hp>=0&&p2.Hp>=0;{
		var chose ,chose2 int

		var a1 Action=&p1

		fmt.Scanf("%d",&chose)
		judge:=1
		switch chose{
		case 1:a1.Hit(&p2,&p1)
		case 2:
			if p1.Mp>=2 {
				a1.SHit(&p2,&p1)
			}else{
				fmt.Printf("%s的蓝量不足\n",p1.Name)
			}
		case 3:judge=rand.Intn(3)
			if judge==2{
				break
			}else{
				fmt.Println("逃跑失败，回来吧你")
			}
		}
		if judge==2{
			break
		}

		fmt.Printf("现在是%s时间\n\n",p2.Name)
		chose2=rand.Intn(2)
		switch chose2{
		case 0:a1.Hit(&p1,&p2)
		case 1:
			if p2.Mp>=2 {
				a1.SHit(&p1,&p2)
			}else{
				fmt.Printf("%s的蓝量不足\n",p2.Name)
			}
		}
		times++
		fmt.Printf("第%d回合结束，双方的状态为\n",times)
		status(p1)
		status(p2)
		if p2.Hp<=10&&p2.Hp>=0{
			fmt.Printf("%s已经菠萝菠萝哒\n",p2.Name)
		}

	}
	if die(p1)||p2.Hp>0{
		fmt.Println("菜")
	}else if die(p2){
		fmt.Printf("成功干掉了%s打的不错",p2.Name)
	}

}
func main(){
	fmt.Printf("亲爱的旅行者，欢迎来到提瓦特！\n请在双子中任选一人，1为空，2为荧：\n")
	var sex int

	for ; ;{
		fmt.Scanf("%d",&sex)
		if sex==1{
			fmt.Printf("您选择了龙哥\n")

			break
			} else if sex==2{
				fmt.Printf("您选择了妹妹\n")

			break
			} else{
			fmt.Printf("只能选择哥哥或者妹妹哦，请重新选择\n")
		}
	}
	fmt.Println("请输入您的昵称")
	var name string
	fmt.Scanf("%s",&name)
	fmt.Printf("角色创建成功，%s的初始能力为\n",name)
	player:=Traveller{
		Name:name,
		ATK: 5,
		Hp: 30,
		Mp: 16,
	}
	shiLaiMu:=Traveller{
		Name: "史莱姆",
		ATK: 1,
		Hp:5,
	}
	container:=Traveller{
		Name:"天理的维系者",
		ATK: 10,
		Hp: 50,
		Mp:100,
	}


	status(player)
    fmt.Println()
	fmt.Println("你来到了一块陌生的土地，四周都是迷雾.\n<<*************************** (ﾟДﾟ≡ﾟдﾟ)!? *************************>>\n请选择：\n1:向左\n2:向右\n3从寒天之钉上自由落体")
	var choose3 int
	fmt.Scanf("%d",&choose3)
	switch choose3{
	case 1:
		fmt.Println()
		fmt.Println("Ծ‸Ծﾉ✧   /ᐠ｡ꞈ｡ᐟ\\")
		fmt.Println("你来到了璃月，遇见了一个听书的男人.他给了你一块奇怪的石头，你感到攻击力和血量增加了！")
		player.ATK+=10;player.Hp+=30
		status(player)
	case 2:
		fmt.Println()
		fmt.Println("(๑`н´๑)   /ᐠ｡ꞈ｡ᐟ\\")

		battle(player,shiLaiMu)
	case 3:
		player.Hp=-1
	}
	if die(player){
		fmt.Println("菜")
		return
	}
	fmt.Println()
	fmt.Printf("(ー`′ー)     (°ー°〃)\n")
	battle(player,container)

}


