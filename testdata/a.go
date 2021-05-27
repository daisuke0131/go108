package a

func a_call1(){
	return
}

func a_call2() string{
	return ""
}

func a_call3() int{
	return 1
}

func a_call4() string{
	a := ""
	return a
}

func a_call5() string{
	a := func() string{
		return ""
	}
	return a()
}

func a_call6() string{
	a := func() string{
		return func() string{
			return ""
		}()
	}
	return a()
}