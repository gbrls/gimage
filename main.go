package main

import "github.com/gabrielSchneider100/gimage/cmd"

/*

TODO:

* - Add CLI written in cobra
* - Write Image.Pix in file
* - Scale down images to problably something like 80px x 80px

*/

func main() {
	// err := img.GetImagesURLs("People")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cmd.Execute()
}
