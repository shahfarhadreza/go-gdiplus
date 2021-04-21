package gdiplus

type Brush struct {
	nativeBrush *GpBrush
}

type SolidBrush struct {
	Brush
}

func (b *Brush) Dispose() {
	GdipDeleteBrush(b.nativeBrush)
}

func (b *Brush) GetBrushType() (brushType BrushType) {
	GdipGetBrushType(b.nativeBrush, (*GpBrushType)(&brushType))
	return
}

func (b *Brush) Clone() *Brush {
	clone := &Brush{}
	GdipCloneBrush(b.nativeBrush, &clone.nativeBrush)
	return clone
}

func NewSolidBrush(color *Color) *SolidBrush {
	b := &SolidBrush{}
	var solidFill *GpSolidFill
	GdipCreateSolidFill(color.GetValue(), &solidFill)
	b.nativeBrush = &solidFill.GpBrush
	return b
}

func (b *SolidBrush) AsBrush() *Brush {
	return &b.Brush
}

func (b *SolidBrush) SetColor(color *Color) {
	GdipSetSolidFillColor(b.nativeBrush, color.GetValue())
}

func (b *SolidBrush) GetColor() (color Color) {
	GdipGetSolidFillColor(b.nativeBrush, &color.Argb)
	return
}
