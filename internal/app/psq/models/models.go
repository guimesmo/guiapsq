package models

type Psq interface {}

type PsqImp struct {
	ID int `db:"_id" json:"id"`
	Nome string `db:"nome" json:"nome" validate:"required"`
	Endereco string `db:"email" json:"email" validate:"required,email"`
	Apresentacao string `db:apresentacao" json:"apresentacao" validate:"required"`
}