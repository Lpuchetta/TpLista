package diccionario

import (
	"fmt"
	"hash/fnv" //Una funcion de hash que provee go. Referencia: https://pkg.go.dev/hash/fnv#New64a
	TDALista "tdas/lista"
)

const (
	_FACTOR_CARGA_SUP  = 1.2
	_FACTOR_CARGA_INF  = 0.3
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

type iterDiccionario[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	itLista   TDALista.IteradorLista[parClaveValor[K, V]]
	posActual int
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

func CrearIteradorDiccionario[K comparable, V any](h *hashAbierto[K, V]) IterDiccionario[K, V] {
	posActual := 0
	for posActual < len(h.casillas) && h.casillas[posActual].EstaVacia() {
		posActual++
	}
	var itLista TDALista.IteradorLista[parClaveValor[K, V]]
	if posActual < len(h.casillas) {
		itLista = h.casillas[posActual].Iterador()
	}
	return &iterDiccionario[K, V]{
		hash:      h,
		itLista:   itLista,
		posActual: posActual,
	}
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, pertenece := h.buscar(clave)
	return pertenece
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	it, encontrado := h.buscar(clave)
	if encontrado {
		par := it.VerActual()
		par.valor = valor
		it.Borrar()
		it.Insertar(par)
		return
	}

	nuevo := parClaveValor[K, V]{clave, valor}
	it.Insertar(nuevo)
	h.cantidad++

	if float64(h.cantidad)/float64(len(h.casillas)) > _FACTOR_CARGA_SUP {
		nuevoTam := len(h.casillas) * 2
		h.redimensionar(nuevoTam)
	}

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	it, encontrado := h.buscar(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	par := it.Borrar()

	valor := par.valor
	h.cantidad--
	factorDeCarga := float64(h.cantidad) / float64(len(h.casillas))
	if factorDeCarga < _FACTOR_CARGA_INF && len(h.casillas) > _CAPACIDAD_INICIAL {
		nuevoTam := len(h.casillas) / 2
		h.redimensionar(nuevoTam)
	}

	return valor
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	it, encontrado := h.buscar(clave)
	var par parClaveValor[K, V]
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	par = it.VerActual()

	valor := par.valor

	return valor

}

func (h *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for _, casilla := range h.casillas {
		for it := casilla.Iterador(); it.HaySiguiente(); it.Siguiente() {
			par := it.VerActual()
			clave, valor := par.clave, par.valor
			if !visitar(clave, valor) {
				return
			}
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return CrearIteradorDiccionario(h)
}

func (h *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	if nuevoTam == len(h.casillas) || nuevoTam <= 0 {
		return
	}

	nuevas := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)
	for i := range nuevas {
		nuevas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for _, casilla := range h.casillas {
		for it := casilla.Iterador(); it.HaySiguiente(); {
			par := it.VerActual()
			indice := h.indexDe(par.clave, nuevoTam)
			nuevas[indice].InsertarUltimo(par)
			it.Siguiente()
		}
	}

	h.casillas = nuevas
}

func (h *hashAbierto[K, V]) buscar(clave K) (TDALista.IteradorLista[parClaveValor[K, V]], bool) {
	indice := h.indexDe(clave, len(h.casillas))
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

func (it *iterDiccionario[K, V]) HaySiguiente() bool {
	return it.itLista != nil && it.itLista.HaySiguiente()
}

func (it *iterDiccionario[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	par := it.itLista.VerActual()
	return par.clave, par.valor
}

func (it *iterDiccionario[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.itLista.Siguiente()

	if !it.itLista.HaySiguiente() {
		it.buscarProxPos()
	}
}

func (it *iterDiccionario[K, V]) buscarProxPos() {
	it.posActual++
	for it.posActual < len(it.hash.casillas) && it.hash.casillas[it.posActual].EstaVacia() {
		it.posActual++
	}
	if it.posActual < len(it.hash.casillas) {
		it.itLista = it.hash.casillas[it.posActual].Iterador()
	} else {
		it.itLista = nil
	}

}

func (h *hashAbierto[K, V]) indexDe(clave K, tamCasillas int) int {
	hv := h.hashClave(clave)
	return int(hv % uint64(tamCasillas))
}

// hashClave usa FNV-1a para generar un hash de la clave.
func (h *hashAbierto[K, V]) hashClave(clave K) uint64 {
	hf := fnv.New64a()
	hf.Write(convertirABytes(clave))
	return hf.Sum64()
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
