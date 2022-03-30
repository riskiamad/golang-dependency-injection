// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// Injectors from injector.go:

func InitializedService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMongoDB := NewDatabaseMongoDB()
	databaseRepository := NewDatabaseRepository(databasePostgreSQL, databaseMongoDB)
	return databaseRepository
}

func InitializedFooBarService() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

func InitializedHelloService() *HelloService {
	sayHelloImpl := NewSayHelloImpl()
	helloService := NewHelloService(sayHelloImpl)
	return helloService
}

func InitializedFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func InitializedFooBarUsingValue() *FooBar {
	foo := _wireFooValue
	bar := _wireBarValue
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

var (
	_wireFooValue = fooValue
	_wireBarValue = barValue
)

func InitializedReader() io.Reader {
	reader := _wireFileValue
	return reader
}

var (
	_wireFileValue = os.Stdin
)

func InitializedConfiguration() *Configuration {
	application := NewApplication()
	configuration := application.Configuration
	return configuration
}

func InitializedConnection(name string) (*Connection, func()) {
	file, cleanup := NewFile(name)
	connection, cleanup2 := NewConnection(file)
	return connection, func() {
		cleanup2()
		cleanup()
	}
}

// injector.go:

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

var HelloSet = wire.NewSet(
	NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

var fooValue = &Foo{}

var barValue = &Bar{}
