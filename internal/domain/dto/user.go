package dto

import "github.com/google/uuid"

type UserResponses struct {
	ID        uuid.UUID `json:"id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910"`
	Name      string    `json:"name" example:"Anda Raiwin"`
	PicUrl    string    `json:"picUrl" example:"https://anda-daf-bridge.s3.amazonaws.com/users/profile-pic/48a18dd9-48c3-45a5-b4f3-e8d7a60e2910.png"`
	Email     string    `json:"email" example:"andaraiwin@gmail.com"`
	Role      string    `json:"role" example:"User"`
	UpdatedAt string    `json:"updatedAt" example:"2025-01-24T13:22:10.532645Z"`
}

type ProfileResponses struct {
	ID        uuid.UUID `json:"id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910"`
	FirstName string    `json:"firstName" example:"Anda"`
	LastName  string    `json:"lastName" example:"Raiwin"`
	Email     string    `json:"email" example:"andaraiwin@gmail.com"`
	Phone     string    `json:"phone" example:"08123456789"`
	PicUrl    string    `json:"picUrl" example:"https://anda-daf-bridge.s3.amazonaws.com/users/profile-pic/48a18dd9-48c3-45a5-b4f3-e8d7a60e2910.png"`
	Language  string    `json:"language" example:"Indonesia"`
	Role      string    `json:"role" example:"User"`
	UpdateAt  string    `json:"updatedAt" example:"2025-01-24T13:22:10.532645Z"`
}

type SiguUpRequest struct {
	Name     string `json:"name" example:"Anda Raiwin" validate:"required"`
	Email    string `json:"email" example:"andaraiwin@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"$2a$10$GEMNCwJCpl2yRm.UirLrUuIG55oc8oLCcP4HRe0uPlTizoIVRAS6K" validate:"required"`
	Role     string `json:"role" example:"User" validate:"required"`
	Provider string `json:"provider" example:"local" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"andaraiwin@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"$2a$10$GEMNCwJCpl2yRm.UirLrUuIG55oc8oLCcP4HRe0uPlTizoIVRAS6K" validate:"required"`
}
