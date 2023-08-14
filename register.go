package tlbin

var registerMap = map[int32]func() TLObject{}

func PutRegister(id int32, f func() TLObject) {
	registerMap[id] = f
}

func NewTLObjectByClassID(classId int32) TLObject {
	f, ok := registerMap[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) bool {
	_, ok := registerMap[classId]
	return ok
}
