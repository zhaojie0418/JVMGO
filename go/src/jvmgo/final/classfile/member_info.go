package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// MemberInfo 为避免重复定义，使用同一的MemberInfo来定义字段表和方法表
type MemberInfo struct {
	//使用cp字段保存常量池指针
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// read field or method table
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	//使用make方法分配对应的成员变量slice
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

//供上面的readMembers方法调用
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// Name 用于从常量池中获取字段或者方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Descriptor 用于从常量池中获取字段或者方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (self *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (self *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return self.getUnparsedAttributeData("AnnotationDefault")
}

func (self *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
