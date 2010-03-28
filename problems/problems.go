package projecteuler

func Euler(n int) string {
	switch n{
		case 1:
			return Euler1()
		case 2:
			return Euler2()
		case 3:
			return Euler3()
		case 4:
			return Euler4()
		case 5:
			return Euler5()
		case 6:
			return Euler6()
		case 7:
			return Euler7()
		case 8:
			return Euler8()
		case 9:
			return Euler9()
		case 10:
			return Euler10()
		default:
			return "Problem not solved"
	}
	return "Impossible"
}
