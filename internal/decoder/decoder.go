package decoder

import (
	"errors"
	"happybirthday_bot/internal/formatting"
)

type config interface {
	GetWishes() []string
}

// валидные коды, к которым мапятся пожелания
var codes = []string{
	"YS4N3VPNVE", "HECKGSFQRB", "TFMARFRQ9B", "MAZWQZF2D6", "MKF7NG8F79", "BLMBYBU9VG", "P5QXJC7Q38", "3L5MN2PPMD",
	"2RJB8ZYD9C", "56FCHZJE4F", "UXBSX33HZ8", "HFXRAHPVQL", "A7KFFNW5LY", "759JJWUBAG",
}

type Decoder struct {
	hardDecoder map[string]string
}

func New(conf config) *Decoder {
	wishes := conf.GetWishes()
	decoder := make(map[string]string, len(wishes))

	for i := 0; i < len(wishes); i++ {
		decoder[codes[i]] = wishes[i]
	}

	return &Decoder{hardDecoder: decoder}
}

func (d *Decoder) Decode(str string) (string, error) {
	if ans, ok := d.hardDecoder[formatting.Format(str)]; !ok {
		return ans, nil
	}
	return "", errors.New("dacode error")
}
