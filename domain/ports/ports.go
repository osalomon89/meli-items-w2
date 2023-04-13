package ports

import (
	
	dom "github.com/osalomon89/neocamp-meli-w2/domain"
	
)

type UseCasesService interface{
	UpdateItem(dom.Item)error
}


type Repository interface{
	UpdateItem(dom.Item) error
}