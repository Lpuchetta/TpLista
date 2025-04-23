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
}

func TestInsertarPrimeroInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	lista.InsertarUltimo(21)
	require.Equal(t, 21, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 21, lista.VerPrimero())
	require.False(t, lista.EstaVacia())

	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 21, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
		
}



