package common

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/kelleyk/go-cachecash/ccmsg"
	"github.com/kelleyk/go-cachecash/colocationpuzzle"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

func EncryptTicketL2(p *colocationpuzzle.Puzzle, t *ccmsg.TicketL2) ([]byte, error) {
	plaintext, err := proto.Marshal(t)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal ticket L2")
	}

	block, err := aes.NewCipher(p.Key())
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct block cipher")
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, p.IV())
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

func DecryptTicketL2(secret, ciphertext []byte) (*ccmsg.TicketL2, error) {
	p := &colocationpuzzle.Puzzle{Secret: secret}

	block, err := aes.NewCipher(p.Key())
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct block cipher")
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, p.IV())
	stream.XORKeyStream(plaintext, ciphertext)

	msg := &ccmsg.TicketL2{}
	if err := proto.Unmarshal(plaintext, msg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal plaintext message")
	}

	return msg, nil
}