package gdiplus

type Pen struct {
	nativePen *GpPen
}

func NewPen(color *Color, width float32) *Pen {
	p := &Pen{}
	GdipCreatePen1(color.GetValue(), width, UnitWorld, &p.nativePen)
	return p
}

func NewPenFromBrush(brush *Brush, width float32) *Pen {
	p := &Pen{}
	GdipCreatePen2(brush.nativeBrush, width, UnitWorld, &p.nativePen)
	return p
}

func (p *Pen) Dispose() {
	GdipDeletePen(p.nativePen)
}

func (p *Pen) Clone() *Pen {
	clone := &Pen{}
	GdipClonePen(p.nativePen, &clone.nativePen)
	return clone
}

func (p *Pen) SetWidth(width float32) {
	GdipSetPenWidth(p.nativePen, width)
}

func (p *Pen) GetWidth() (width float32) {
	GdipGetPenWidth(p.nativePen, &width)
	return
}

func (p *Pen) SetLineCap(startCap, endCap LineCap, dashCap DashCap) {
	GdipSetPenLineCap197819(p.nativePen, GpLineCap(startCap), GpLineCap(endCap), GpDashCap(dashCap))
}

func (p *Pen) SetStartCap(startCap LineCap) {
	GdipSetPenStartCap(p.nativePen, GpLineCap(startCap))
}

func (p *Pen) SetEndCap(endCap LineCap) {
	GdipSetPenEndCap(p.nativePen, GpLineCap(endCap))
}

func (p *Pen) SetDashCap(dashCap DashCap) {
	GdipSetPenDashCap197819(p.nativePen, GpDashCap(dashCap))
}

func (p *Pen) GetStartCap() (startCap LineCap) {
	GdipGetPenStartCap(p.nativePen, (*GpLineCap)(&startCap))
	return
}

func (p *Pen) GetEndCap() (endCap LineCap) {
	GdipGetPenEndCap(p.nativePen, (*GpLineCap)(&endCap))
	return
}

func (p *Pen) GetDashCap() (dashCap DashCap) {
	GdipGetPenDashCap197819(p.nativePen, (*GpDashCap)(&dashCap))
	return
}

func (p *Pen) SetLineJoin(lineJoin LineJoin) {
	GdipSetPenLineJoin(p.nativePen, GpLineJoin(lineJoin))
}

func (p *Pen) GetLineJoin() (lineJoin LineJoin) {
	GdipGetPenLineJoin(p.nativePen, (*GpLineJoin)(&lineJoin))
	return
}

func (p *Pen) SetCustomStartCap(customCap *GpCustomLineCap) {
	GdipSetPenCustomStartCap(p.nativePen, customCap)
}

func (p *Pen) GetCustomStartCap() (customCap *GpCustomLineCap) {
	GdipGetPenCustomStartCap(p.nativePen, &customCap)
	return
}

func (p *Pen) SetCustomEndCap(customCap *GpCustomLineCap) {
	GdipSetPenCustomEndCap(p.nativePen, customCap)
}

func (p *Pen) GetCustomEndCap() (customCap *GpCustomLineCap) {
	GdipGetPenCustomEndCap(p.nativePen, &customCap)
	return
}

func (p *Pen) SetMiterLimit(miterLimit float32) {
	GdipSetPenMiterLimit(p.nativePen, miterLimit)
}

func (p *Pen) GetMiterLimit() (miterLimit float32) {
	GdipGetPenMiterLimit(p.nativePen, &miterLimit)
	return
}

func (p *Pen) SetMode(penMode PenAlignment) {
	GdipSetPenMode(p.nativePen, GpPenAlignment(penMode))
}

func (p *Pen) GetMode() (penMode PenAlignment) {
	GdipGetPenMode(p.nativePen, (*GpPenAlignment)(&penMode))
	return
}

func (p *Pen) SetTransform(matrix *GpMatrix) {
	GdipSetPenTransform(p.nativePen, matrix)
}

func (p *Pen) GetTransform(matrix *GpMatrix) {
	GdipGetPenTransform(p.nativePen, matrix)
}

func (p *Pen) ResetTransform() {
	GdipResetPenTransform(p.nativePen)
}

func (p *Pen) MultiplyTransform(matrix *GpMatrix, order MatrixOrder) {
	GdipMultiplyPenTransform(p.nativePen, matrix, GpMatrixOrder(order))
}

func (p *Pen) TranslateTransform(dx, dy float32, order MatrixOrder) {
	GdipTranslatePenTransform(p.nativePen, dx, dy, GpMatrixOrder(order))
}

func (p *Pen) ScaleTransform(sx, sy float32, order MatrixOrder) {
	GdipScalePenTransform(p.nativePen, sx, sy, GpMatrixOrder(order))
}

func (p *Pen) RotateTransform(angle float32, order MatrixOrder) {
	GdipRotatePenTransform(p.nativePen, angle, GpMatrixOrder(order))
}

func (p *Pen) SetColor(color *Color) {
	GdipSetPenColor(p.nativePen, color.GetValue())
}

func (p *Pen) GetColor() (color Color) {
	GdipGetPenColor(p.nativePen, &color.Argb)
	return
}

func (p *Pen) SetBrush(brush *Brush) {
	GdipSetPenBrushFill(p.nativePen, brush.nativeBrush)
}

func (p *Pen) GetBrush() *Brush {
	brush := &Brush{}
	GdipGetPenBrushFill(p.nativePen, &brush.nativeBrush)
	return brush
}

func (p *Pen) GetPenType() (penType PenType) {
	GdipGetPenFillType(p.nativePen, (*GpPenType)(&penType))
	return
}

func (p *Pen) GetDashStyle() (dashStyle DashStyle) {
	GdipGetPenDashStyle(p.nativePen, (*GpDashStyle)(&dashStyle))
	return
}

func (p *Pen) SetDashStyle(dashStyle DashStyle) {
	GdipSetPenDashStyle(p.nativePen, GpDashStyle(dashStyle))
}

func (p *Pen) GetDashOffset() (offset float32) {
	GdipGetPenDashOffset(p.nativePen, &offset)
	return
}

func (p *Pen) SetDashOffset(offset float32) {
	GdipSetPenDashOffset(p.nativePen, offset)
}

func (p *Pen) GetDashCount() (count int32) {
	GdipGetPenDashCount(p.nativePen, &count)
	return
}

func (p *Pen) SetDashArray(dash []float32) {
	GdipSetPenDashArray(p.nativePen, &dash[0], int32(len(dash)))
}

func (p *Pen) GetDashArray(dash *float32, count int32) {
	GdipGetPenDashArray(p.nativePen, dash, count)
}

func (p *Pen) GetCompoundCount() (count int32) {
	GdipGetPenCompoundCount(p.nativePen, &count)
	return
}

func (p *Pen) SetCompoundArray(dash []float32) {
	GdipSetPenCompoundArray(p.nativePen, &dash[0], int32(len(dash)))
}

func (p *Pen) GetCompoundArray(dash *float32, count int32) {
	GdipGetPenCompoundArray(p.nativePen, dash, count)
}
