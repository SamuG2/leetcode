package main


func onlyOnes(s string) bool{
	for _, c:= range s{
		if c != '1'{
			return false
		}
	}
	return true
}

func smallestGoodBase(n string) string {
    for i:=2;i<n
}