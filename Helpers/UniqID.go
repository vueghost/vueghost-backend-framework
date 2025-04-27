/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Helpers

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Unique ID.
type UniqIDParams struct {
	Prefix      string
	MoreEntropy bool
}

func UniqID(params UniqIDParams) string {
	id := ""
	entropy := int64(math.Floor(rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * 0x75bcd15))
	// Set prefix for unique id
	if params.Prefix != "" {
		id += params.Prefix
	}
	id += format(time.Now().Unix(), 8)
	// Increment global entropy value
	entropy++
	id += format(entropy, 5)
	// If we have more entropy add this
	if params.MoreEntropy == true {
		number := rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * 10
		id += strconv.FormatFloat(number, 'E', -1, 64)[0:10]
	}
	return id
}

func format(number int64, width int) string {
	hex := strconv.FormatInt(number, 16)

	if width <= len(hex) {
		// so long we split
		return hex[0:width]
	}

	for len(hex) < width {
		hex = "0" + hex
	}

	return hex
}
