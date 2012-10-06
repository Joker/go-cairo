package main

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/Joker/go-cairo"
	"math"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	wid, _ :=   xproto.NewWindowId(X)
	screen :=   xproto.Setup(X).DefaultScreen(X)
		    xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
					0, 0, 240, 240, 0,
					xproto.WindowClassInputOutput, 
					screen.RootVisual,
					xproto.CwBackPixel | xproto.CwEventMask,
					[]uint32{ // values must be in the order defined by the protocol
						0xffffffff,
						xproto.EventMaskStructureNotify |
						xproto.EventMaskKeyPress |
						xproto.EventMaskKeyRelease})

	xproto.MapWindow(X, wid)
	fmt.Printf("%d %d\n", screen.AllowedDepths[0].Visuals[0].VisualId, screen.RootVisual)
	

	var (
		ux, uy float64 = 1, 1
		
		fe cairo.FontExtents
		te cairo.TextExtents
		text = "joy"
		x, y, px, dashlength float64
	)




	surface := cairo.NewSurfaceFromXCB(xproto.Drawable(wid), screen.AllowedDepths[0].Visuals[0], 240, 240)
	surface.Scale(240, 240);
	surface.SetFontSize(0.5);

	/* Drawing code goes here */
	surface.SetSourceRGB(0.0, 0.0, 0.0);
	surface.SelectFontFace("Georgia", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD);
	surface.FontExtents(&fe);

	ux, uy = surface.DeviceToUserDistance(ux, uy);
	if ux > uy {
		px = ux
	} else {
		px = uy
	}

	surface.FontExtents(&fe);
	surface.TextExtents(text, &te);
	x = 0.5 - te.X_bearing - te.Width  / 2;
	y = 0.5 - fe.Descent   + fe.Height / 2;

	/* baseline, descent, ascent, height */
	surface.SetLineWidth(4*px);
	dashlength = 9*px;
	surface.SetDash(&dashlength, 1, 0);
	surface.SetSourceRGBA(0, 0.6, 0, 0.5);
	surface.MoveTo(x + te.X_bearing, y);
	surface.RelLineTo(te.Width, 0);
	surface.MoveTo(x + te.X_bearing, y + fe.Descent);
	surface.RelLineTo(te.Width, 0);
	surface.MoveTo(x + te.X_bearing, y - fe.Ascent);
	surface.RelLineTo(te.Width, 0);
	surface.MoveTo(x + te.X_bearing, y - fe.Height);
	surface.RelLineTo(te.Width, 0);
	surface.Stroke();

	/* extents: width & height */
	surface.SetSourceRGBA(0, 0, 0.75, 0.5);
	surface.SetLineWidth(px);
	dashlength = 3*px;
	surface.SetDash(&dashlength, 1, 0);
	surface.Rectangle(x + te.X_bearing, y + te.Y_bearing, te.Width, te.Height);
	surface.Stroke();

	/* text */
	surface.MoveTo(x, y);
	surface.SetSourceRGB(0, 0, 0);
	surface.ShowText(text);

	/* bearing */
	surface.SetDash(nil, 0, 0);
	surface.SetLineWidth(2 * px);
	surface.SetSourceRGBA(0, 0, 0.75, 0.5);
	surface.MoveTo(x, y);
	surface.RelLineTo(te.X_bearing, te.Y_bearing);
	surface.Stroke();

	/* text's advance */
	surface.SetSourceRGBA(0, 0, 0.75, 0.5);
	surface.Arc(x + te.X_advance, y + te.Y_advance, 5 * px, 0, 2 * math.Pi);
	surface.Fill();

	/* reference point */
	surface.Arc(x, y, 5 * px, 0, 2 * math.Pi);
	surface.SetSourceRGBA(0.75, 0, 0, 0.5);
	surface.Fill();

	surface.Finish()
	surface.Destroy()


	
	
	for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			// fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			// fmt.Printf("Event: %s\n", ev)
		}
		if xerr != nil {
			// fmt.Printf("Error: %s\n", xerr)
		}
	}
}
