package v1

func ErrGetRequest() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal membaca request data",
			EN: "Failed to parse data",
		}))
}

func ErrCannotUpdateData() *ResponseError {
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: "Tidak diperbolehkan mengubah data",
		EN: "Changing data is not allowed",
	}))
}
