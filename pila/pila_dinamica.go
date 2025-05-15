package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	_TAM_INICIAL        = 5
	_FACTOR_REDIMENSION = 2
	_LIM_REDUCCION      = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _TAM_INICIAL),
		cantidad: 0,
	}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * _FACTOR_REDIMENSION)
	}
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	dato := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if pila.cantidad == len(pila.datos)/_LIM_REDUCCION {
		nuevoTam := len(pila.datos) / _FACTOR_REDIMENSION
		if nuevoTam >= _TAM_INICIAL {
			pila.redimensionar(nuevoTam)
		}
	}
	return dato
}

func (pila *pilaDinamica[T]) redimensionar(nuevoTam int) {
	nuevos := make([]T, nuevoTam)
	copy(nuevos, pila.datos)
	pila.datos = nuevos
}
