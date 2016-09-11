package serial

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	//lg "github.com/hiromaily/golibs/log"
	"github.com/ugorji/go/codec"
)

//-----------------------------------------------------------------------------
// TODO:work in progress
//-----------------------------------------------------------------------------

// Vector is vector
type Vector struct {
	x, y, z int
}

// MarshalBinarys is to marshal binary
func (v Vector) MarshalBinarys() ([]byte, error) {
	// A simple encoding: plain text.
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinarys is to modifie the receiver so it must take a pointer receiver.
func (v *Vector) UnmarshalBinarys(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

//-----------------------------------------------------------------------------
// Gob
//-----------------------------------------------------------------------------

// ToGOB64 is binary encoder
func ToGOB64(data interface{}) (string, error) {
	// TODO: when passed slice, is it possible to handle by just interface type?
	// TODO: What kind of type is possible to convert like int, map, user-defined type?
	//data := User{Id: 10, Name: "harry"}

	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(data)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
		return "", err
	}
	//Iv+BAwEBBFVzZXIB/4IAAQIBAklkAQQAAQROYW1lAQwAAAAR/4IBFAEKaGFycnkgZGF5bwA=
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// FromGOB64 is binary decoder
func FromGOB64(str string, tData interface{}) error {
	//u := User{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
		return err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(tData)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
		return err
	}
	return nil
}

//-----------------------------------------------------------------------------
// github.com/ugorji/go/codec
//  it may be faster than go-msgpack
//-----------------------------------------------------------------------------
var mh = &codec.MsgpackHandle{RawToString: true}

// CodecEncode is encoder using github.com/ugorji/go/codec
func CodecEncode(data interface{}) []byte {
	//lg.Debugf("data: %+v", data)

	buf := &bytes.Buffer{}
	//&Buffer{buf: []byte(s)}

	enc := codec.NewEncoder(buf, mh)
	enc.Encode(data)
	//buf.Reset()

	//lg.Debugf("buf x: %x", buf.Bytes()) //基数16、10以上の数には小文字(a-f)を使用
	//lg.Debugf("buf %o: %o", buf.Bytes()) //%oは基数8だが、ここではsliceのuint8なのでエラー / uint8の配列
	//lg.Debugf("buf v: %v", buf.Bytes()) //基数8 = uint8の配列
	//lg.Debugf("buf s: %s", buf.String()) //string
	//lg.Debugf("buf s: %s", hex.EncodeToString(buf.Bytes())) //string

	return buf.Bytes()
}

// CodecDecode is decoder using github.com/ugorji/go/codec
func CodecDecode(sData string, tData interface{}) error {
	bData, err := hex.DecodeString(sData)
	if err != nil {
		return err
	}

	r := bytes.NewReader(bData)

	dec := codec.NewDecoder(r, mh)
	dec.Decode(tData)
	//r.Seek(0, os.SEEK_SET)

	return nil
}

//-----------------------------------------------------------------------------
// MsgPack
//-----------------------------------------------------------------------------
