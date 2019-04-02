/*
 * Copyright 2019 The go-vite Authors
 * This file is part of the go-vite library.
 *
 * The go-vite library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The go-vite library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with the go-vite library. If not, see <http://www.gnu.org/licenses/>.
 */

package vnode

import (
	"bytes"
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"testing"
)

func compare(n, n2 *Node, rest bool) error {
	if n.ID != n2.ID {
		return fmt.Errorf("different ID %s %s", n.ID, n2.ID)
	}

	if !bytes.Equal(n.Host, n2.Host) {
		return fmt.Errorf("different Host %v %v", n.Host, n2.Host)
	}

	if n.Port != n2.Port {
		return fmt.Errorf("different Port %d %d", n.Port, n2.Port)
	}

	if n.Net != n2.Net {
		return fmt.Errorf("different Net %d %d", n.Net, n2.Net)
	}

	if rest && !bytes.Equal(n.Ext, n2.Ext) {
		return fmt.Errorf("different Ext %v %v", n.Ext, n2.Ext)
	}

	return nil
}

func TestNodeID_IsZero(t *testing.T) {
	if !ZERO.IsZero() {
		t.Fail()
	}
}

func TestNode_Serialize(t *testing.T) {
	var conds = [4][2]bool{
		{true, true},
		{true, false},
		{false, false},
		{false, true},
	}

	for i := 0; i < 4; i++ {
		n := MockNode(conds[i][0], conds[i][1])
		data, err := n.Serialize()
		if err != nil {
			t.Error(err)
		}

		n2 := new(Node)
		err = n2.Deserialize(data)
		if err != nil {
			t.Error(err)
		}

		if err = compare(n, n2, true); err != nil {
			t.Error(err)
		}
	}
}

func TestParseNode_Net(t *testing.T) {
	n, err := ParseNode("vite.org/2")

	if err != nil {
		t.Error(err)
	}

	if n.Net != 2 {
		t.Fail()
	}

	n, err = ParseNode("vite.org")
	if err != nil {
		t.Error(err)
	}
	if n.Net != 0 {
		t.Fail()
	}
}

func TestDistance(t *testing.T) {
	for i := uint(0); i < IDBits; i++ {
		var a, b NodeID

		byt := i / 8
		bit := i % 8

		// set i bit to 1
		a[byt] = a[byt] | (1 << (7 - bit))

		fmt.Printf("%b\n", a)
		fmt.Printf("%b\n", b)

		fmt.Println()

		if Distance(a, b) != IDBits-i {
			fmt.Printf("%b\n", a)
			fmt.Printf("%b\n", b)

			t.Errorf("distance should equal %d, but get %d", i, Distance(a, b))
		}
	}
}

func TestRandFromDistance(t *testing.T) {
	var id, id2 NodeID
	var d uint

	for i := 0; i < 1000000; i++ {
		crand.Read(id[:])
		d = uint(rand.Intn(IDBits))
		id2 = RandFromDistance(id, d)

		if Distance(id, id2) != d {
			fmt.Printf("%b\n", id)
			fmt.Printf("%b\n", id2)
			fmt.Println(Distance(id, id2), d)
			t.Fatal()
		}
	}
}
