# factory design pattern

## simple factory

由于 Go 没有构造函数，一般我们采用 NewXXX 的方式创建对象，当返回接口的时候，其实就是简单工厂模式

## factory method

当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作，推荐使用工厂方法模式，
将复杂的创建逻辑拆分到多个工厂类中，每个工厂类都不会过于复杂

### examples

#### docker/distribution

* [distribution/registry/storage/driver/storagedriver.go](https://github.com/docker/distribution/tree/master/registry/storage/driver/storagedriver.go)
* [distribution/registry/storage/driver/factory/factory.go](https://github.com/docker/distribution/blob/master/registry/storage/driver/factory/factory.go)
* [distribution/registry/storage/driver/filesystem/driver.go](https://github.com/docker/distribution/blob/master/registry/storage/driver/filesystem/driver.go)

## abstract factory

不太常用

# reference list

* [Factory patterns in Go](https://www.sohamkamani.com/golang/2018-06-20-golang-factory-patterns/)
* [golang factory pattern](https://wangzhezhe.github.io/2019/06/16/golang-factory-pattern/)
* [Factory Design Pattern in Go](https://golangbyexample.com/golang-factory-design-pattern/)
* [Go — Factory Pattern](https://medium.com/@haluan/go-factory-design-pattern-e5301e0f3283)
* [golang factory pattern and examples](https://wangzhezhe.github.io/2019/06/16/golang-factory-pattern/)
