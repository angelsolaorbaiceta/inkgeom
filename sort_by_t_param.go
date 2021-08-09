package inkgeom

/*
ByTParamValue implements sort.Interface for []TParam based on the value field.
*/
type ByTParamValue []TParam

func (a ByTParamValue) Len() int {
	return len(a)
}

func (a ByTParamValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTParamValue) Less(i, j int) bool {
	return a[i].value < a[j].value
}
