package main

import (
	"encoding/json"
	"fmt" 
	"io/ioutil"
)
//Only use structs when you know the structure of the response
// type PhotosData struct {
// 	Data PhotoData
// 	Paging PagingPhotoData
// }

// type PhotoData struct {
// 	Created_Time string
// 	Name string
// 	Id string
// }

// type PhotoId struct {
// 	Id string
// }

// type PagingPhotoData struct {
// 	Cursors []interface{}
// 	Next string
// 	Previous string
// }

func main() {

	SetEnvironmentVariables()

	resp, err := GetPhotoIds()
	fmt.Println(resp.StatusCode)

	//TODO: Try to use structs to save the data
	// var photos PhotosData
	// var savePagingData []interface{}
	

	var photoData map[string]interface{}
	var photoIds []string

	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(body), &photoData)
	}

	data := photoData["data"].([]interface{})
	// fmt.Println(data)

	for _, value := range data {
		// Each value is an interface{} type, that is type asserted as a string
		// fmt.Println(key, value.(map[string]interface{}))
		image, _ := value.(map[string]interface{})
		id  := image["id"].(string)
		photoIds = append(photoIds, id)
		fmt.Println(photoIds)
	  }

	defer resp.Body.Close()

	fmt.Println("------------------------- Photo Data ------------------------- ")
	fmt.Println(photoData)
	fmt.Println("------------------------- Paging Data ------------------------- ")
	// fmt.Println(savePagingData)


}