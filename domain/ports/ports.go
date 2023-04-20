package ports

import (
	dom "github.com/osalomon89/neocamp-meli-w2/domain"
)


type UseCasesService interface{
	UpdateItem(dom.Item)error
	GetItem(dom.Item)error
	DeleteItem(int)bool
}


type Repository interface{
	UpdateItem(dom.Item) error
}