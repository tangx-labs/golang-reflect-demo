# golang 反射

## 反射第一定律

`interface{}` 通过 `reflect.TypeOf/ValueOf` 转换为 `反射对象`

`reflect.Ptr -> reflect.Container`

`rv.CanAddr()`

`reflect.Container -> reflect.Ptr`


## 反射第二定律

`反射对象` 通过 `rv.Interface()` 转换为 `interface{} 对象`

`rv.CanInterface()`

## 反射第三定律

反射对象可以被被修改

`rv.CanSet()`

