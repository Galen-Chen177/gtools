package backend

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string, name2 string) string {
	return fmt.Sprintf("Hello %s, It's show time! and %s", name, name2)
}
