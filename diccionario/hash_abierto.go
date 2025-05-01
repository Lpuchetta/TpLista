package diccionario

import (
	"fmt"
	"hash/fnv"
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

type iteradorHash[K comparable, V any] struct{
	hash  *hashAbierto[K,V]
	casilla int
	iterLista	TDALista.IteradorLista[parClaveValor[K, V]]
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

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, pertenece := h.buscar(clave)
	return pertenece
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	it, pertenece := h.buscar(clave)
	if pertenece {
		par := it.VerActual()
		par.valor = valor
		it.Borrar()
		it.Insertar(par)
		return 
	}
	nuevo := parClaveValor[K, V]{clave, valor}
	it.Insertar(nuevo)
	h.cantidad++
	//Cambie esta parte porque si la clave ya esta en el hash actualizamos el valor y no aumentamos la cantindad, antes pasaba
	//Que aunque la clave estaba se pisaba el valor y se aumentamaba igual

	if float64(h.cantidad/len(h.casillas)) > _FACTOR_CARGA_SUP {
		nuevoTam := 2 * len(h.casillas)
		h.redimensionar(nuevoTam)
	}

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	it, encontrado := h.buscar(clave)
	var par parClaveValor[K, V]
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	} else {
		par = it.Borrar()
	}

	valor := par.valor
	h.cantidad--

	if float64(h.cantidad/len(h.casillas)) < _FACTOR_CARGA_INF {
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
	} else {
		par = it.VerActual()
	}
	valor := par.valor

	return valor
}

func (h *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	nuevas := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)
	for i := 0; i < nuevoTam; i++{ //Cambio esto porque el range no se puede hacer sobre enteros
		nuevas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for _, casilla := range h.casillas {

		for it := casilla.Iterador(); it.HaySiguiente(); it.Siguiente() {
			par := it.VerActual()
			indice := h.indexDe(par.clave, nuevoTam)
			nuevas[indice].InsertarUltimo(par)
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

// indexDe define en que casilla del arreglo de listas enlazadas debe caer el par clave-valor.
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

// convertirABytes convierte una clave generica a []byte.
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K,V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K,V]{
		hash: h,
		casilla: 0,
		iterLista: h.casillas[0].Iterador(),
	}
}

func (it *iteradorHash[K, V]) avanzarASiguienteNoVacia() bool {
	for it.casilla < len(it.hash.casillas) {
		if it.iterLista.HaySiguiente() {
			return true
		}
		it.casilla++
		if it.casilla < len(it.hash.casillas) {
			it.iterLista = it.hash.casillas[it.casilla].Iterador()
		}
	}
	return false
}



func (it *iteradorHash[K, V]) HaySiguiente() bool{
	return it.avanzarASiguienteNoVacia()
}

func (it *iteradorHash[K, V]) Siguiente(){
	if !it.avanzarASiguienteNoVacia(){
		panic("El iterador termino de iterar")
	}
	it.iterLista.Siguiente()
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.avanzarASiguienteNoVacia(){
		panic("El iterador termino de iterar")
	}
	par := it.iterLista.VerActual()
	return par.clave, par.valor
}


