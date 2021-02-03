package generator

//Read/Write into
// 1. Log file for tracking the process
// 2. Question and answer file for the operator 
// 3. Reads the config file to get max number that can be used for generating random number 

import ("os"
"fmt"
"bufio"
"strconv"
"strings"
"log"
)

func createlogger() *log.Logger{ 
    filename :=  "logger.txt"
    fo,err := os.Create(filename)
    if err != nil {
        fmt.Println("Error in create logger file", err)
        return nil
    }
    logger := log.New(fo,"App: ",log.Lshortfile)
	return logger 
}

func Createfile(name string, logger *log.Logger) *os.File{
    
    filename :=  name + ".txt"
    fo,err := os.Create(filename) 
	if err != nil {
        logger.Println("Error in create file")
        return nil
	}
	return fo
}

func Writefile(fo *os.File, data string){
	fo.WriteString(data)
}

func getMaxnumber(opt string, logger *log.Logger)int{
    fi,err := os.Open("config.txt") 
    if err != nil{ 
        logger.Println("Error in file open", err)
    }
    scanner := bufio.NewScanner(fi)
    scanner.Split(bufio.ScanLines) 
    var datamapping []string
 	for scanner.Scan() {
		datamapping = append(datamapping, scanner.Text())
	}
 	fi.Close()
 	for _,dm := range datamapping{
         a := strings.Split(dm," ")
         if a[0] == opt{
             n,err := strconv.Atoi(a[1])
             if err != nil{ 
                logger.Println("Error in Number conversion", err)
            }
             return n
         }
     }
     return 0
}