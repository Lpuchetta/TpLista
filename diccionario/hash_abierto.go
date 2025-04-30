package diccionario

import (
	"fmt"
	"hash/fnv" //Una funcion de hash que prove go, si queres fijate en esta pagina. https://pkg.go.dev/hash/fnv#New64a
	TDALista "tdas/lista"
)

const (
	_FACTOR_CARGA      = 1.2
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

//Aca tengo la duda de si esta bien o no crear el hash vacio. Entiendo que no tiene mucho sentido porque en la interfaz en ningun lugar aclara que
//si se quiere hacer algo sobre el diccionario tire un panic en caso de que este vacio. Revisar.

// No hay drama porque cuando lo creas, no puede ser nil
func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	if h.Cantidad() == 0 {
		return false
	}
	indice := h.indexDe(clave)
	lista := h.casillas[indice]
	if lista.EstaVacia() {
		return false
	}
	it := lista.Iterador()
	for it.HaySiguiente() {
		if it.VerActual().clave == clave {
			return true
		}
		it.Siguiente()
	}
	return false

	// Acá se reduce a:

	// indice := h.indexDe(clave)
	// _, pertenece := h.buscar(clave, indice)
	// return pertenece
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	indice := h.indexDe(clave) // Problema1: Tuve que hacerlo así porque sino no tenía el índice y no podía acceder a la lista en (*)
	it, pertenece := h.buscar(clave, indice)
	if pertenece {
		par := it.VerActual()
		par.valor = valor
	} else {
		nuevo := parClaveValor[K, V]{
			clave: clave,
			valor: valor,
		}

		it.Insertar(nuevo)
	}

	h.cantidad++

	// TODO: Acá hay que pensar en el factor de carga y en la redimensión.

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	if !h.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	indice := h.indexDe(clave)
	lista := h.casillas[indice]

	var valor V
	for it := lista.Iterador(); it.HaySiguiente(); it.Siguiente() {
		actual := it.VerActual()
		if actual.clave == clave {
			borrado := it.Borrar()
			valor = borrado.valor
		}
	}
	h.cantidad--

	//TODO: Acá hay que pensar en el factor de carga y en la redimensión.

	return valor
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	if !h.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	indice := h.indexDe(clave)
	lista := h.casillas[indice]
	it := lista.Iterador()
	var valor V
	for it.HaySiguiente() {
		actual := it.VerActual()
		if actual.clave == clave {
			valor = actual.valor
			break
		}
		it.Siguiente()
	}
	return valor
}

func (h *hashAbierto[K, V]) buscar(clave K, indice int) (TDALista.IteradorLista[parClaveValor[K, V]], bool) {
	lista := h.casillas[indice]
	it := lista.Iterador()
	for ; it.HaySiguiente(); it.Siguiente() {
		actual := it.VerActual()
		if actual.clave == clave {
			return it, true
		}
	}
	return it, false
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
