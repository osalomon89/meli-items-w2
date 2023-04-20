package usecases

import (
"errors"
repo "github.com/osalomon89/neocamp-meli-w2/controller/repository"
dom "github.com/osalomon89/neocamp-meli-w2/domain"

)
type usecases struct{

}



func NewUseCases() *usecases{
	return &usecases{}
}


func (u *usecases) UpdateItem(item dom.Item)error{
	// var db []repo.itemRepository
	
	// for i := range db{
	// 	if db[i].ID == item.ID {
	// 		db[i] = item		//PREGUNTAR ESTO.
	// 		db[i].Title = item.Title
	// 		db[i].Description = item.Description
	// 		db[i].Code = item.Code
	// 		db[i].Price = item.Price

			//ctx.JSON(http.StatusOK, gin.H{"data": db[i]})
			return errors.New("se actualizó")
		// }
	// }
	
	return errors.New("no se actualizó")
}

func (u *usecases) GetItem(item dom.Item)error{
	
	return nil
}