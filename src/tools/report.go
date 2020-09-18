package tools

type Data interface {
	Get()
}

func GetValues(src Data) {
	src.Get()

}
