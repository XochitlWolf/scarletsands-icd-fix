// Code generated by "pipeline"; DO NOT EDIT.
package tartaglia

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
		attack_5,
		attack_6,
	}
	attackE = [][][]float64{
		{attackE_1},
		{attackE_2},
		{attackE_3},
		{attackE_4},
		{attackE_5},
		attackE_6,
	}
)

var (
	// attack: aim = [6]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.9282000064849854,
		0.9894000291824341,
		1.0506000518798828,
		1.111799955368042,
		1.1729999780654907,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.41280001401901245,
		0.446399986743927,
		0.47999998927116394,
		0.527999997138977,
		0.5616000294685364,
		0.6000000238418579,
		0.6528000235557556,
		0.7056000232696533,
		0.758400022983551,
		0.8159999847412109,
		0.8736000061035156,
		0.9312000274658203,
		0.9887999892234802,
		1.0463999509811401,
		1.1039999723434448,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.46268001198768616,
		0.5003399848937988,
		0.5379999876022339,
		0.5917999744415283,
		0.6294599771499634,
		0.6725000143051147,
		0.7316799759864807,
		0.7908599972724915,
		0.8500400185585022,
		0.9146000146865845,
		0.9791600108146667,
		1.043720006942749,
		1.1082799434661865,
		1.1728399991989136,
		1.2374000549316406,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.5538399815559387,
		0.5989199876785278,
		0.6439999938011169,
		0.7084000110626221,
		0.7534800171852112,
		0.8050000071525574,
		0.8758400082588196,
		0.9466800093650818,
		1.0175199508666992,
		1.0947999954223633,
		1.1720800399780273,
		1.2493599653244019,
		1.326640009880066,
		1.40392005443573,
		1.4811999797821045,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.5701799988746643,
		0.6165900230407715,
		0.6629999876022339,
		0.7293000221252441,
		0.7757099866867065,
		0.8287500143051147,
		0.9016799926757812,
		0.9746099710464478,
		1.0475399494171143,
		1.1270999908447266,
		1.2066600322723389,
		1.2862199544906616,
		1.365779995918274,
		1.4453400373458862,
		1.524899959564209,
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.6088799834251404,
		0.6584399938583374,
		0.7080000042915344,
		0.7788000106811523,
		0.8283600211143494,
		0.8849999904632568,
		0.96288001537323,
		1.0407600402832031,
		1.1186399459838867,
		1.2036000490188599,
		1.2885600328445435,
		1.373520016670227,
		1.4584800004959106,
		1.5434399843215942,
		1.6283999681472778,
	}
	// attack: attack_6 = [5]
	attack_6 = []float64{
		0.7275599837303162,
		0.7867799997329712,
		0.8460000157356262,
		0.9305999875068665,
		0.9898200035095215,
		1.0575000047683716,
		1.1505600214004517,
		1.2436200380325317,
		1.3366800546646118,
		1.4381999969482422,
		1.539720058441162,
		1.6412400007247925,
		1.7427599430084229,
		1.8442800045013428,
		1.9457999467849731,
	}
	// attack: fullaim = [7]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.3559999465942383,
		2.4800000190734863,
		2.634999990463257,
		2.7899999618530273,
		2.944999933242798,
	}
	// attack: rtBurst = [9]
	rtBurst = []float64{
		0.6200000047683716,
		0.6664999723434448,
		0.7129999995231628,
		0.7749999761581421,
		0.8215000033378601,
		0.8679999709129333,
		0.9300000071525574,
		0.9919999837875366,
		1.0540000200271606,
		1.1160000562667847,
		1.1779999732971191,
		1.2400000095367432,
		1.3174999952316284,
		1.3949999809265137,
		1.472499966621399,
	}
	// attack: rtFlash = [8]
	rtFlash = []float64{
		0.12399999797344208,
		0.13330000638961792,
		0.14259999990463257,
		0.1550000011920929,
		0.16429999470710754,
		0.1736000031232834,
		0.1860000044107437,
		0.19840000569820404,
		0.21080000698566437,
		0.2231999933719635,
		0.23559999465942383,
		0.24799999594688416,
		0.26350000500679016,
		0.27900001406669617,
		0.2944999933242798,
	}
	// skill: attackE_1 = [1]
	attackE_1 = []float64{
		0.38872000575065613,
		0.4203599989414215,
		0.4519999921321869,
		0.49720001220703125,
		0.5288400053977966,
		0.5649999976158142,
		0.6147199869155884,
		0.6644399762153625,
		0.7141600251197815,
		0.7684000134468079,
		0.8226400017738342,
		0.8768799901008606,
		0.931119978427887,
		0.9853600263595581,
		1.0396000146865845,
	}
	// skill: attackE_2 = [2]
	attackE_2 = []float64{
		0.41624000668525696,
		0.4501200020313263,
		0.48399999737739563,
		0.5324000120162964,
		0.5662800073623657,
		0.6050000190734863,
		0.658240020275116,
		0.7114800214767456,
		0.7647200226783752,
		0.8227999806404114,
		0.8808799982070923,
		0.9389600157737732,
		0.9970399737358093,
		1.0551199913024902,
		1.1131999492645264,
	}
	// skill: attackE_3 = [3]
	attackE_3 = []float64{
		0.5633000135421753,
		0.6091499924659729,
		0.6549999713897705,
		0.7204999923706055,
		0.7663499712944031,
		0.8187500238418579,
		0.8907999992370605,
		0.9628499746322632,
		1.0348999500274658,
		1.1134999990463257,
		1.1921000480651855,
		1.2706999778747559,
		1.3493000268936157,
		1.427899956703186,
		1.506500005722046,
	}
	// skill: attackE_4 = [4]
	attackE_4 = []float64{
		0.5994200110435486,
		0.6482099890708923,
		0.6970000267028809,
		0.766700029373169,
		0.8154900074005127,
		0.8712499737739563,
		0.947920024394989,
		1.024590015411377,
		1.1012599468231201,
		1.1849000453948975,
		1.2685400247573853,
		1.352180004119873,
		1.4358199834823608,
		1.5194599628448486,
		1.6030999422073364,
	}
	// skill: attackE_5 = [5]
	attackE_5 = []float64{
		0.5529800057411194,
		0.5979899764060974,
		0.6430000066757202,
		0.7073000073432922,
		0.7523099780082703,
		0.8037499785423279,
		0.8744800090789795,
		0.9452099800109863,
		1.0159399509429932,
		1.0930999517440796,
		1.170259952545166,
		1.2474199533462524,
		1.3245799541473389,
		1.4017399549484253,
		1.4788999557495117,
	}
	// skill: attackE_6 = [6 7]
	attackE_6 = [][]float64{
		{
			0.3543199896812439,
			0.3831599950790405,
			0.41200000047683716,
			0.45320001244544983,
			0.4820399880409241,
			0.5149999856948853,
			0.5603200197219849,
			0.6056399941444397,
			0.6509600281715393,
			0.7003999948501587,
			0.7498400211334229,
			0.7992799878120422,
			0.8487200140953064,
			0.8981599807739258,
			0.9476000070571899,
		},
		{
			0.3766799867153168,
			0.40733999013900757,
			0.43799999356269836,
			0.48179998993873596,
			0.5124599933624268,
			0.5475000143051147,
			0.5956799983978271,
			0.6438599824905396,
			0.6920400261878967,
			0.7445999979972839,
			0.7971600294113159,
			0.8497200012207031,
			0.9022799730300903,
			0.9548400044441223,
			1.0074000358581543,
		},
	}
	// skill: eCharge = [8 9]
	eCharge = [][]float64{
		{
			0.6019999980926514,
			0.6510000228881836,
			0.699999988079071,
			0.7699999809265137,
			0.8190000057220459,
			0.875,
			0.9520000219345093,
			1.0290000438690186,
			1.1059999465942383,
			1.190000057220459,
			1.2740000486373901,
			1.3580000400543213,
			1.4420000314712524,
			1.5260000228881836,
			1.6100000143051147,
		},
		{
			0.7198200225830078,
			0.778410017490387,
			0.8370000123977661,
			0.9207000136375427,
			0.9792900085449219,
			1.0462499856948853,
			1.138319969177246,
			1.230389952659607,
			1.3224600553512573,
			1.4228999614715576,
			1.5233399868011475,
			1.6237800121307373,
			1.7242200374603271,
			1.8246599435806274,
			1.9250999689102173,
		},
	}
	// skill: rtSlash = [8]
	rtSlash = []float64{
		0.6019999980926514,
		0.6510000228881836,
		0.699999988079071,
		0.7699999809265137,
		0.8190000057220459,
		0.875,
		0.9520000219345093,
		1.0290000438690186,
		1.1059999465942383,
		1.190000057220459,
		1.2740000486373901,
		1.3580000400543213,
		1.4420000314712524,
		1.5260000228881836,
		1.6100000143051147,
	}
	// skill: skill = [0]
	skill = []float64{
		0.7200000286102295,
		0.7739999890327454,
		0.828000009059906,
		0.8999999761581421,
		0.9539999961853027,
		1.0080000162124634,
		1.0800000429153442,
		1.1519999504089355,
		1.2239999771118164,
		1.2960000038146973,
		1.3680000305175781,
		1.440000057220459,
		1.5299999713897705,
		1.6200000047683716,
		1.7100000381469727,
	}
	// burst: burst = [2]
	burst = []float64{
		3.7839999198913574,
		4.067800045013428,
		4.351600170135498,
		4.730000019073486,
		5.013800144195557,
		5.297599792480469,
		5.676000118255615,
		6.0543999671936035,
		6.432799816131592,
		6.811200141906738,
		7.189599990844727,
		7.567999839782715,
		8.041000366210938,
		8.513999938964844,
		8.987000465393066,
	}
	// burst: meleeBurst = [0]
	meleeBurst = []float64{
		4.639999866485596,
		4.98799991607666,
		5.335999965667725,
		5.800000190734863,
		6.1479997634887695,
		6.495999813079834,
		6.960000038146973,
		7.423999786376953,
		7.888000011444092,
		8.35200023651123,
		8.815999984741211,
		9.279999732971191,
		9.859999656677246,
		10.4399995803833,
		11.020000457763672,
	}
	// burst: rtBlast = [1]
	rtBlast = []float64{
		1.2000000476837158,
		1.2899999618530273,
		1.3799999952316284,
		1.5,
		1.590000033378601,
		1.6799999475479126,
		1.7999999523162842,
		1.9199999570846558,
		2.0399999618530273,
		2.1600000858306885,
		2.2799999713897705,
		2.4000000953674316,
		2.549999952316284,
		2.700000047683716,
		2.8499999046325684,
	}
)
