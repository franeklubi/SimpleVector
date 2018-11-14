package SimpleVector

import (
    "math"
)

type SVector struct {
    X, Y, Z float64
}

func (sv SVector) Magnitude() (float64) {
    return math.Sqrt(sv.X*sv.X+sv.Y*sv.Y+sv.Z*sv.Z)
}

func (sv SVector) Subtract(sv2 SVector) (SVector) {
    return SVector{sv.X-sv2.X, sv.Y-sv2.Y, sv.Z-sv2.Z}
}

func (sv SVector) SubtractN(n float64) (SVector) {
    return SVector{sv.X-n, sv.Y-n, sv.Z-n}
}

func (sv SVector) Add(sv2 SVector) (SVector) {
    return SVector{sv.X+sv2.X, sv.Y+sv2.Y, sv.Z+sv2.Z}
}

func (sv SVector) AddN(n float64) (SVector) {
    return SVector{sv.X+n, sv.Y+n, sv.Z+n}
}

func (sv SVector) Multiply(sv2 SVector) (SVector) {
    return SVector{sv.X*sv2.X, sv.Y*sv2.Y, sv.Z*sv2.Z}
}

func (sv SVector) MultiplyN(n float64) (SVector) {
    return SVector{sv.X*n, sv.Y*n, sv.Z*n}
}

func (sv SVector) Divide(sv2 SVector) (SVector) {
    if ( sv2.X == 0 || sv2.Y == 0 || sv2.Z == 0 ) {
        panic("Can't divide by 0!")
    }
    return SVector{sv.X/sv2.X, sv.Y/sv2.Y, sv.Z/sv2.Z}
}

func (sv SVector) DivideN(n float64) (SVector) {
    if ( n == 0 ) {
        panic("Can't divide by 0!")
    }
    return SVector{sv.X/n, sv.Y/n, sv.Z/n}
}

func (sv SVector) Normalize() (SVector) {
    mag := sv.Magnitude()
    if (mag != 0 && mag != 1) {
        return sv.DivideN(mag)
    }
    return sv
}

func (sv SVector) Distance(sv2 SVector) (float64) {
    return sv.Subtract(sv2).Magnitude()
}

func (sv SVector) Limit(n float64) (SVector) {
    if ( sv.Magnitude() > n) {
        return sv.Normalize().MultiplyN(n)
    } else {
        return sv
    }
}
