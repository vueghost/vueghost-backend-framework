/*
 * Vueghost Tech,FZE (c) 2018-2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express
 * permission of Vueghost Tech, FZE.
 */

// Vueghost © 2018-2019
// Helpers Mask_test
// December, 2019

package Helpers

import (
	"fmt"
	"testing"
)

func TestMaskEmail(t *testing.T) {
	fmt.Println(MaskEmail("mohamed_ym@yahoo.com"), MaskEmail("webdevelopment@google.com"), MaskEmail("a@a.com"), MaskEmail("web.app@yahoo.com"))
}
