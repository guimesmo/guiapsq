package models

type Psq struct {
	ID           int    `db:"_id" json:"id"`
	Nome         string `db:"nome" json:"nome" validate:"required"`
	Email        string `db:"email" json:"email" validate:"required,email"`
	Endereco     string `db:"endereco" json:"endereco" validate:"required"`
	Apresentacao string `db:"apresentacao" json:"apresentacao" validate:"required"`
}
