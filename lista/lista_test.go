package lista_test

import(
	"testing"
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista" 
	
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t,"La lista esta vacia", func(){
		lista.BorrarPrimero()
	})
	require.PanicsWithValue(t,"La lista esta vacia", func(){
		lista.VerPrimero()
	})
	require.PanicsWithValue(t,"La lista esta vacia", func(){
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
	
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.BorrarPrimero()
	})
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	n := 10000
	for i := 0; i < n/2; i++{
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i + 1, lista.Largo())
	} 
	for i := n / 2; i < n; i++{
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i + 1, lista.Largo())
	}
	require.Equal(t, n, lista.Largo())
	largoEsperado := n
	for i := n / 2 - 1; i >= 0; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
		largoEsperado--
		require.Equal(t, largoEsperado , lista.Largo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, n - 1, lista.VerUltimo())
	}
	for i := n / 2; i < n ; i++{
		require.Equal(t, i, lista.BorrarPrimero())
		largoEsperado--
		require.Equal(t,largoEsperado,lista.Largo())
		if largoEsperado > 0{
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
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.BorrarPrimero()
	})
}

func TestCasosBordes(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func(){
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func(){
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
	require.PanicsWithValue(t, "La lista esta vacia", func(){
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
