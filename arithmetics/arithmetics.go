package arithmetics

func swap(s []byte, a,b int) {
    s[a],s[b] = s[b],s[a]
}

func Permute(str []byte) int{
    length := len(str)
    key := length-1
    newkey := key
    for ;key > 0 && str[key] <= str[key-1];{
        key--
    }
    key--
    if key < 0{
        return 0
    }
    newkey = length-1
    for ;newkey > key && str[newkey] <= str[key];{
        newkey--
    }
    swap(str,key,newkey)
    length--
    key++
    for ;length>key;{
        swap(str,length,key)
        key++
        length--
    }
    return 1
}

