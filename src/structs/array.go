package structs

import "errors"

type Array struct{
	Data []string
	Lenght int
}

func NewArray(len int) *Array{
	return &Array{
		Data: make([]string, len),
		Lenght: len,
	}
}

func (array *Array) Aset(index int, value string) error {
    if index < 0 || index >= array.Lenght {
		return errors.New("-->incorrect index value")
	}
    array.Data[index] = value
    return nil
}

func (array *Array) Aget(index int) (string, error){
	if index < 0 || index >= array.Lenght {
        return "", errors.New("-->incorrect index value")
    }
    return array.Data[index], nil
}