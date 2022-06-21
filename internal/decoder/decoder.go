package decoder

import (
	"errors"
	"happybirthday_bot/internal/formatting"
)

type Decoder struct{}

func New() *Decoder {
	return &Decoder{}
}

var hardDecoder = map[string]string{
	"YS4N3VPNVE": "УДАЧИ",
	"HECKGSFQRB": "РАДОСТИ",
	"TFMARFRQ9B": "СЧАСТЬЯ",
	"MAZWQZF2D6": "МИНЕТОВ",
	"MKF7NG8F79": "ЗДОРОВЬЯ",
	"BLMBYBU9VG": "КАРБОНАРЫ",
	"P5QXJC7Q38": "ЯСНЫХ ДНЕЙ",
	"3L5MN2PPMD": "МНОГО ДЕНЯК",
	"2RJB8ZYD9C": "МЕНЬШЕ БЕРЁЗ",
	"56FCHZJE4F": "ВСЕЙ КОЛЫ МИРА",
	"UXBSX33HZ8": "СЛАДКИХ ГРАНАТОВ",
	"HFXRAHPVQL": "ИСПОЛНЕНИЯ ЖЕЛАНИЙ",
	"A7KFFNW5LY": "ПРОЦВЕТАНИЯ ПИРАТСТВУ",
	"759JJWUBAG": "ХОРОШЕЙ ИГРЫ ;)",
}

func (d *Decoder) Decode(str string) (string, error) {
	if ans, ok := hardDecoder[formatting.Format(str)]; !ok {
		return ans, nil
	}
	return "", errors.New("dacode error")
}
