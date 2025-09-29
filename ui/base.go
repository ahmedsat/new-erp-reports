package ui

// import (
// 	"image/color"
// 	"log"
// 	"os"

// 	"gioui.org/app"
// 	"gioui.org/layout"
// 	"gioui.org/op/clip"
// 	"gioui.org/op/paint"
// )

// func BaseUi(draw func(window *app.Window) error) {
// 	go func() {
// 		window := new(app.Window)
// 		err := draw(window)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Exit(0)
// 	}()
// 	app.Main()
// }

// func SetBackgroundColor(gtx layout.Context, color color.NRGBA) {
// 	defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
// 	paint.Fill(gtx.Ops, color)
// }
