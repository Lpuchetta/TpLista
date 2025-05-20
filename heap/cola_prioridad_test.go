package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapifyArregloExistente(t *testing.T) {
	arr := []int{9, 3, 7, 1, 5}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 9, heap.VerMax())

	expectedOrder := []int{9, 7, 5, 3, 1}
	for _, expected := range expectedOrder {
		require.Equal(t, expected, heap.Desencolar())
	}
}

func TestElementosIguales(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	elementos := []int{5, 5, 5, 5}
	for _, e := range elementos {
		heap.Encolar(e)
	}

	for range elementos {
		require.Equal(t, 5, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}
