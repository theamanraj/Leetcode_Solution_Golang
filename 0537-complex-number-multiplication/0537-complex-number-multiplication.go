type Complex struct{
    Real int
    Imag int
}

func ToComplex(num string) Complex {
    plusIndex := strings.Index(num, "+")
    imagIndex := strings.Index(num, "i")
    r, _ := strconv.Atoi(num[:plusIndex])
    c, _ := strconv.Atoi(num[plusIndex + 1 : imagIndex])
    return Complex {
        Real: r,
        Imag: c,
    }
}

func Mul(c1, c2 Complex) Complex {
    return Complex {
        Real: c1.Real * c2.Real - c1.Imag * c2.Imag, 
        Imag: c1.Real * c2.Imag + c2.Real * c1.Imag,
    }
} 

func complexNumberMultiply(num1 string, num2 string) string {
    result := Mul(ToComplex(num1), ToComplex(num2))
    return fmt.Sprintf("%v+%vi", result.Real, result.Imag)
}