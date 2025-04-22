package lista

type NodoLista[T any] struct{
	proximo *NodoLista[T]
	dato T
}

type ListaEnlazada[T any] struct{
	primero *NodoLista[T]
	ultimo *NodoLista[T]
	largo int
}

func CrearListaEnlazada[T any]() *ListaEnlazada[T]{
	return &ListaEnlazada[T]{
		primero : nil,
		ultimo : nil,
		largo : 0,
	}
}
