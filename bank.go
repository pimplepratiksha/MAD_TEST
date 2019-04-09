package main
import (
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
)
func myfunc(i int,cust <-chan int,final chan<- int,timecust int){
	t:=((time.Now()).Format("2006-01-02 03:04:05"))
	for j :=range cust{
		fmt.Println(t,"--> cashier :",i,"customer ",j,"started")
		time.Sleep(time.Duration(timecust) * time.Second)
		fmt.Println(t,"--> cashier :",i,"customer ",j,"Completed")
		final <- 1
	}
}

func main(){
	//fmt.Println("hiiiiii",time.Now())
	t:=((time.Now()).Format("2006-01-02 03:04:05"))
	fmt.Println(t,"--> Bank Simulation started")
	 
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	
    nocashier1:=strings.Split(arg1,"=")
	nocust1:=strings.Split(arg2,"=")
	timecust1:=strings.Split(arg3,"=")
	
	nocashier,_:=strconv.Atoi(nocashier1[1])
	nocust,_:=strconv.Atoi(nocust1[1])
	timecust,_:=strconv.Atoi(timecust1[1])
	
	cust:=make(chan int,nocust)
	final:=make(chan int,nocust)

	for i:=0;i<nocashier;i++{
		go myfunc(i,cust,final,timecust)
	}

	for j:=1;j<=nocust;j++{
		cust <- j
	}
	close(cust)

	for k:=1;k<=nocust;k++{
		<- final
	}
	close(final)

	t=((time.Now()).Format("2006-01-02 03:04:05"))
 	fmt.Println(t,"--> Bank Simulation end")
}