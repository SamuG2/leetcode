package main

import "fmt"

// idea: buscar los máximos de altura me ir calculando uno por uno cuanta lluvia hay entre los dos. Si haciendo esto se quintan otros máximos ya no se buscan
// func trap(height []int) int {

// 	var heights map[int][]int
// 	max_h := 0
// 	for i :=range height {
// 		max_h = max(max_h, height[i])
// 		heights[height[i]] =append(heights[height[i]], i)
// 	}

// 	for max_h >0{

// 	}

// }

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

//idea 2: explorar de izquiera a derecha hasta encontrar máximo local (h[i]>h[i-1] && h[i]> h[i+1]). Una vez encontrado buscarel siguiente máximo local, calcular y restar. Seguir buscando maximos
func trap(height []int) int {
	height = append(height, 0) //metemos un cero al final
	res := 0
	l := len(height) - 1
	for i := 0; i < l; i++ {
		if height[i] > height[i+1] {
			base := 0
			temp_base := 0
			temp_res := 0
			max_h_idx := l
			for j := i + 1; j < l; j++ {
				if height[j] >= height[i] {
					temp_res = min(height[i], height[j])*(j-i-1) - base
					i = j - 1
					break
				}
				if height[j] >= height[max_h_idx] {
					max_h_idx = j
					temp_base = base
				}

				base += height[j]
			}
			if temp_res == 0 && max_h_idx != 0 {
				temp_res = height[max_h_idx]*(max_h_idx-i-1) - temp_base
				i = max_h_idx - 1
			}
			res += temp_res
		}
	}
	return res
}

func main() {
	fmt.Println(trap([]int{4, 2, 3}))
}
