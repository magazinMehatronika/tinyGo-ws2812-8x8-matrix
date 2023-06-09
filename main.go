// WS2812 RGB LED 8x8
// on rPi PicoW
// in JetBrains GoLand 2023.1
// by magazinMehatronika 2023

package main

import (
	"fmt"
	"image/color"
	"machine"
	"strconv"
	"time"
	"tinygo.org/x/drivers/ws2812"
)

var neo = machine.GPIO8
var leds [64]color.RGBA

var colorString, numberHEX string
var rgbOUTA color.RGBA

var first2, middle2, last2 string

func showOnLed(drawingHEX []int) {
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)

	for i := 0; i <= 63; i++ {
		var s = int64(drawingHEX[i])
		numberHEX = strconv.FormatInt(int64(s), 16)
		numberHEX = fmt.Sprintf("%06s", numberHEX)

		first2 = numberHEX[0 : len(numberHEX)-4]
		middle2 = numberHEX[len(numberHEX)-4 : len(numberHEX)-2]
		last2 = numberHEX[len(numberHEX)-2:]

		R, _ := strconv.ParseUint(first2, 16, 32)
		rgbOUTA.R = uint8(R)

		G, _ := strconv.ParseUint(middle2, 16, 32)
		rgbOUTA.G = uint8(G)

		B, _ := strconv.ParseUint(last2, 16, 32)
		rgbOUTA.B = uint8(B)

		leds[i] = color.RGBA{R: rgbOUTA.R, G: rgbOUTA.G, B: rgbOUTA.B}

	}
	ws.WriteColors(leds[:])

}

