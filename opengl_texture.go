package main

import (
	"github.com/autovelop/playthos"
	_ "github.com/autovelop/playthos/glfw/keyboard"
	_ "github.com/autovelop/playthos/opengl"
	_ "github.com/autovelop/playthos/platforms/web"
	_ "github.com/autovelop/playthos/webgl"
	// _ "github.com/autovelop/playthos/platforms/linux"
	"github.com/autovelop/playthos/animation"
	"github.com/autovelop/playthos/keyboard"
	_ "github.com/autovelop/playthos/platforms/windows"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

func main() {
	eng := engine.New("TestOpenGLTexture", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	r := std.Vector3{0, 0, 0}
	NewGameObject(eng.NewEntity()).
		NewTransform(&std.Vector3{0, 0, 5}, &r, &std.Vector3{5, 5, 1}).
		NewMaterial("background.png", &std.Color{1, 1, 1, 1}).
		NewAnimation(1000, &r, []std.Animatable{&std.Vector3{0, 0, 0}, &std.Vector3{0, 0, 360}})

	kb := eng.Listener(&keyboard.Keyboard{})

	kb.On(keyboard.KeyEscape, func(action ...int) {
		switch action[0] {
		case keyboard.ActionRelease:
			eng.Stop()
		}
	})

	eng.Start()
}

type GameObject struct {
	entity *engine.Entity
}

func NewGameObject(e *engine.Entity) *GameObject {
	return &GameObject{e}
}

func (g *GameObject) NewTransform(p *std.Vector3, r *std.Vector3, s *std.Vector3) *GameObject {
	t := std.NewTransform()
	t.Set(p, r, s)
	g.entity.AddComponent(t)
	return g
}

func (g *GameObject) NewMaterial(f string, c *std.Color) *GameObject {
	m := render.NewMaterial()
	m.Set(c)
	i := render.NewImage()
	i.LoadImage(f)
	t := render.NewTexture(i)
	m.SetTexture(t)
	g.entity.AddComponent(m)

	q := render.NewMesh()
	q.Set(std.QuadMesh)
	g.entity.AddComponent(q)
	return g
}

func (g *GameObject) NewAnimation(d float64, v std.Animatable, kfs []std.Animatable) *GameObject {
	a := animation.NewClip(10, d, v)
	for i, kf := range kfs {
		a.AddKeyFrame(float64(i)*d, 0, kf)
	}
	g.entity.AddComponent(a)
	return g
}
