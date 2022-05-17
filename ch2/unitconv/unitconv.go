package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Inch float64
type Meter float64

type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (i Inch) String() string  { return fmt.Sprintf("%ginch", i) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (p Pound) String() string    { return fmt.Sprintf("%gp", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) } //摄氏温度->华氏温度

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) } //华氏温度->摄氏温度

func IToM(i Inch) Meter { return Meter(i * 0.0254) } //英寸 -> 米

func MToI(m Meter) Inch { return Inch(m * 39.37) } //米 -> 英寸

func PToK(p Pound) Kilogram { return Kilogram(p * 0.45) } // 磅 -> 千克
func KToP(k Kilogram) Pound { return Pound(k * 2.2) }     //千克 -> 磅