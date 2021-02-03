package generator
//implements the addition operator 

import ("strconv"
)

type addition struct{
    operator string
} 

func newAddition() *addition{ 
     return &addition{operator : "+"}
}

func (add *addition)question(n1 int, n2 int)string { 
    var s string
    s = strconv.Itoa(n1) + add.operator + strconv.Itoa(n2)
    return s
} 


func (add *addition)answer(n1 int, n2 int) string{ 
    var s string
    s = strconv.Itoa(n1) + add.operator + strconv.Itoa(n2) + " = " + strconv.Itoa(n1+n2)
    return s 
}