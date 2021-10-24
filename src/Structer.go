package main

import "fmt"

type Video struct {
	 VideoMain
	 Author

}
type VideoBase struct{
	bsNum   int         //弹幕数量
	Viewer  int			//观看人数
	title   string      //题目

}
type VideoMain struct{
	VideoBase
	bsIsOpen bool		//是否开启弹幕
	inViewer  int       //正在观看人数
	publicTime int      //发布时间
	rank       int      //全站排行
	isPlaying bool      //是否正在播放
	speed     int		//播放速度
	definition int		//清晰度
	volume   int        //音量(1——100)
	isScreenExtend  bool//是否设置为宽屏
	isFallExtend    bool//是否设置为全屏
	isScreenSmall   bool//是否设置为小窗播放
	isWebFallExtend bool//是否设置为网页全屏
	bsOut     string    //发送弹幕
	likeNum   int       //点赞数
	coinNum   int       //硬币数
	collectNum int      //收藏数
	shareNum   int      //转发数

}
type Author struct {
	Name string             //名字
	SVIP bool               //是否是高贵的年度带会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}
func (v Video) coin(i int){ //投币
	if(i>2||i<=0){
		fmt.Print("最多只能投两个币哦")
		return
	}
	switch(i){
	case 1:v.coinNum++
	case 2:v.coinNum+=2
	}
}

func (v Video) like(){   //点赞
	v.likeNum++
}

func (v Video) collection(){ //收藏
	v.collectNum++
}

func (v Video) longPressThree(){  //一键三连
	v.like()
	v.collection()
	v.coin(1)
}

func outVideo(AuthorName string , videoName string)(v Video){ //发布视频
	v.Name=AuthorName
	v.title=videoName
	return v
}

/*func main(){
    var 毕导ICU Video
	毕导ICU.Name="毕导ICU"
	毕导ICU.title="xxxxxxx"

}
*/
