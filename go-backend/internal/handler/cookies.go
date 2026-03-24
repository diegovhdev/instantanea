package handler

import "net/http"

func SetCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name: name,
		Value: value,
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: true,
	})
}

func RemoveCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name: name,
		Value: "",
		Path: "/",
		MaxAge: -1,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})
}