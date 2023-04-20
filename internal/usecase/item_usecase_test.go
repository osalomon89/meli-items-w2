package usecase_test

import (
	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/domain/port"
	"gigigarino/challengeMELI/internal/usecase"
	"reflect"
	"testing"
)



func Test_itemUsecase_Index(t *testing.T) {
	type fields struct {
		repo port.ItemRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []domain.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewItemUsecase(nil)
			if got := u.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemUsecase.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemUsecase_GetListaInicial(t *testing.T) {
	type fields struct {
		repo port.ItemRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []domain.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewItemUsecase(nil)
			if got := u.GetListaInicial(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemUsecase.GetListaInicial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemUsecase_GetAllItems(t *testing.T) {
	type fields struct {
		repo port.ItemRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []domain.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewItemUsecase(nil)
			if got := u.GetAllItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemUsecase.GetAllItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemUsecase_GetItemById(t *testing.T) {
	type fields struct {
		repo port.ItemRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewItemUsecase(nil)
			if got := u.GetItemById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemUsecase.GetItemById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemUsecase_AddItem(t *testing.T) {
	type fields struct {
		repo port.ItemRepository
	}
	type args struct {
		item domain.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewItemUsecase(nil)
			got, err := u.AddItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("itemUsecase.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itemUsecase.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
