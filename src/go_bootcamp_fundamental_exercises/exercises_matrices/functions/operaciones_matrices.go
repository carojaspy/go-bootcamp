package functions

import ("fmt")

func suma_matrices(m1 [][]int, m2 [][]int) [][]int {
	/* */ 
	fmt.Println("suma_matrices")
	m3 := [][]int{}
	for i, row := range m1 {
		tmp_row := []int{}
		for j := range row {
			tmp_row = append(tmp_row, m1[i][j]+m2[i][j])
			fmt.Println(tmp_row)
		} // end col
		m3 = append(m3, tmp_row)
	} // end row
	fmt.Println("Matriz final", m3)
	return m3
}

func resta_matrices(m1 [][]int, m2 [][]int) [][]int {
	/* */ 
	fmt.Println("resta_matrices")
	m3 := [][]int{}
	for i, row := range m1 {
		tmp_row := []int{}
		for j := range row {
			tmp_row = append(tmp_row, m1[i][j]-m2[i][j])
			fmt.Println(tmp_row)
		} // end col
		m3 = append(m3, tmp_row)
	} // end row
	fmt.Println("Matriz final", m3)
	return m3
}

func mult_matrices(m1 [][]int, m2 [][]int) [][]int {
	/* */ 
	fmt.Println("mult_matrices")
	m3 := [][]int{}
	fmt.Println(m3)
	return nil
}

func determinante_matrix(m1 [][]int) [][]int {
	/* */ 
	fmt.Println("determinante_matrix")
	return nil
}





