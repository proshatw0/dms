package structs

import "errors"

type Array struct {
	Data   []string
	Lenght int
}

func NewArray(len int) *Array {
	return &Array{
		Data:   make([]string, len),
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

func (array *Array) Aget(index int) (string, error) {
	if index < 0 || index >= array.Lenght {
		return "", errors.New("-->incorrect index value")
	}
	return array.Data[index], nil
}

func (array *Array) Aindex(value string) (int, error) {
	for i := 0; i < array.Lenght; i++ {
		if array.Data[i] == value {
			return i, nil
		}
	}
	return -1, errors.New("-->element not found")
}

func (array *Array) Adel(index int) (string, error) {
	if index < 0 || index >= array.Lenght {
		return "", errors.New("-->incorrect index value")
	}
	value := array.Data[index]
	array.Data[index] = ""
	return value, nil
}

func (array *Array) Adel_value(value string) (string, error) {
	index, err := array.Aindex(value)
	if err != nil {
		return "", err
	}
	array.Data[index] = ""
	return value, nil
}

func (array *Array) Apush(value string) error {
	index, err := array.Aindex("")
	if err != nil {
		return errors.New("-->array is full")
	}
	array.Data[index] = value
	return nil
}

func (array *Array) index_last_element() (int, error) {
	for i := array.Lenght - 1; i >= 0; i-- {
		if array.Data[i] != "" {
			return i, nil
		}
	}
	return -1, errors.New("--array is clear")
}

func (array *Array) Apop() (string, error) {
	index, err := array.index_last_element()
	if err != nil {
		return "", err
	}
	value := array.Data[index]
	array.Data[index] = ""
	return value, nil
}
