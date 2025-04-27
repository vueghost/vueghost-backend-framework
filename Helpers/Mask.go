/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Helpers

import (
	"strings"
)

func MaskEmail(email string) interface{} {
	emailArr := strings.Split(email, "@")
	result := ""
	for index, value := range emailArr {
		if index == 1 {
			result += "@" + strings.Repeat("*", len(value)/2) + value[len(value)/2:]
		} else {

			result += value[0:len(value)/2] + strings.Repeat("*", len(value)/2)
		}
	}

	return result
}
