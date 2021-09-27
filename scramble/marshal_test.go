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
	"crypto/sha256"
	"fmt"
	"sort"
	"testing"
)

type myStrMap map[string]string

var (
	tval = myStrMap{
		"toto":    "tata",
		"prout":   "join",
		"chikada": "coulet",
	}
)

func (m myStrMap) hash() []byte {
	var clist []string

	// global hash
	gh := sha256.New()

	for key, val := range m {
		// hash svc, rpc tuples
		str := fmt.Sprintf("k:%s.v:%s", key, val)
		clist = append(clist, str)
	}

	sort.Strings(clist)
	for _, c := range clist {
		c256 := sha256.Sum256([]byte(c))
		gh.Write(c256[:])
	}

	ghash := gh.Sum(nil)
	return ghash[:]
}

func TestMarshal(t *testing.T) {
	var out myStrMap

	k := make([]byte, 32)
	_, err := rand.Read(k)
	if err != nil {
		t.Fatalf("rand error: %v\n", err)
	}

	tval64, err := Marshal(k, nil, tval)
	if err != nil {
		t.Fatalf("marshal error: %v\n", err)
	}

	tvalhash := tval.hash()

	t.Logf("marshalled: %s\n", tval64)

	err = Unmarshal(k, nil, tval64, &out)
	if err != nil {
		t.Fatalf("unmarshal error: %v\n", err)
	}

	outhash := out.hash()

	if !bytes.Equal(tvalhash, outhash) {
		t.Fatalf("orig: %x VS unmarshalled: %x\n", tvalhash, outhash)
	}
}
