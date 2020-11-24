package proof

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/bandprotocol/chain/x/oracle/types"
)

func hexToBytes(hexstr string) []byte {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return b
}

func leafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func branchHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func TestEncodeRelay(t *testing.T) {
	block := BlockRelayProof{
		MultiStoreProof: MultiStoreProof{
			AccToGovStoresMerkleHash:          hexToBytes("0D4AF3F5FFA02A56B1DEED7BC8C16732AEB8FD003C67EEF26048B314C351FAE8"),
			MainAndMintStoresMerkleHash:       hexToBytes("68BF06D17DBF1F5870D3092E1433A99FDAF6E263EFD5F8C82C533691D87592B7"),
			OracleIAVLStateHash:               hexToBytes("4F900B8B425CF85AB2A1ED2907D4830BF674703C64954847DAAEE9B81A09BA31"),
			ParamsStoresMerkleHash:            hexToBytes("2B6A7E0F44ED9C179A47A40F93D5824189A5426D6C3F77692DE28E50E20A33DD"),
			SlashingToUpgradeStoresMerkleHash: hexToBytes("F5E26E9E91F18051F41453B5A5FC82279C364351155CEA5564B9D61BC12BE58A"),
		},
		BlockHeaderMerkleParts: BlockHeaderMerkleParts{
			VersionAndChainIdHash:             hexToBytes("3561783E9C3BDF932A16580FC0C9CEFFEC4C509073FFF65A42871BFAB64408AE"),
			Height:                            3021518,
			TimeSecond:                        1605721438,
			TimeNanoSecond:                    605059026,
			LastBlockIDAndOther:               hexToBytes("21114E3076A55C6853B4730FB8678B5BF2314C1DF6DCE169ACEE9AECE893C60F"),
			NextValidatorHashAndConsensusHash: hexToBytes("EA01CD62E714B603323A21A4A7382F8AB04788C867A0C99ADE687D00E7D5FE62"),
			LastResultsHash:                   hexToBytes("AA3C7CBEFF135291E6415ECA2528FC98D263B120C67BCECD8D8CCD3253EFECC1"),
			EvidenceAndProposerHash:           hexToBytes("68D9EF5EB2AFAF2E36940299C8CDA2F43ACB015FC2D6CAFD2C577CA48F1B2C26"),
		},
		Signatures: []TMSignature{
			{
				R:                hexToBytes("D090EF654A5C8B59EB97346EB46E601A6D57DDA9174C1E535C40485FD5A8D414"),
				S:                hexToBytes("47A7B3E582103C2B28B8D3D3C8D9B7D715D852E449711C7B91BDF9EDCB3F7ADA"),
				V:                28,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510DAD8C69B02320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("26ED90EE89D4F6B5D172904DDB82A18419E6833BFC74522416800E3A7D2E3AE0"),
				S:                hexToBytes("0C2DF3F307106861195AA41643CF6DA1F5BE195B11C74B9C90329F99B1E24CA7"),
				V:                27,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A9F1DCC402320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("5CA760BDA037610B623B83CBF8E0F55FECA37A9A621EBAD17360E5725F8E3425"),
				S:                hexToBytes("3DF55E237CB9D1EAE87243CD4D3F5F1091B9FF2E2F16A2B083928F2B0F09C7BB"),
				V:                28,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A196B3C502320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("37466D99BB8ED9EA462B6A7E988189224A628FFA973A05A29BB401C4C4FC9B2D"),
				S:                hexToBytes("7FD076B0E28B7E6A35CC80AD6FA688958E9F22093CE00C22D4FE67124C633B04"),
				V:                28,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD05108BFF92C502320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("99CFB8E6848F039863D98C48AC985F4AD6A27FE0B602A8735D44654DF1B06DE2"),
				S:                hexToBytes("75C7162F5ABADA35FB48460F2196E5348334E98BA05DE1236D59DB329EC2E4AA"),
				V:                27,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD05108E8BADC302320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("780733F9013D10F88A2FA936BC2789A61146749B779FB847C0F46279CA31543E"),
				S:                hexToBytes("28EB691288B02C06D29CECC3DDEA3387CB10F69AEC9F6C0A51ED39784D6675D0"),
				V:                27,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510F0CDC5D402320F62616E642D6775616E79752D706F61"),
			},
			{
				R:                hexToBytes("2914567728CAE2ABE2B707933AE63660C7C79786C7B897CB867B68F67AAC07A8"),
				S:                hexToBytes("04F683B20F6B043A0CA79C39DD68167869262BF262CEF1B1CE8E30B4A06B0D5D"),
				V:                27,
				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A7CEC19A02320F62616E642D6775616E79752D706F61"),
			},
		},
	}
	result, err := block.encodeToEthData()
	require.Nil(t, err)
	expect := hexToBytes("0d4af3f5ffa02a56b1deed7bc8c16732aeb8fd003c67eef26048b314c351fae868bf06d17dbf1f5870d3092e1433a99fdaf6e263efd5f8c82c533691d87592b74f900b8b425cf85ab2a1ed2907d4830bf674703c64954847daaee9b81a09ba312b6a7e0f44ed9c179a47a40f93d5824189a5426d6c3f77692de28e50e20a33ddf5e26e9e91f18051f41453b5a5fc82279c364351155cea5564b9d61bc12be58a3561783e9c3bdf932a16580fc0c9ceffec4c509073fff65a42871bfab64408ae00000000000000000000000000000000000000000000000000000000002e1ace000000000000000000000000000000000000000000000000000000005fb55d5e00000000000000000000000000000000000000000000000000000000241077d221114e3076a55c6853b4730fb8678b5bf2314c1df6dce169acee9aece893c60fea01cd62e714b603323a21a4a7382f8ab04788c867a0c99ade687d00e7d5fe62aa3c7cbeff135291e6415eca2528fc98d263b120c67bcecd8d8ccd3253efecc168d9ef5eb2afaf2e36940299c8cda2f43acb015fc2d6cafd2c577ca48f1b2c2600000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000066000000000000000000000000000000000000000000000000000000000000007c00000000000000000000000000000000000000000000000000000000000000920d090ef654a5c8b59eb97346eb46e601a6d57dda9174c1e535c40485fd5a8d41447a7b3e582103c2b28b8d3d3c8d9b7d715d852e449711c7b91bdf9edcb3f7ada000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510dad8c69b02320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000026ed90ee89d4f6b5d172904ddb82a18419e6833bfc74522416800e3a7d2e3ae00c2df3f307106861195aa41643cf6da1f5be195b11c74b9c90329f99b1e24ca7000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a9f1dcc402320f62616e642d6775616e79752d706f610000000000000000000000000000000000000000000000000000005ca760bda037610b623b83cbf8e0f55feca37a9a621ebad17360e5725f8e34253df55e237cb9d1eae87243cd4d3f5f1091b9ff2e2f16a2b083928f2b0f09c7bb000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a196b3c502320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000037466d99bb8ed9ea462b6a7e988189224a628ffa973a05a29bb401c4c4fc9b2d7fd076b0e28b7e6a35cc80ad6fa688958e9f22093ce00c22d4fe67124c633b04000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd05108bff92c502320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000099cfb8e6848f039863d98c48ac985f4ad6a27fe0b602a8735d44654df1b06de275c7162f5abada35fb48460f2196e5348334e98ba05de1236d59db329ec2e4aa000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd05108e8badc302320f62616e642d6775616e79752d706f61000000000000000000000000000000000000000000000000000000780733f9013d10f88a2fa936bc2789a61146749b779fb847c0f46279ca31543e28eb691288b02c06d29cecc3ddea3387cb10f69aec9f6c0a51ed39784d6675d0000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510f0cdc5d402320f62616e642d6775616e79752d706f610000000000000000000000000000000000000000000000000000002914567728cae2abe2b707933ae63660c7c79786c7b897cb867b68f67aac07a804f683b20f6b043a0ca79c39dd68167869262bf262cef1b1ce8e30b4a06b0d5d000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a7cec19a02320f62616e642d6775616e79752d706f61000000000000000000000000000000000000000000000000000000")

	require.Equal(t, expect, result)
}

