package user

import "time"

type ListByBusinessDto struct {
	Name       string    `json:"name"`
	FullName   string    `json:"fullName"`
	AvatarURL  string    `json:"avatarUrl"`
	Roles      []string  `json:"roles"`
	IsVerified bool      `json:"isVerified"`
	IsCurrent  bool      `json:"isCurrent"`
	JoinAt     time.Time `json:"joinAt"`
	BirthDate  time.Time `json:"birthDate"`
	CreatedAt  time.Time `json:"createdAt"`
}
