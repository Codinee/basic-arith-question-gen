package generator

// Controls the entire flow of the package. 
// random numbers are genereated and question and answer files are created for the arithmetic operator specified
// add, sub, mul and div are the operator for which 2 files are created  
import (
"math/rand"
"os"
"flag"
"time"
)


func Runfirst(){
    var q,a,opt string

	rand.Seed(time.Now().UnixNano())
	opt = getOperator()
    o:=getNewOperator(opt)
    logger := createlogger()
    logger.Println("Process started")
    fo := Createfile("sheet-ques",logger)
    f1 := Createfile("sheet-ans",logger)   
    n := getMaxnumber(opt,logger) //From config file
    for i:=0;i<50;i+=2{
        n1 := rand.Intn(n)
        n2 := rand.Intn(n)
        q= o.question(n1,n2) + "\n"  
        Writefile(fo,q)
        a = o.answer(n1,n2) + "\n"
        Writefile(f1,a)
    }
}
//operator for which file is generated is obtained from stdin thru flags

func getOperator()string{
    fs := flag.NewFlagSet("Flags", flag.ExitOnError)
    var args []string 
    args = os.Args[1:]
    fs.Parse(args) 
    op := fs.Arg(0)
	return op
}

//for each operator option new object is created
func getNewOperator(opt string) operator{
	switch opt{
	case "mul":
		o := newMultiply()
		return o
	case "div" :
		o := newDivide()
        return o
    case "add":
        o:= newAddition()
        return o
    case "sub":
        o:=newSubtract()
        return o
	default:
		return nil
	}
}

