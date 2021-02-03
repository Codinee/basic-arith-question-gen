package generator

//implements the division operator 

import ("strconv"
)

type divide struct{
    operator string
} 

func newDivide() *divide{ 
     return &divide{operator : "/"}
}

func (div *divide)question(n1 int, n2 int)string { 
    var s string
    s = strconv.Itoa(n1*n2) + div.operator + strconv.Itoa(n2)
    return s
} 


func (div *divide)answer(n1 int, n2 int) string{ 
    var s string
    s = strconv.Itoa(n1*n2) + div.operator + strconv.Itoa(n2) + " = " + strconv.Itoa(n1)
    return s 

}