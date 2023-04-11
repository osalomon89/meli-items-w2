package main

import (
	"fmt"
	"time"
)

//solo creacion
func actualizarCreateAt(item *Item){
	item.CreatedAt = time.Now()
}

func actualizarUpdateAt(item *Item){
	item.UpdateAt = time.Now()
}

func actualizarId(item *Item){
	item.ID = obtenerId()
}

//agregar
func appendItemToArticulos(item Item){
	articulos = append(articulos, item)
}

//no puntero
//actualizar 
func updateItemNuevo(item Item){
	//le pasamos el puntero
	actualizarUpdateAt(&item)
	for i, v := range articulos {
		if v.ID == item.ID {
			articulos[i] = item

		}
	}
}


// codigo debe ser unico
// funcion repetido
func codigoRepetido(item *Item) bool {
	var repetido bool
	//validar si es desigual a nill
	if item == nil{
		return true
	}
	for _, valor := range articulos {
		if valor.Code == item.Code {
			repetido = true
		}
	}
	return repetido
}

//id autoincremental, obtiene el proximo ID libre
func obtenerId() int {
	var idSiguiente int
	for _, itemAnterior := range articulos {
		if idSiguiente < itemAnterior.ID {
			idSiguiente = itemAnterior.ID
		}
	}
	idSiguiente += 1
	return idSiguiente
}

// validar status
func validateStatus(item *Item) {
	if item.Stock > 0 {
		item.Status = "ACTIVE"
	} else {
		item.Status = "INACTIVE"
	}
}

func informacionCompleta(item *Item) error{
	if item.Code == "" || item.Title == "" || item.Description == "" || item.Price == 0 || item.Status == "" {
		return fmt.Errorf("el campo no debe estar vacio")
	}
	return nil
}





