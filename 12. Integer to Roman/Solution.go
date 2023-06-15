
func  intToRoman(num int) string {
    result := ""
    numThousands := num / 1000
    for i := 0; i < numThousands; i++ {
        result += "M"
    }

    num = num % 1000
    if num/100 == 9 {
        result += "CM"
        num -= 900
    }

    num5Hundreds := num / 500
    for i := 0; i < num5Hundreds; i++ {
        result += "D"
    }

    num = num % 500
    if num/100 == 4 {
        result += "CD"
        num -= 400
    }

    numHundreds := num / 100
    for i := 0; i < numHundreds; i++ {
        result += "C"
    }

    num = num % 100
    if num/10 == 9 {
        result += "XC"
        num -= 90
    }

    numFifties := num / 50
    for i := 0; i < numFifties; i++ {
        result += "L"
    }

    num = num % 50
    if num/10 == 4 {
        result += "XL"
        num -= 40
    }

    numTens := num / 10
    for i := 0; i < numTens; i++ {
        result += "X"
    }

    num = num % 10
    if num == 9 {
        result += "IX"
        num -= 9
    }

    numFives := num / 5
    for i := 0; i < numFives; i++ {
        result += "V"
    }

    num = num % 5
    if num == 4 {
        result += "IV"
        num -= 4
    }

    for i := 0; i < num; i++ {
        result += "I"
    }

    return result
}

