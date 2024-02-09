// Code generated by "pipeline"; DO NOT EDIT.
package nahida

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.40304800868034363,
		0.4332770109176636,
		0.46350499987602234,
		0.5038099884986877,
		0.5340390205383301,
		0.5642669796943665,
		0.6045719981193542,
		0.644877016544342,
		0.6851819753646851,
		0.7254859805107117,
		0.7657909989356995,
		0.8060960173606873,
		0.8564770221710205,
		0.9068580269813538,
		0.9572389721870422,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.3697440028190613,
		0.397475004196167,
		0.4252060055732727,
		0.4621799886226654,
		0.4899109899997711,
		0.5176420211791992,
		0.5546159744262695,
		0.5915899872779846,
		0.6285650134086609,
		0.665539026260376,
		0.7025139927864075,
		0.7394880056381226,
		0.7857059836387634,
		0.8319240212440491,
		0.8781419992446899,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.45874398946762085,
		0.493149995803833,
		0.5275560021400452,
		0.5734300017356873,
		0.6078360080718994,
		0.6422420144081116,
		0.6881160140037537,
		0.7339900135993958,
		0.779865026473999,
		0.8257390260696411,
		0.8716139793395996,
		0.9174879789352417,
		0.9748309850692749,
		1.032173991203308,
		1.0895169973373413,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.5840640068054199,
		0.6278690099716187,
		0.6716740131378174,
		0.7300800085067749,
		0.7738850116729736,
		0.8176900148391724,
		0.8760960102081299,
		0.9345020055770874,
		0.9929090142250061,
		1.0513149499893188,
		1.1097220182418823,
		1.1681280136108398,
		1.2411359548568726,
		1.3141440153121948,
		1.3871519565582275,
	}
	// attack: charge = [4]
	charge = []float64{
		1.3200000524520874,
		1.4190000295639038,
		1.5180000066757202,
		1.649999976158142,
		1.7489999532699585,
		1.8480000495910645,
		1.9800000190734863,
		2.111999988555908,
		2.24399995803833,
		2.375999927520752,
		2.507999897003174,
		2.640000104904175,
		2.805000066757202,
		2.9700000286102295,
		3.134999990463257,
	}
	// skill: skillHold = [1]
	skillHold = []float64{
		1.3040000200271606,
		1.4018000364303589,
		1.4996000528335571,
		1.6299999952316284,
		1.7278000116348267,
		1.825600028038025,
		1.9559999704360962,
		2.086400032043457,
		2.2167999744415283,
		2.3471999168395996,
		2.47760009765625,
		2.6080000400543213,
		2.7709999084472656,
		2.934000015258789,
		3.0969998836517334,
	}
	// skill: skillPress = [0]
	skillPress = []float64{
		0.984000027179718,
		1.057800054550171,
		1.131600022315979,
		1.2300000190734863,
		1.3037999868392944,
		1.3775999546051025,
		1.4759999513626099,
		1.5743999481201172,
		1.6727999448776245,
		1.7711999416351318,
		1.8696000576019287,
		1.968000054359436,
		2.0910000801086426,
		2.2139999866485596,
		2.3369998931884766,
	}
	// skill: triKarmaAtk = [2]
	triKarmaAtk = []float64{
		1.031999945640564,
		1.1094000339508057,
		1.1868000030517578,
		1.2899999618530273,
		1.367400050163269,
		1.4448000192642212,
		1.5479999780654907,
		1.6512000560760498,
		1.7544000148773193,
		1.8575999736785889,
		1.960800051689148,
		2.063999891281128,
		2.193000078201294,
		2.322000026702881,
		2.4509999752044678,
	}
	// skill: triKarmaEM = [3]
	triKarmaEM = []float64{
		2.063999891281128,
		2.2188000679016113,
		2.3736000061035156,
		2.5799999237060547,
		2.734800100326538,
		2.8896000385284424,
		3.0959999561309814,
		3.3024001121520996,
		3.5088000297546387,
		3.7151999473571777,
		3.921600103378296,
		4.127999782562256,
		4.386000156402588,
		4.644000053405762,
		4.9019999504089355,
	}
	// burst: burstTriKarmaCDReduction = [2 3]
	burstTriKarmaCDReduction = [][]float64{
		{
			0.24799999594688416,
			0.26660001277923584,
			0.28519999980926514,
			0.3100000023841858,
			0.3285999894142151,
			0.3472000062465668,
			0.3720000088214874,
			0.3968000113964081,
			0.42160001397132874,
			0.446399986743927,
			0.47119998931884766,
			0.4959999918937683,
			0.5270000100135803,
			0.5580000281333923,
			0.5889999866485596,
		},
		{
			0.3720000088214874,
			0.39989998936653137,
			0.4277999997138977,
			0.4650000035762787,
			0.492900013923645,
			0.520799994468689,
			0.5580000281333923,
			0.5952000021934509,
			0.6323999762535095,
			0.6696000099182129,
			0.7067999839782715,
			0.7440000176429749,
			0.7904999852180481,
			0.8370000123977661,
			0.8834999799728394,
		},
	}
	// burst: burstTriKarmaDmgBonus = [0 1]
	burstTriKarmaDmgBonus = [][]float64{
		{
			0.14880000054836273,
			0.15996000170707703,
			0.17112000286579132,
			0.1860000044107437,
			0.197160005569458,
			0.2083200067281723,
			0.2231999933719635,
			0.2380799949169159,
			0.2529599964618683,
			0.2678399980068207,
			0.28271999955177307,
			0.29760000109672546,
			0.31619998812675476,
			0.33480000495910645,
			0.35339999198913574,
		},
		{
			0.2231999933719635,
			0.23994000256061554,
			0.2566800117492676,
			0.27900001406669617,
			0.295740008354187,
			0.31248000264167786,
			0.33480000495910645,
			0.35712000727653503,
			0.3794400095939636,
			0.4017600119113922,
			0.4240800142288208,
			0.446399986743927,
			0.47429999709129333,
			0.5022000074386597,
			0.5300999879837036,
		},
	}
	// burst: burstTriKarmaDurationExtend = [4 5]
	burstTriKarmaDurationExtend = [][]float64{
		{
			3.3440001010894775,
			3.5947999954223633,
			3.845599889755249,
			4.179999828338623,
			4.430799961090088,
			4.681600093841553,
			5.015999794006348,
			5.350399971008301,
			5.684800148010254,
			6.019199848175049,
			6.353600025177002,
			6.688000202178955,
			7.105999946594238,
			7.52400016784668,
			7.941999912261963,
		},
		{
			5.015999794006348,
			5.392199993133545,
			5.768400192260742,
			6.269999980926514,
			6.646200180053711,
			7.02239990234375,
			7.52400016784668,
			8.02560043334961,
			8.527199745178223,
			9.028800010681152,
			9.530400276184082,
			10.031999588012695,
			10.659000396728516,
			11.28600025177002,
			11.913000106811523,
		},
	}
)
