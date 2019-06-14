package balance

import(
	"crypto/rand"
	"math/big"
	"fmt"
)
//Server 服务器
type Server struct{
	Host string
	Port int
	Value int 
}

//ServerList 服务器列表
var ServerList []Server

func init(){
	ServerList=append(ServerList,Server{"127.0.0.1",8081,1})
	ServerList=append(ServerList,Server{"127.0.0.1",8082,1})
	ServerList=append(ServerList,Server{"127.0.0.1",8083,1})
}

/****************************************************/
//随机算法
/****************************************************/

//RandomBalance  随机算法
func RandomBalance(){
	lens:= int64(len(ServerList))
	for i:=0;i<10;i++{
		randNum,_:= rand.Int(rand.Reader,big.NewInt(lens))
		fmt.Print(randNum)
	}
}


var curIndex int
//RoundRobinBalance 轮询算法
func RoundRobinBalance(){
	fmt.Printf("%v\n",ServerList[curIndex]) 
	curIndex++
	//curIndex = (curIndex + 1) % len(ServerList)
	if curIndex>len(ServerList)-1{
		curIndex=0
	}
}

