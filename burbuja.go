package main

func Burbuja(s []int64)  {
    intercambios := true
    numPasada := len(s)-1
    for numPasada > 0 && intercambios{
       intercambios = false
       for i, v := range s {
            if i < len(s)-1{
                if v > s[i+1]{
                    intercambios = true
                    s[i] = s[i+1]
                    s[i+1] = v
                }
            }
        }
	   numPasada -= 1
	}
}

func main()  {
	
}