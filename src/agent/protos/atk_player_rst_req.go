package protos

import . "types"
import "misc/packet"

func _atk_player_rst_req(sess *Session, reader *packet.Packet) (ret []byte, err error) {
	tbl, _ := pktread_atk_player_rst_req(reader)
	println(tbl.F_rst)
	return
}