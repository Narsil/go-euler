package arithmetics

func swap(s []byte, a, b int) {
	s[a], s[b] = s[b], s[a]
}

func Permute(str []byte) bool {
	length := len(str)
	key := length - 1
	newkey := key
	for key > 0 && str[key] <= str[key-1] {
		key--
	}
	key--
	if key < 0 {
		return false
	}
	newkey = length - 1
	for newkey > key && str[newkey] <= str[key] {
		newkey--
	}
	swap(str, key, newkey)
	length--
	key++
	for length > key {
		swap(str, length, key)
		key++
		length--
	}
	return true
}

func Rotate(str []byte) int {
	tmp := str[0]
	for i := 0; i < len(str)-1; i++ {
		str[i] = str[i+1]
	}
	str[len(str)-1] = tmp
	return 1
}
