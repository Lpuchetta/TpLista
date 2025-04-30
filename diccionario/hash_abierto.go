package diccionario

import (
	"fmt"
	"hash/fnv" //Una funcion de hash que prove go, si queres fijate en esta pagina. https://pkg.go.dev/hash/fnv#New64a
	TDALista "tdas/lista"
)

const (
	_FACTOR_CARGA      = 0.7
	_CAPACIDAD_INICIAL = 7
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	casillas []TDALista.Lista[parClaveValor[K, V]]
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	casillas := make([]TDALista.Lista[parClaveValor[K, V]], _CAPACIDAD_INICIAL)
	for i := range casillas {
		casillas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{
		casillas: casillas,
		cantidad: 0,
	}
}


func (h *hashAbierto[K, V]) buscarIterador(clave K) (TDALista.IteradorLista[parClaveValor[K, V]], bool) {
	indice := h.indexDe(clave)
	lista := h.casillas[indice]
	it := lista.Iterador()
	for it.HaySiguiente() {
		if it.VerActual().clave == clave {
			return it, true
		}
		it.Siguiente()
	}
	return nil, false
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, ok := h.buscarIterador(clave)
	return ok
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	_, ok := h.buscarIterador(clave)
	if !ok{
		panic("La clave no pertence al diccionario")
	}

	indice := h.indexDe(clave)
	lista := h.casillas[indice]

	nuevo := parClaveValor[K, V]{
		clave: clave,
		valor: valor,
	}

	lista.InsertarUltimo(nuevo)
	h.cantidad++

	// TODO: Acá hay que pensar en el factor de carga y en la redimensión.

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	it, ok := h.buscarIterador(clave)
	if !ok{
		panic ("La clave no pertenece al dicciona")
	}
	borrado := it.Borrar()
	valor := borrado.valor
	h.cantidad--

	//TODO: Acá hay que pensar en el factor de carga y en la redimensión.

	return valor
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	it, ok := h.buscarIterador(clave)
	if !ok{
		panic("La clave no pertence al diccionario")
	}
	return it.VerActual().valor
}

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
