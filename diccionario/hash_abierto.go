package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const (
	_CAPACIDAD_INICIAL = 23 // Es mejor que sea un número primo.
	_FACTOR_CARGA      = 1.2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	casillas []TDALista.Lista[parClaveValor[K, V]]
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		casillas: make([]TDALista.Lista[parClaveValor[K, V]], _CAPACIDAD_INICIAL),
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

func (hash *hashAbierto[K, V]) buscar()

// indexDe define en que casilla del arreglo de listas enlazadas debe caer el par clave-valor.
func (h *hashAbierto[K, V]) indexDe(clave K) int {
	hv := h.hashClave(clave)
	return int(hv % uint64(len(h.casillas)))
}

// hashClave usa FNV-1a para generar un hash de la clave.
func (h *hashAbierto[K, V]) hashClave(clave K) uint64 {
	hf := fnv.New64a()
	hf.Write(convertirABytes(clave))
	return hf.Sum64()
}

// convertirABytes convierte una clave generica a []byte.
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
