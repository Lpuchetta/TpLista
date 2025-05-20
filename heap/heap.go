package cola_prioridad

const _CAPACIDAD_INICIAL = 5

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, _CAPACIDAD_INICIAL),
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return nil
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return false
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {

}

func (cola *colaConPrioridad[T]) VerMax() T {

}

func (cola *colaConPrioridad[T]) Desencolar() T {

}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}
