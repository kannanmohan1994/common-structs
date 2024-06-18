package commonstructs

type BusinessError struct {
	Err string
}

func (b *BusinessError) Error() string {
	return b.Err
}
