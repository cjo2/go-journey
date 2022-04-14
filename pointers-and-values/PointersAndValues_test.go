package PointersAndValues

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompany_ChangeName(t *testing.T) {
	apple := &Company{
		Name: "Apple Computers Inc.",
	}

	assert.Equal(t, "Apple Computers Inc.", apple.Name)
	apple.ChangeName("Apple Inc.")
	assert.Equal(t, "Apple Inc.", apple.Name)
}

func TestApp_ChangePublisher(t *testing.T) {
	bungie := &Company{
		Name: "Bungie",
	}

	microsoft := &Company{
		Name: "Microsoft",
	}

	halo := &App{
		AppName:   "Halo",
		Publisher: bungie,
	}

	assert.Equal(t, "Bungie", halo.Publisher.Name)
	halo.ChangePublisher(microsoft)
	assert.Equal(t, "Microsoft", halo.Publisher.Name)
}

func TestApp_ChangeCreator(t *testing.T) {
	bungie := &Company{
		Name: "Bungie",
	}

	microsoft := &Company{
		Name: "Microsoft",
	}

	halo := &App{
		AppName:   "Halo",
		Publisher: bungie,
		Creator:   *bungie,
	}

	assert.Equal(t, "Bungie", halo.Publisher.Name)
	halo.ChangePublisher(microsoft)
	assert.Equal(t, "Microsoft", halo.Publisher.Name)
	microsoft.Name = "Microsoft Xbox"
	assert.Equal(t, "Microsoft Xbox", halo.Publisher.Name)

	// %p in Printf will print the address

	// Since Publisher stores a pointer to a value of type Company,
	// we should see the same memory address for bungie and halo.Publisher
	fmt.Printf("%p %p \n", bungie, halo.Publisher)
	// 0xc0000550a0 0xc0000550b0

	// Creator is of type Company, so we dereference bungie.
	// We should expect to see two different memory addresses
	fmt.Printf("%p %p \n", bungie, &halo.Creator)
	// 0xc000055040 0xc00007b368
}
