package magma

/* Implementation of GOST 28147-89 "Magma" */

type Method int8

const (
	ECB Method = iota // Electronic Codebook (Режим простой замены)
	CBC               // Cipher Block Chaining (Имитовставка)
	CFB               // Cipher Feedback (Гаммирование)
	OFB               // Output Feedback (Гаммирование с обратной связью)
)

type Magma struct {
	Method Method
}
