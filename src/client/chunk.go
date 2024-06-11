package internal

type Chunk struct {
	Start    int64
	End      int64
	FileName string
	PartNum  int
	Content  byte
}

func NewChunk(Start int64, End int64, FileName string) *Chunk {
	return &Chunk{
		Start:    Start,
		End:      End,
		FileName: FileName,
	}
}

func (ck *Chunk) SetPartNumber(PartNum int) *Chunk {
	ck.PartNum = PartNum
	return ck
}

func (ck *Chunk) ChunkStart(Start int64) *Chunk {
	ck.Start = Start
	return ck
}

func (ck *Chunk) ChunkEnd(End int64) *Chunk {
	ck.End = End
	return ck
}

func (ck *Chunk) ChunkContent(Content byte) *Chunk {
	ck.Content = Content
	return ck
}
