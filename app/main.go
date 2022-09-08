package main

import "go-img-poc/libs"

func main() {

	libs.CreateImageLabel("GOLANG TESTE IMAGE - 123444555", "imglabel.png")
	libs.CreateImageFogleman("../img/wolf-02.jpeg", "WOLF ID 10204080160")
	libs.CreateMergeImage("WOLF ID 10204080160")
	libs.CreateWaterMark()

}
