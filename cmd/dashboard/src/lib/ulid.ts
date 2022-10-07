const encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

function ulid(id: Uint8Array): string {
    var result = ''

    result += encoding[(id[0]&224)>>5]
	result += encoding[id[0]&31]
	result += encoding[(id[1]&248)>>3]
	result += encoding[((id[1]&7)<<2)|((id[2]&192)>>6)]
	result += encoding[(id[2]&62)>>1]
	result += encoding[((id[2]&1)<<4)|((id[3]&240)>>4)]
	result += encoding[((id[3]&15)<<1)|((id[4]&128)>>7)]
	result += encoding[(id[4]&124)>>2]
	result += encoding[((id[4]&3)<<3)|((id[5]&224)>>5)]
	result += encoding[id[5]&31]

	// 16 bytes of entropy
	result += encoding[(id[6]&248)>>3]
	result += encoding[((id[6]&7)<<2)|((id[7]&192)>>6)]
	result += encoding[(id[7]&62)>>1]
	result += encoding[((id[7]&1)<<4)|((id[8]&240)>>4)]
	result += encoding[((id[8]&15)<<1)|((id[9]&128)>>7)]
	result += encoding[(id[9]&124)>>2]
	result += encoding[((id[9]&3)<<3)|((id[10]&224)>>5)]
	result += encoding[id[10]&31]
	result += encoding[(id[11]&248)>>3]
	result += encoding[((id[11]&7)<<2)|((id[12]&192)>>6)]
	result += encoding[(id[12]&62)>>1]
	result += encoding[((id[12]&1)<<4)|((id[13]&240)>>4)]
	result += encoding[((id[13]&15)<<1)|((id[14]&128)>>7)]
	result += encoding[(id[14]&124)>>2]
	result += encoding[((id[14]&3)<<3)|((id[15]&224)>>5)]
	result += encoding[id[15]&31]

    return result
}

export {ulid};
