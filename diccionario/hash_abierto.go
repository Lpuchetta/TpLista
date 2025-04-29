package diccionario

import (
	TDALista "tdas/lista"
)

const _CAPACIDAD_INICIAL = 23 // Es mejor que sea un número primo.

type hashAbierto[K comparable, V any] struct {
	casillas []TDALista.Lista[K]
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		casillas: make([]TDALista.Lista[K], _CAPACIDAD_INICIAL),
		cantidad: 0,
	}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {

}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	return false
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {

}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return nil
}
