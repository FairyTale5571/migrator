package models

type Config struct {
	Host     string `json:"ip"`
	Port     string `json:"port"`
	Name     string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}
