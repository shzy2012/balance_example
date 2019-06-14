package balance

import(
	"testing"
)
func init(){

}

func Test_Random(t *testing.T){
	RandomBalance()
}

func Test_RoundRobin(t *testing.T){
	for i:=0;i<100;i++{
		RoundRobinBalance()
	}
}