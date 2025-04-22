package lista

type Lista[T any] interface {

	// EstaVacia revisa si la lista esta vacia y devuelve un booleano
	EstaVacia() bool

	// InsertarPrimero inserta un elemento en la primer posicion
	InsertarPrimero(T)

	// InsertarUltimo inserta un elemento en la ultima posicion
	InsertarUltimo(T)

	// BorrarPrimero borra el primer elemento de la lista
	// Si la lista esta se devuelve un panic "La lista esta vacia
	BorrarPrimero() T

	// VerPrimero ve el primer elemento de la lista y lo devuelve
	// Si la lista esta se devuelve un panic "La lista esta vacia
	VerPrimero() T

	// VerUltimo ve el ultimo elemento de la lista y lo devuelve
	// Si la lista esta se devuelve un panic "La lista esta vacia
	VerUltimo() T

	// Largo devuelve la longiud de la lista
	Largo() int

	// Iterar permite recorrer la lista con una funcion
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador para recorrer la lista
	Iterador() IteradorLista[T]
}
