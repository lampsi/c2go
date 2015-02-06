// +build ignore

package gc

import (
	"cmd/internal/obj"
	"strconv"
)

const (
	fmtMinus = 1 << iota
	fmtShort
	fmtSharp
	fmtByte
	fmtLong
	fmtComma
	fmtPlus
	fmtUnsigned
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (n *Node) Line() string {
	return obj.Linklinefmt(Ctxt, int(n.Lineno), false, false)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isalnum(c int) bool {
	return isalpha(c) || isdigit(c)
}

func isalpha(c int) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func isdigit(c int) bool {
	return '0' <= c && c <= '9'
}

func yyparse() int {
	panic("yyparse")
}

func plan9quote(s string) string {
	return "'" + strings.Replace(s, "'", "''", -1) + "'"
}