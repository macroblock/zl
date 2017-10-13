package hal

type idestroyer interface {
	Destroy()
}

func destroy(o interface{}) {
	// if o == nil {
	// 	log.Warning(true, "destroy: can't get address of argument <nil>")
	// 	return
	// }
	switch v := o.(type) {
	case nil:
	case idestroyer:
		v.Destroy()
	default:
		log.Warning(true, "destroy: variable haven't got Destroy() method")
	}
	o = nil
}
