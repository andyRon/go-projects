package main

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"os"
)

// 将图像缩放到指定大小
func resizeImage(img image.Image, width, height int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return newImg
}

// 将图像转换为灰度图像
func toGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ { // TODO
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return grayImg
}

// 计算图像的感知哈希值
func perceptualHash(img image.Image) string {
	img = resizeImage(img, 8, 8)
	grayImg := toGrayscale(img)
	sum := 0.0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sum += float64(grayImg.At(x, y).(color.Gray).Y)
		}
	}
	avg := sum / 64.0
	hash := ""
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if float64(grayImg.At(x, y).(color.Gray).Y) > avg {
				hash += "1"
			} else {
				hash += "0"
			}
		}
	}
	return hash
}

// 计算两个哈希值的汉明距离
func hammingDistance(hash1, hash2 string) int {
	distance := 0
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			distance++
		}
	}
	return distance
}

func main() {
	file1, err := os.Open("./image1.jpg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file1.Close()

	file2, err := os.Open("./image2.jpg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	//fmt.Printf("%+v", file1)

	img1, _, err := image.Decode(file1)
	if err != nil {
		fmt.Println("Error decoding image1:", err)
		return
	}

	img2, _, err := image.Decode(file2)
	if err != nil {
		fmt.Println("Error decoding image2:", err)
		return
	}

	hash1 := perceptualHash(img1)
	hash2 := perceptualHash(img2)

	distance := hammingDistance(hash1, hash2)
	fmt.Printf("Hamming Distance: %d\n", distance)
	if distance <= 10 {
		fmt.Println("Images are similar")
	} else {
		fmt.Println("Images are not similar")
	}
}
