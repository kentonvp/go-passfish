package stringutils

import "bytes"

func CenterString(str string, total_field_width int) string {

	str_len := len(str)
	spaces_to_pad := total_field_width - str_len
	var tmp_spaces float64
	var lr_spaces int

	tmp_spaces = float64(spaces_to_pad) / 2
	lr_spaces = int(tmp_spaces)

	buffer := bytes.NewBufferString("")

	spaces_remaining := total_field_width

	for i := 0; i < lr_spaces; i++ {
		buffer.WriteString(" ")
		spaces_remaining = spaces_remaining - 1
	}
	buffer.WriteString(str)
	spaces_remaining = spaces_remaining - str_len
	for i := spaces_remaining; i > 0; i-- {
		buffer.WriteString(" ")
	}

	return buffer.String()
}
