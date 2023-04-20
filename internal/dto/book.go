package dto

import "bibliotekaProject/internal/entity"

type BookRequest struct {
	ID            string `json:"id" `
	Title         string `json:"title" validate:"required"`
	Genre         string `json:"genre" validate:"required"`
	CodeISBN      string `json:"code_isbn" validate:"required"`
	IdAuthorBooks string `json:"id_author_books" validate:"required"`
}
type BookResponse struct {
	IdAuthorBooks string `json:"id_author_books"`
	ID            string `json:"id"`
	Title         string `json:"title"`
	Genre         string `json:"genre"`
	CodeISBN      string `json:"code_isbn"`
}

func ParseFromBook(src entity.Book) (dst BookResponse) {
	dst = BookResponse{
		ID:            src.ID,
		Title:         *src.Title,
		Genre:         *src.Genre,
		CodeISBN:      *src.CodeISBN,
		IdAuthorBooks: *src.IdAuthorBooks,
	}

	return
}

func ParseFromBooks(src []entity.Book) (dst []BookResponse) {
	for _, data := range src {
		dst = append(dst, ParseFromBook(data))
	}

	return
}
