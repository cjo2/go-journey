package PointersAndValues

type App struct {
	AppName   string   // Resolves to a value of a string
	Publisher *Company // Resolves to a pointer to value of type Company
	Creator   Company
}

type Company struct {
	Name string
}

type Organization interface {
	ChangeName() string
}

// ChangeName has a receiver argument of Company
func (c *Company) ChangeName(newName string) {
	c.Name = newName
}

// ChangePublisher requires an argument of a pointer to a value of type Company
func (a *App) ChangePublisher(newPublisher *Company) {
	a.Publisher = newPublisher
}

func (a *App) ChangeCreator(company Company) {
	a.Creator = company
}
