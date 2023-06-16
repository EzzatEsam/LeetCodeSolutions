func validIPAddress(queryIP string) string {
    splitted := strings.Split(queryIP , ".")
    if len(splitted )== 4 {
        foundError := false
        for _ , part := range splitted {
            num , e :=strconv.Atoi(part)
            if (len(part) >3 || e != nil || num > 255 || num < 0 || ( len(part) >1 && part[0] == '0') ) {
                foundError = true;
                break;
            }

        }

        if ! foundError {
            return "IPv4"
        }
    }
    splitted = strings.Split(queryIP , ":")
    if len(splitted )== 8 {
        foundError := false
        for _ , part := range splitted {
            part := strings.ToLower(part)
			if len(part) > 4 || len(part) ==0 {
				foundError = true;
			}
            for _, ltr := range part {
				if !isHexLetter(ltr) && !isNum(ltr) {
					foundError = true
					break
				}
			}

			if foundError {
				break
			}
        }

        if ! foundError {
            return "IPv6"
        }
    }
    return "Neither"
}
func isNum(ltr rune) bool  {
	if ltr >= '0' && ltr <= '9'  {
		return true
	}	
	return false
}

func isHexLetter(ltr rune) bool  {  
	if ltr >= 'a' && ltr <= 'f'   {
		return true
	}
	return false
}