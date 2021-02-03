package generator

// implements multiplication process 

import ("strconv"
)

type multiply struct{
    operator string
} 

func newMultiply() *multiply{ 
     return &multiply{operator : "*"}
}

func (mul *multiply)question(n1 int, n2 int)string { 
    var s string
    s = strconv.Itoa(n1) + mul.operator + strconv.Itoa(n2)
    return s
} 


func (mul *multiply)answer(n1 int, n2 int) string{ 
    var s string
    s = strconv.Itoa(n1) + mul.operator + strconv.Itoa(n2) + " = " + strconv.Itoa(n1*n2)
    return s 
}