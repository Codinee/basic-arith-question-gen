package generator

//Operator interface - is implemented by addition, subtraction, multiplication and division process 

type operator interface{
    question(int, int)string
    answer(int, int)string
}