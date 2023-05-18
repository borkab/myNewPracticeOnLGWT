package arrays

func SumArray(numbers [5]int) int { //a SumArray funcio argumense egy elore meghatarozott 5 elemu lista, es egy egesz szamot ad vissza
	sum := 0                         //letrehozok egy sum nevu valtozot aminek a 0 erteket adom
	for _, number := range numbers { //letrehozok egy uj valtozot, aminek az erteke a numbers nevu 5elemu list aktualis erteke a for loopban
		sum += number //vegigmegyek a numbers lista osszes elemet, amikor az elso elemhez erek, azt az erteket hozzaadom a sum valtozo ertekehez
	}
	return sum // a funcio visszaadja a sum vegso erteket, ami a numbers 5elemu lista elemeinek az osszege
}

func SumSlice(numbers []int) int { //a SumSlice funkcio argumense egy slice of ints, es egy integert ad vissza
	sum := 0 // letrehozok egy int tipusu valtozot, aminek az erteke nulla

	for _, number := range numbers { //letrehozok egy number nevu valtozot, aminek megadom erteknek a numbers nevu listam aktualis elemenek az erteket
		sum += number //vegigmegyek a numbers nevu lista osszes elemen, es amikor az aktualis elemhez erek, ennek az elemnek az erteket hozzaadom a sum valtozom ertekehez.
	}
	return sum //a funcio kikopi a sum vegso erteket, amikor vegigert a loop-on,
	//tehat a numbers lista elemeinek az osszeget
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int //letrehozok egy sums nevu valtozot slice fo ints tipussal

	for _, numbers := range numbersToSum { //letrehozok egy numbers nevu egesz szamokbol allo listat, aminek megadom erteknek a numbersToSum listakat tartalmazo lista elso elemet. vegigmegyek a numberesToSum listam minden egyes elemen
		sums = append(sums, SumSlice(numbers)) //amikor az aktualis elemehez ertem, a korabban letrehozott sums nevu listamhoz hozzaadom a numbers lista elemeinek az osszeget.
	}
	return sums // a funcio kikopi a sums valtozo vegso erteket, amikor vegigert a loopon,
	// tehat a numbersToSum lista minden listaja osszegenek az osszeget.
}

func SumAllTails(numberesToSum ...[]int) []int { //a SumAllTails funcio argumense egy listakat tartalmazo lista
	var sums []int                          //letrehozok egy meg ures listat
	for _, numbers := range numberesToSum { //letrehozok egy numbers nevu listat, aminek az erteke mindig a numbersToSum listakat tartalmazo lista aktualis eleme
		//vegigmegyek a sumall listakat tartalmazo listan, es amikor az aktualis listahoz erek,
		if len(numbers) == 0 { //megnezem h a lista ures e. ha ures, akkor a sums listahoz hozzaadok egy 0 erteket
			sums = append(sums, 0)
		} else { //ha a lista nem ures, azaz mindne mas esetben:
			tail := numbers[1:]                 // a tail nevu lista aktualis elemenek megadom a numbers lista elemeit az elso elemevel kezdve
			sums = append(sums, SumSlice(tail)) //az elobb letrehozott lista elemeit osszeadom a Sumslice funkcioval, es az osszeget hozzaadom a sums nevu listamhoz
		}
	}
	return sums // a funkcio visszaadja a sums lista elemeit, ami egy lista a listak osszeadott elemeirol az 1. elemektol kezdve(a listak nulladik elemet nem adtuk hozza a listak elemeinek osszegehez)
}
