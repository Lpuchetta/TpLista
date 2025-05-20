package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return nil
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return nil
}
