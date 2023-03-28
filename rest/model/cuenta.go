package model

import "github.com/dgrijalva/jwt-go"

type Login struct {
	Correo string `json:"correo"`
	Clave  string `json:"clñve"`
}

type EstructuraToken struct {
	Correo string `json:"correo"`
	jwt.StandardClaims
}

func name() {
	//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

}
