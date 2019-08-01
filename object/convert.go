package object

func ToStringArray(any []interface{}) []string {
	var array []string
	{
		array = make([]string, len(any), cap(any))
		for i, item := range any {
			array[i] = item.(string)
		}
	}
	return array
}

func ToIntArray(any []interface{}) []int {
	var array []int
	{
		array = make([]int, len(any), cap(any))
		for i, item := range any {
			array[i] = item.(int)
		}
	}
	return array
}