package generator
// implements the subtraction process 

import ("strconv"
)

type subtract struct{
    operator string
} 

func newSubtract() *subtract{ 
     return &subtract{operator : "-"}
}

func (sub *subtract)question(n1 int, n2 int)string { 
    var s string
    s = strconv.Itoa(n1) + sub.operator + strconv.Itoa(n2)
    return s
} 


func (sub *subtract)answer(n1 int, n2 int) string{ 
    var s string
    s = strconv.Itoa(n1) + sub.operator + strconv.Itoa(n2) + " = " + strconv.Itoa(n1-n2)
    return s 
}