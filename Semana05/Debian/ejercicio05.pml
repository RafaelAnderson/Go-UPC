
#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

byte S=1
byte c=0

active[2] proctype p(){
	do
	::
		wait(S)
		c++
		assert(c<2)
		c--
		signal(S)
	od
}
