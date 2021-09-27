//go:build go1.17
// +build go1.17

/*
 *
 * BSD 3-Clause License
 *
 * Copyright (c) 2020 Eric Augé <eau [plus] pkisauce [a.t.] unix4fun [D.O.T] net> as Ermites.IO
 * All rights reserved.
 *
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 * list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation
 * and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */
package scramble

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"

	xcha "golang.org/x/crypto/chacha20poly1305"
)

var (
	sb64      = base64.RawURLEncoding
	nonceSize = 32
	keySize   = 32
)

func Marshal(key, ad []byte, d interface{}) (str string, err error) {
	var blob bytes.Buffer

	if len(key) != keySize {
		err = ErrScrambleInvalidInput
		return
	}

	// we should gob, encrypt and b64 to store
	enc := gob.NewEncoder(&blob)
	err = enc.Encode(d)
	if err != nil {
		return
	}

	// encryption..
	// TODO: sanity checks.
	nonce := make([]byte, nonceSize)

	_, err = rand.Read(nonce)
	if err != nil {
		err = ErrScrambleRand
		return
	}

	aead, err := xcha.NewX(key)
	if err != nil {
		err = ErrScrambleCrypto
		return
	}

	ct := aead.Seal(nonce, nonce[0:24], blob.Bytes(), ad)
	// now armor..
	str = sb64.EncodeToString(ct)
	return
}

func Unmarshal(key, ad []byte, blob64 string, o interface{}) (err error) {
	if len(key) != keySize || len(blob64) < nonceSize+16 {
		err = ErrScrambleInvalidInput
		return
	}
	//h = make(Hosts)
	ct, err := sb64.DecodeString(blob64)
	if err != nil {
		return
	}

	// TODO decrypt here.
	nonce := ct[0:32]
	aead, err := xcha.NewX(key)
	if err != nil {
		err = ErrScrambleCrypto
		return
	}

	blob, err := aead.Open(nil, nonce[0:24], ct[32:], ad)
	if err != nil {
		err = ErrScrambleCrypto
		return
	}

	bh := bytes.NewBuffer(blob)
	dec := gob.NewDecoder(bh)
	err = dec.Decode(o)
	return
}
