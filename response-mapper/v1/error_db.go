package v1

func ErrDB() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			EN: "An error occurred while querying db",
			ID: "Terjadi kesalahan pada saat query db",
		}))
}

func ErrUpdatedDB() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Terjadi kesalahan pada saat perbarui data ke db",
			EN: "An error occurred while updating db",
		}))
}

func ErrCreatedDB() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Terjadi kesalahan pada saat menambahkan data ke db",
			EN: "An error occurred while creating db",
		}))
}

func ErrDeletedDB() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Terjadi kesalahan pada saat menghapus data ke db",
			EN: "An error occurred while deleting db",
		}))
}

func ErrFailedSendEmail() *ResponseError {
	return NewError(ErrDatabase, NewResponseMultiLang(
		MultiLanguages{
			ID: "Gagal mengirim surel",
			EN: "Failed to send email",
		}))
}
