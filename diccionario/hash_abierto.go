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
