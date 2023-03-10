package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

//这里控制其只要解析过一回，就将其className存储下来，避免重符解析

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
