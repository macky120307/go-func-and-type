package main

func main() {
	ms := []int{19, 86, 1, 12}

	var sum int
	for i := 0; i < len(ms); i++ {
		sum += ms[i]
	}

	println(sum)

	/* -------------------- */
	
	type Game struct {
		user 	string
		score 	int
	}
	
	sample := Game {
		user:  "user1",
		score: 80,
	}

	println(sample.user)

	/* -------------------- */

	n1, m1 := swap(10, 20)
	println(n1, m1)

	/* -------------------- */

	n2, m2 := 10, 20
	swap2(&n2, &m2)
	println(n2, m2)
}

func swap(n, m int) (n1, m1 int) {
	n1, m1 = m, n
	return
}

func swap2(n, m *int) (n1, m1 *int) {
	*n, *m = *m, *n
	return
}
