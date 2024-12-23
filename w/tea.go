package w

var Tea *tea

func init() {
	Tea = &tea{}
}

func (c *tea) String(x string) *string {
	return &x
}

func (c *tea) In64(x int64) *int64 {
	return &x
}

func (c *tea) Int(x int) *int {
	return &x
}

func (c *tea) Bool(x bool) *bool {
	return &x
}

func (c *tea) Uint64(x uint64) *uint64 {
	return &x
}

func (c *tea) Uint32(x uint32) *uint32 {
	return &x
}

func (c *tea) Float64(x float64) *float64 {
	return &x
}
func (c *tea) Float32(x float32) *float32 {
	return &x
}
