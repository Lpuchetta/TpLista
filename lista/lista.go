package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no contiene elementos.
	EstaVacia() bool

	// InsertarPrimero inserta un elemento en la primera posición.
	InsertarPrimero(T)

	// InsertarUltimo inserta un elemento en la última posición.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y lo devuelve.
	// Si la lista está vacía, lanza un panic: "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero ve el primer elemento de la lista y lo devuelve
	// Si la lista esta vacia, lanza un panic: "La lista esta vacia".
	VerPrimero() T

	// VerPrimero devuelve el primer elemento de la lista sin eliminarlo.
	// Si la lista está vacía, lanza un panic: "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recorre la lista desde el primero al ultimo elemento.
	// llamado a la funcion visitar por cada uno. Si visitar devuelve false, se corta la iteracion.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador para recorrer la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual sobre el que está parado el iterador.
	// Si el iterador ya recorrió toda la lista, lanza un panic: "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si quedan elementos por recorrer.
	HaySiguiente() bool

	// Siguiente avanza una posición en la lista.
	// Si el iterador ya recorrió toda la lista, lanza un panic: "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un elemento en la posición actual del iterador.
	Insertar(T)

	// Borrar elimina el elemento actual de la lista y lo devuelve.
	// Si el iterador ya recorrió toda la lista, lanza un panic: "El iterador termino de iterar".
	Borrar() T
}
