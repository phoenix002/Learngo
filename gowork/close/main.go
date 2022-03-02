package main
func ReadWrite() bool {
	file.Open("test1")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}
func main(){

}

