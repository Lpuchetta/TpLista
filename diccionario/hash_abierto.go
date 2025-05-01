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

	factorDeCarga := float64(h.cantidad) / float64(len(h.casillas))
	if factorDeCarga > _FACTOR_CARGA_SUP{
		h.redimensionar(len(h.casillas) * 2)
	}
	
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	par, encontrado := h.obtenerParClaveValor(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	it, _ := h.buscar(clave)
	it.Borrar()
	
	valor := par.valor
	h.cantidad--
	factorDeCarga := float64(h.cantidad)/ float64(len(h.casillas))
	if factorDeCarga < _FACTOR_CARGA_INF && len(h.casillas) > _CAPACIDAD_INICIAL {
			nuevoTam := len(h.casillas) / 2
			h.redimensionar(nuevoTam)
	}

	return valor
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	par,encontrado := h.obtenerParClaveValor(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	valor := par.valor
	return valor
}

func (h *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	if nuevoTam <= 0{
		return //Agrego esta linea por el caso en que el nuevo tamaño sea menor o igual a 0. Asi no redimensionamos y nos evitamos dividir por zero o cosas asi.
	}

	nuevas := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)
	for i := 0; i < nuevoTam; i++{ //Cambio esto porque el range no se puede hacer sobre enteros
		nuevas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	cantNueva := 0
	for _, casilla := range h.casillas {

		for it := casilla.Iterador(); it.HaySiguiente(); it.Siguiente() {
			par := it.VerActual()
			indice := h.indexDe(par.clave, nuevoTam)
			nuevas[indice].InsertarUltimo(par)
			cantNueva++
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

func (h *hashAbierto[K,V]) obtenerParClaveValor(clave K) (parClaveValor[K,V], bool) {
	it, encontrado := h.buscar(clave)
	if !encontrado {
		return parClaveValor[K,V]{},false
		
	}
	return it.VerActual(), true
}

func (h *hashAbierto[K,V]) Iterar(visitar func(clave K, valor V) bool){
	it := h.Iterador()

	for it.HaySiguiente(){
		clave, valor := it.VerActual()
		if !visitar(clave, valor){
			return
		}
		it.Siguiente()
	}
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

func (it *iteradorHash[K, V]) HaySiguiente() bool{
	return it.avanzarCasilla()
}

func (it *iteradorHash[K, V]) Siguiente(){
	if !it.avanzarCasilla(){
		panic("El iterador termino de iterar")
	}
	it.iterLista.Siguiente()
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.avanzarCasilla(){
		panic("El iterador termino de iterar")
	}
	par := it.iterLista.VerActual()
	return par.clave, par.valor
}

func (it *iteradorHash[K, V]) avanzarCasilla() bool{
	for it.casilla < len(it.hash.casillas){
		if it.iterLista.HaySiguiente(){
			return true
		}
		it.casilla++
		if it.casilla < len(it.hash.casillas){
			it.iterLista = it.hash.casillas[it.casilla].Iterador()
		}
	}
	return false
}