func TestEncodeVerify(t *testing.T) {
	data := OracleDataProof{
		Version: 180,
		RequestPacket: types.OracleRequestPacketData{
			ClientID:       "test",
			OracleScriptID: 1,
			Calldata:       hexToBytes("000000034254430000000000000064"),
			AskCount:       4,
			MinCount:       4,
		},
		ResponsePacket: types.OracleResponsePacketData{
			ClientID:      "test",
			RequestID:     1,
			AnsCount:      4,
			RequestTime:   1592898765,
			ResolveTime:   1592898773,
			ResolveStatus: 1,
			Result:        hexToBytes("00000000000eb9e6"),
		},
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("108873CE6589ADD98E40996E7EA55D6592CFEE6EBF140CFBE7F0C7EAF554B8BD"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  2,
				SubtreeSize:    3,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("796AEB2094F52E848A9C9ACC57C380F92010FA19956CCF9E2C6E8E7D1E490F0C"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  3,
				SubtreeSize:    5,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("423F98BE29BD72613A6815E696ECF229E2B489198E9385D3BF1F743B99CA1573"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  4,
				SubtreeSize:    8,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("386E60DB569E57EAF6F2DBFA0E87AC5DAFB849EBF1DB7C6B29C0231D644AD92E"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  5,
				SubtreeSize:    21,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("465252AF4FE160AD23C446CE54DFC3EECAD65B96BDA8C68871A244033424616A"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  6,
				SubtreeSize:    41,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("203800627751EE71822B76CF8C24F806D2528DC82247F5A49E42AF9DF04ADC65"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  7,
				SubtreeSize:    77,
				SubtreeVersion: 190,
				SiblingHash:    hexToBytes("68B4E7B02288DCE3C92C590CF09C96B64D18FF48C92D57D1F14D877C6A2720F9"),
			},
		},
	}

	result, err := data.encodeToEthData(191)
	require.Nil(t, err)
	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000000bf00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000b4000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f000000034254430000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000005ef1b4cd000000000000000000000000000000000000000000000000000000005ef1b4d50000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000eb9e6000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000b4108873ce6589add98e40996e7ea55d6592cfee6ebf140cfbe7f0c7eaf554b8bd00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000b4796aeb2094f52e848a9c9acc57c380f92010fa19956ccf9e2c6e8e7d1e490f0c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000b4423f98be29bd72613a6815e696ecf229e2b489198e9385d3bf1f743b99ca157300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000b4386e60db569e57eaf6f2dbfa0e87ac5dafb849ebf1db7c6b29c0231d644ad92e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001500000000000000000000000000000000000000000000000000000000000000b4465252af4fe160ad23c446ce54dfc3eecad65b96bda8c68871a244033424616a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002900000000000000000000000000000000000000000000000000000000000000b4203800627751ee71822b76cf8c24f806d2528dc82247f5a49e42af9df04adc6500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000004d00000000000000000000000000000000000000000000000000000000000000be68b4e7b02288dce3c92c590cf09c96b64d18ff48c92d57d1f14d877c6a2720f9")
	require.Equal(t, expect, result)
}
