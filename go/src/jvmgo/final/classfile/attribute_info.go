package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()  //属性名索引并不是真的属性名本身，而是指向常量池中的字符串
	attrName := cp.getUtf8(attrNameIndex) //注意attrName是预定好的，只需要根据其名称进行分类即可
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	// case "MethodParameters":
	// case "RuntimeInvisibleAnnotations":
	// case "RuntimeInvisibleParameterAnnotations":
	// case "RuntimeInvisibleTypeAnnotations":
	// case "RuntimeVisibleAnnotations":
	// case "RuntimeVisibleParameterAnnotations":
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	// case "SourceDebugExtension":
	// case "StackMapTable":
	case "Synthetic":
		return _attrSynthetic
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
