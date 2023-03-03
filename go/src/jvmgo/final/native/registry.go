package native

import "jvmgo/final/rtda"

//对本地方法进行注册

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	//类名+方法名+方法描述符 = 组成本地方法的KEY
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" {
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}
	return nil
}
