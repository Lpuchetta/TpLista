package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarPrimeroInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(5)
	require.Equal(t, 5, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())

	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())

	lista.InsertarUltimo(6)
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 3, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())

	lista.InsertarUltimo(7)
	require.Equal(t, 7, lista.VerUltimo())
	require.Equal(t, 3, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 5, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 6, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 6, lista.BorrarPrimero())
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 7, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	n := 10000
	for i := 0; i < n/2; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i+1, lista.Largo())
	}
	for i := n / 2; i < n; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i+1, lista.Largo())
	}
	require.Equal(t, n, lista.Largo())
	largoEsperado := n
	for i := n/2 - 1; i >= 0; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
		largoEsperado--
		require.Equal(t, largoEsperado, lista.Largo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, n-1, lista.VerUltimo())
	}
	for i := n / 2; i < n; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
		largoEsperado--
		require.Equal(t, largoEsperado, lista.Largo())
		if largoEsperado > 0 {
			require.False(t, lista.EstaVacia())
			require.Equal(t, n-1, lista.VerUltimo())
		}
	}
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestCasosBordes(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
	lista.InsertarPrimero(44)
	require.Equal(t, 44, lista.VerPrimero())
	require.Equal(t, 44, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 44, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())

	lista.InsertarPrimero(100)
	require.Equal(t, 100, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})

	lista.InsertarUltimo(30)
	require.Equal(t, 30, lista.VerUltimo())
	require.Equal(t, 30, lista.VerPrimero())
	require.Equal(t, 30, lista.BorrarPrimero())

	lista.InsertarPrimero(7)
	lista.InsertarPrimero(7)
	lista.InsertarUltimo(7)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 7, lista.BorrarPrimero())
}

func TestDiferentesTiposDeDatos(t *testing.T) {
	listaStr := TDALista.CrearListaEnlazada[string]()

	require.True(t, listaStr.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaStr.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaStr.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaStr.BorrarPrimero()
	})
	require.Equal(t, 0, listaStr.Largo())
	listaStr.InsertarPrimero("A")
	require.Equal(t, 1, listaStr.Largo())
	require.False(t, listaStr.EstaVacia())
	require.Equal(t, "A", listaStr.VerPrimero())
	require.Equal(t, "A", listaStr.VerUltimo())
	listaStr.InsertarUltimo("B")
	require.Equal(t, "B", listaStr.VerUltimo())
	require.Equal(t, "A", listaStr.VerPrimero())
	require.Equal(t, 2, listaStr.Largo())
	require.False(t, listaStr.EstaVacia())
	require.Equal(t, "A", listaStr.BorrarPrimero())
	require.Equal(t, "B", listaStr.BorrarPrimero())
	require.True(t, listaStr.EstaVacia())

	listaBool := TDALista.CrearListaEnlazada[bool]()
	require.True(t, listaBool.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaBool.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaBool.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaBool.BorrarPrimero()
	})
	require.Equal(t, 0, listaBool.Largo())
	listaBool.InsertarPrimero(true)
	listaBool.InsertarUltimo(false)
	require.Equal(t, true, listaBool.VerPrimero())
	require.Equal(t, false, listaBool.VerUltimo())
	require.Equal(t, 2, listaBool.Largo())
	require.False(t, listaBool.EstaVacia())
	require.Equal(t, true, listaBool.BorrarPrimero())
	require.Equal(t, false, listaBool.BorrarPrimero())
	require.True(t, listaBool.EstaVacia())

	type Persona struct {
		Nombre string
		Edad   int
	}
	listaPersona := TDALista.CrearListaEnlazada[Persona]()
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaPersona.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaPersona.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		listaPersona.BorrarPrimero()
	})
	require.True(t, listaPersona.EstaVacia())
	require.Equal(t, 0, listaPersona.Largo())

	p1 := Persona{Nombre: "Alan", Edad: 58}
	p2 := Persona{Nombre: "Barbara", Edad: 12}
	listaPersona.InsertarPrimero(p1)
	listaPersona.InsertarUltimo(p2)
	require.False(t, listaPersona.EstaVacia())
	require.Equal(t, 2, listaPersona.Largo())
	require.Equal(t, p1, listaPersona.VerPrimero())
	require.Equal(t, p2, listaPersona.VerUltimo())
	require.Equal(t, p1, listaPersona.BorrarPrimero())
	require.Equal(t, p2, listaPersona.BorrarPrimero())
	require.True(t, listaPersona.EstaVacia())
	require.Equal(t, 0, listaPersona.Largo())

}

func TestCrearIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	it := lista.Iterador()

	require.False(t, it.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Siguiente()
	})

}

func TestIteradorInsertarSobreListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	it := lista.Iterador()

	it.Insertar(467)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 467, lista.VerPrimero())
	require.Equal(t, 467, it.VerActual())
	require.True(t, it.HaySiguiente())

	it.Siguiente()

	require.False(t, it.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Siguiente()
	})

}

func TestIteradorInsertaAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	it := lista.Iterador()

	it.Insertar(467)

	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 467, lista.VerPrimero())
	require.Equal(t, 467, it.VerActual())
	require.True(t, it.HaySiguiente())

}

func TestIteradorInsertaAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	it := lista.Iterador()
	for ; it.HaySiguiente(); it.Siguiente() {

	}

	it.Insertar(467)

	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 467, lista.VerUltimo())
	require.Equal(t, 467, it.VerActual())
	require.True(t, it.HaySiguiente())

}

func TestIteradorInsertarVariosValores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)

	it := lista.Iterador()
	it.Siguiente()

	it.Insertar(467)

	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 467, it.VerActual())
	require.True(t, it.HaySiguiente())

	it.Siguiente()
	it.Insertar(101)

	require.Equal(t, 6, lista.Largo())
	require.Equal(t, 101, it.VerActual())
	require.True(t, it.HaySiguiente())

	for ; it.HaySiguiente(); it.Siguiente() {

	}

	it.Insertar(23)

	require.Equal(t, 7, lista.Largo())
	require.Equal(t, 23, lista.VerUltimo())
	require.Equal(t, 23, it.VerActual())
	require.True(t, it.HaySiguiente())

	it.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Siguiente()
	})

}

func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	it := lista.Iterador()

	borrado := it.Borrar()

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, borrado)
	require.Equal(t, 2, it.VerActual())
	require.True(t, it.HaySiguiente())

}

func TestIteradorBorraUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	it := lista.Iterador()

	for i := 0; i < 3; i++ {
		it.Siguiente()
	}

	borrado := it.Borrar()

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 4, borrado)
	require.False(t, it.HaySiguiente())

}

func TestIteradorBorraAlTerminarDeIterarLanzaError(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	it := lista.Iterador()

	for it.HaySiguiente() {
		it.Siguiente()
	}

	require.False(t, it.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Siguiente()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		it.Borrar()
	})

}

func TestIteradorBorrarElementoDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(467)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	it := lista.Iterador()

	it.Siguiente()
	it.Siguiente()

	borrado := it.Borrar()

	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 467, borrado)
	require.Equal(t, 3, it.VerActual())
	require.True(t, it.HaySiguiente())

}

func TestIteradorInternoSobreListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	cont := 0
	lista.Iterar(func(dato int) bool {
		cont++
		return true
	})

	require.Equal(t, 0, cont)
}

func TestIteradorInternoSobreListaConUnSoloElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(467)

	suma := 0
	lista.Iterar(func(dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, 467, suma)
}

func TestIteradorInternoSeDetieneEnElPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}

	cont := 0
	lista.Iterar(func(dato int) bool {
		cont++
		return false
	})

	require.Equal(t, 1, cont)
}

func TestIteradorInternoSeDetieneEnElTercero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}

	cont := 0
	lista.Iterar(func(dato int) bool {
		if cont == 3 {
			return false
		}
		cont++
		return true
	})

	require.Equal(t, 3, cont)
}

func TestIteradorInternoVisitaTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	nums := []int{467, 23, 101, 19}

	for _, num := range nums {
		lista.InsertarUltimo(num)
	}

	var vistos []int
	lista.Iterar(func(dato int) bool {
		vistos = append(vistos, dato)
		return true
	})

	require.Equal(t, nums, vistos)
}

func TestIteradorInternoSumaElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 100; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, 5050, suma)

}

func TestIteradorInternoSumaParcialElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 6; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(dato int) bool {
		if suma <= 10 {
			suma += dato
			return true
		}
		return false
	})

	require.Equal(t, 15, suma)

}

func TestIteradorInternoSumarPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 6; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(dato int) bool {
		if dato%2 == 0 {
			suma += dato
		}
		return true
	})

	require.Equal(t, 12, suma)

}

func TestIteradorInternoSumarLosPrimerosTresPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 100; i++ {
		lista.InsertarUltimo(i)
	}

	suma, cont := 0, 0
	lista.Iterar(func(dato int) bool {
		if cont == 3 {
			return false
		}

		if dato%2 == 0 {
			suma += dato
			cont++
		}

		return true
	})

	require.Equal(t, 12, suma)

}

func TestIteradorInternoContarImpares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < 100; i++ {
		lista.InsertarUltimo(i)
	}

	cont := 0
	lista.Iterar(func(dato int) bool {
		if dato%2 == 1 {
			cont++
		}
		return true
	})

	require.Equal(t, 50, cont)

}

func TestIteradorInternoBusquedaElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	nums := []int{467, 23, 101, 19}

	for _, num := range nums {
		lista.InsertarUltimo(num)
	}

	elemBuscado := 101
	pertenece := false
	lista.Iterar(func(dato int) bool {
		if dato == elemBuscado {
			pertenece = true
			return false
		}
		return true
	})

	require.True(t, pertenece)
}

func TestIteradorInternoConcatenaStrings(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarUltimo("Esto es")
	lista.InsertarUltimo(" una ")
	lista.InsertarUltimo("prueba sobre el ")
	lista.InsertarUltimo("iterador inter")
	lista.InsertarUltimo("no :)")

	var texto string
	lista.Iterar(func(s string) bool {
		texto += s
		return true
	})
	require.Equal(t, "Esto es una prueba sobre el iterador interno :)", texto)
}