func main() {

	mehatronikaHEX := []int{0x15413c, 0x103a34, 0x0c3730, 0x0c3630, 0x0d3731, 0x0d3c36, 0x12463f, 0x13433d,
		0x2d5d58, 0xbed4d3, 0x09443c, 0x054339, 0x043b33, 0x09443c, 0xc7deda, 0x2a625b,
		0x2d605a, 0xa6cac7, 0x8fc0be, 0x074a41, 0x06423a, 0x8ebab7, 0x95c4c2, 0x2a625b,
		0x2d615b, 0x4f978d, 0x70aea7, 0x62a39d, 0x66a39d, 0x6da39d, 0x478980, 0x2c645d,
		0x275e57, 0x38887f, 0x074f45, 0x89bab6, 0x8bbbb7, 0x084940, 0x3b7e75, 0x266059,
		0x0d413b, 0x024238, 0x0e554a, 0x084d43, 0x084c43, 0x06473c, 0x08443c, 0x0d433c,
		0x204f49, 0x80a9a5, 0x4c8a84, 0x679d95, 0x7aaba9, 0x72a5a1, 0x5e9890, 0x1d544c,
		0x123d37, 0x0d3730, 0x0d3832, 0x0d3a33, 0x0f3e37, 0x0e3f37, 0x0f4038, 0x133e38}

	loveHEX := []int{0x010000, 0x1a0000, 0x230000, 0x0d0000, 0x0d0000, 0x230000, 0x1a0000, 0x010000,
		0x340000, 0xd70101, 0xf70101, 0xa30101, 0xa30101, 0xf70101, 0xd70101, 0x340000,
		0x780000, 0xff0101, 0xff0101, 0xfc0101, 0xfc0101, 0xff0101, 0xff0101, 0x780000,
		0x630000, 0xff0101, 0xff0101, 0xff0101, 0xff0101, 0xff0101, 0xff0101, 0x630000,
		0x160000, 0xca0101, 0xfb0101, 0xff0101, 0xff0101, 0xfb0101, 0xca0101, 0x160000,
		0x000000, 0x370000, 0xbd0101, 0xf90101, 0xf90101, 0xbd0101, 0x370000, 0x000000,
		0x000000, 0x000000, 0x2b0000, 0xb00101, 0xaf0101, 0x2b0000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x100000, 0x100000, 0x000000, 0x000000, 0x000000}

	blueHEX := []int{0x00affe, 0x00a7fe, 0x00a7fe, 0x00a7fe, 0x00a7fe, 0x00a7fe, 0x00affe, 0x00affe,
		0x00affe, 0x00a4fb, 0x0025e7, 0x003afe, 0x0025e7, 0x002bfe, 0x002bfe, 0x009dfe,
		0x002bfe, 0x003afe, 0x0025e7, 0x000000, 0xf5efea, 0x000000, 0x009de6, 0x00affe,
		0x00affe, 0x00a4fb, 0xf59d65, 0xf59d65, 0xf59d65, 0x0033f4, 0x009de6, 0x003afe,
		0x00a4fb, 0x002fff, 0x002fff, 0x002ffa, 0x002ffa, 0x002ffa, 0x002fff, 0x00a4fb,
		0x00a4fb, 0x0025e7, 0x002ffa, 0xe8986b, 0xe8986b, 0x002ffa, 0x0025e7, 0x00a4fb,
		0x00affe, 0xf5efea, 0x002ffa, 0xffa160, 0xffa160, 0x002ffa, 0xf5efea, 0x00affe,
		0x00affe, 0x009de6, 0xd90000, 0x009de6, 0x009de6, 0xd90000, 0x009de6, 0x00affe}

	creeterHEX := []int{0x09660e, 0x09660e, 0x055f0a, 0x005202, 0x0a6a10, 0x299a2e, 0x0a690f, 0x09660e,
		0x0a670f, 0x035307, 0x065a0a, 0x0d6b12, 0x1b8220, 0x065709, 0x085c0d, 0x0e6c13,
		0x08610d, 0x000000, 0x000e00, 0x075a0c, 0x0a640f, 0x000000, 0x000200, 0x17771b,
		0x004500, 0x000000, 0x000000, 0x003f00, 0x004400, 0x000000, 0x000000, 0x0f6a14,
		0x004400, 0x003500, 0x001c00, 0x000000, 0x000000, 0x126415, 0x075c0c, 0x0a670f,
		0x09660e, 0x075a0a, 0x000000, 0x000000, 0x000000, 0x000000, 0x004400, 0x06600a,
		0x27982b, 0x0d6310, 0x001d00, 0x0a5c0e, 0x1c7e20, 0x000000, 0x0e6e13, 0x004c00,
		0x0e6e13, 0x0c6910, 0x0c640f, 0x0d6b12, 0x15781a, 0x0c640f, 0x0a670f, 0x07610c}

	heartHEX := []int{0xef9a9a, 0xef9a9a, 0xef9a9a, 0xef9a9a, 0xef9a9a, 0xef9a9a, 0xef9a9a, 0xef9a9a,
		0xef9a9a, 0x000000, 0x000000, 0xef9a9a, 0xef9a9a, 0x000000, 0x000000, 0xef9a9a,
		0x000000, 0xffffff, 0xffffff, 0x000000, 0x000000, 0xb71c1c, 0xb71c1c, 0x000000,
		0x000000, 0xb71c1c, 0xb71c1c, 0xd50000, 0xd50000, 0xd50000, 0xffffff, 0x000000,
		0x000000, 0xd50000, 0xd50000, 0xd50000, 0xb71c1c, 0xb71c1c, 0xb71c1c, 0x000000,
		0xef9a9a, 0x000000, 0xb71c1c, 0xb71c1c, 0xb71c1c, 0xd50000, 0x000000, 0xef9a9a,
		0xef9a9a, 0xef9a9a, 0x000000, 0xb71c1c, 0xb71c1c, 0x000000, 0xef9a9a, 0xef9a9a,
		0xef9a9a, 0xef9a9a, 0xef9a9a, 0x000000, 0x000000, 0xef9a9a, 0xef9a9a, 0xef9a9a}

	pikachuHEX := []int{0x000000, 0x35ed11, 0x35ed11, 0x35ed11, 0x35ed11, 0x000000, 0x000000, 0x35ed11,
		0x35ed11, 0x35ed11, 0xffff00, 0xffea00, 0x35ed11, 0x35ed11, 0x35ed11, 0xffea00,
		0xffea00, 0xffff00, 0xffff00, 0xffff00, 0xffff00, 0x35ed11, 0x35ed11, 0x35ed11,
		0xffea00, 0xffea00, 0x35ed11, 0xffff00, 0x000000, 0xffff00, 0xffff00, 0x000000,
		0xffea00, 0xffff00, 0xffff00, 0xffff00, 0xd50000, 0x35ed11, 0xffea00, 0xffea00,
		0x35ed11, 0x4e342e, 0x35ed11, 0xffff00, 0xffea00, 0xffea00, 0xffea00, 0x35ed11,
		0x35ed11, 0xffff00, 0xffea00, 0xffff00, 0xffea00, 0xffff00, 0x4e342e, 0x35ed11,
		0x35ed11, 0x35ed11, 0xffff00, 0xffea00, 0x4e342e, 0x4e342e, 0xffea00, 0x35ed11}

	for {

		showOnLed(blueHEX)
		time.Sleep(1500 * time.Millisecond)
		showOnLed(loveHEX)
		time.Sleep(1500 * time.Millisecond)
		showOnLed(mehatronikaHEX)
		time.Sleep(1500 * time.Millisecond)
		showOnLed(creeterHEX)
		time.Sleep(1500 * time.Millisecond)
		showOnLed(heartHEX)
		time.Sleep(1500 * time.Millisecond)
		showOnLed(pikachuHEX)
		time.Sleep(1500 * time.Millisecond)
	}
}
