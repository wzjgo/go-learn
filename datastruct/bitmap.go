package datastruct

type Bitmap struct {
	words  []uint64
	length uint64
}

func NewBitmap(size uint64) *Bitmap {
	return &Bitmap{
		words:  make([]uint64, size),
		length: size,
	}
}

func (bitmap *Bitmap) Has(num int) bool {
	word, bit := num/64, uint(num%64)
	return word < len(bitmap.words) && (bitmap.words[word]&(1<<bit)) != 0
}

func (bitmap *Bitmap) Add(num int) {
	word, bit := num/64, uint(num%64)
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

func (bitmap *Bitmap) Len() uint64 {
	return bitmap.length
}
