package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	// si alguno de los dos es 0 el resultado es automaticamente 0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var res [][]int
	var carry int
	// para evitar casos de numeros muy grandes vamos guardando los resultados de las multiplicaciones temporales en arrays
	for i := len(num1) - 1; i >= 0; i-- {
		temp_res := make([]int, len(num1)+len(num2))
		for j := len(num2) - 1; j >= 0; j-- {
			r := int(num1[i]-48)*int(num2[j]-48) + carry // 48 es el código ascii para 0
			carry = 0
			if r > 9 {
				carry = r / 10
				r = r % 10
			}
			temp_res[i+j+1] = r
		}
		if carry != 0 {
			temp_res[i] = carry
		}
		carry = 0
		res = append(res, temp_res)
	}
	//una vez tenemos todas las multiplicaciones temporales sumamos
	var final_res []rune
	for i := len(res[0]) - 1; i >= 0; i-- {
		r := carry
		carry = 0
		for j := len(res) - 1; j >= 0; j-- {
			r += res[j][i]
		}
		if r > 9 {
			carry = r / 10
			r = r % 10
		}
		final_res = append([]rune{rune(r + 48)}, final_res...)
	}

	// quitamos ceros sobrantes por la izquierda
	excess_ceroes := 0
	for i := 0; final_res[i] == 48; i++ {
		excess_ceroes++
	}
	final_res = final_res[excess_ceroes:]

	return string(final_res)
}

//otra opción más simplificada y rápida
func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]byte, len(num1), len(num2))
	for i := range res {
		res[i] = '0'
	}
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := (num1[i] - '0') * (num2[j] - '0')
			tmp_res := res[i+j+1] - '0' + product
			res[i+j+1] = tmp_res%10 + '0'
			res[i+j] = (res[i+j] - '0' + tmp_res/10) + '0'
		}
	}
	if res[0] == '0' {
		res = res[1:]
	}
	return string(res)

}
func main() {
	fmt.Println(multiply("999", "999"))
}
