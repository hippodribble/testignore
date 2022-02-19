package utilfn

type Zoid struct{
	A,B,c,d string
}

func(z *Zoid) Show(){
	println(z.A,z.B)
}