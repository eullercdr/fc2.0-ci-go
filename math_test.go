package main

import "testing"

func testSoma(t *testing.T) {
	  total:= Soma(1,2)
    if total != 3 {
				t.Errorf("Soma(1,2) deveria ser 3, mas foi %d", total)
		}		
	}