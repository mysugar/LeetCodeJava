package pkg


type Fragment interface {
	Exec()error
}
type GetAction struct {


}

func (s GetAction) Exec()error  {
	return nil
}

var frg Fragment = new(GetAction)
var frg2 Fragment = &GetAction{}