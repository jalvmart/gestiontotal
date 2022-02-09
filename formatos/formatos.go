package formatos

import (
	"strconv"
	"strings"
)

func formatoFecha(x string) string {
	dato := strings.Split(x, "-")
	return string(dato[2] + "-" + dato[1] + "-" + dato[0])
}

func formatoNumero(dato float64) string {

	a := strconv.FormatFloat(dato, 'f', 2, 64)
	b := len(a)
	r := ""
	rango1 := []int{7, 8, 9}
	rango2 := []int{10, 11, 12}
	rango3 := []int{13, 14, 15}

	if b <= 6 {
		r = a
	}
	for i, v := range rango1 {
		i++
		if b == v {
			r = a[0:i] + "," + a[i:]
		}
	}

	for i, v := range rango2 {
		i++
		if b == v {
			r = a[0:i] + "," + a[i:i+3] + "," + a[i+3:]
		}
	}

	for i, v := range rango3 {
		i++
		if b == v {
			r = a[0:i] + "," + a[i:i+3] + "," + a[i+3:i+6] + "," + a[i+6:]
		}
	}

	return string(r)

}
