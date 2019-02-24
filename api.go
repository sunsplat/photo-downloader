package main

import (
	"net/http"
)


func GetPhotoIds() (*http.Response, error) {

	resp, err := http.Get("https://graph.facebook.com/v3.2/10103177628059819/photos?access_token=" + token)

	return resp, err
}


func GetAlbums() (*http.Response, error) {

	firstpage := "MTAxMDA0NzE1NDM5OTM5NzkZD"
	resp, err := http.Get("https://graph.facebook.com/v3.2/10103177628059819/albums?access_token=" + token + "?pretty=0&limit=25&before=" + firstpage)

	return resp, err
}


func GetNextPage(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	return resp, err
}
