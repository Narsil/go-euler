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
		case 11:
			return Euler11()
		case 12:
			return Euler12()
        case 13:
            return Euler13()
        case 14:
            return Euler14()
        case 15:
            return Euler15()
        case 16:
            return Euler16()
        case 17:
            return Euler17()
        case 18:
            return Euler18()
		/*
        case 19:
            return Euler19()
        case 20:
            return Euler20()
		*/
		default:
			return "Problem not solved"
	}
	return "Impossible"
}
