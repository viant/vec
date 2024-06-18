package search

import "testing"

func TestCosine(t *testing.T) {
	var tests = []struct {
		v1, v2 []float32
		output float32
	}{
		{v1: []float32{1, 2, 3, 4, 8},
			v2:     []float32{1, 2, 3, 4, 5},
			output: 0.026462317,
		},
		{v1: []float32{1, 2, 3},
			v2:     []float32{1, 2, 3},
			output: 0},
		{v1: []float32{1, 2, 3, 4},
			v2:     []float32{1, 2, 3, 4},
			output: 0},

		{v1: []float32{-0.48152837, -0.9233136, -3.2174091, 0.18817312, 1.0239712, 0.537516, -0.1355207, 1.2107713, -0.16657588, -0.75681144, -0.5795268, -0.7529298, 0.7570292, -1.4908372, -0.92496884, -2.0757003, -1.0849329, -0.06319283, 0.89413524, 0.15037605, -0.8991889, -0.1384988, 0.3735187, 0.9532701, 3.7232246, 0.6240758, 1.2555892, -0.40052778, -0.78026164, 0.78998053, 0.8411393, -0.7796171, -0.081749454, 0.20354177, 1.6018617, -1.2285175, -0.52918386, 0.6926404, 0.6901233, -0.5641664, 0.52443236, -0.8759955, 0.4798163, 0.024029626, 2.5422654, -1.6009299, -0.2767235, 0.40043604, 0.9862171, -0.15648642, -0.62372494, 0.53350985, 1.3108852, 0.11291069, 2.488056, -0.27698934, 0.69433826, -1.7402563, 0.33753332, 0.8440583, 1.0157287, 1.664056, 0.5567892, 1.3590034, 0.5874456, -0.1302745, -0.18488652, 1.0962113, -0.07639228, -1.5551255, 0.90824217, 0.9435294, 0.24992238, 0.046190158, 0.9796589, 0.15338026, 0.33330208, -0.595007, 0.3785809, 0.755604, 0.60600674, -0.7329147, 0.63378763, 0.4955551, 0.30737805, 0.033930197, 0.05713391, -0.47152558, -0.24528675, -0.11068135, 0.649512, 1.21314, 1.2610705, 0.20822904, -1.4788649, 0.13393652, 0.72075874, 0.3123695, -0.12744737, -0.9181958, 0.68768495, 0.98774797, -0.18158372, -0.24322316, 1.4179947, 0.41308767, 0.22775759, -1.1111484, -0.9217407, 0.40848225, -0.45400017, 1.1362616, 0.29854506, -0.8821255, 0.19714928, 0.8859011, -0.4426835, -1.6458354, 0.33348596, 0.8533795, -0.70031506, -1.4923981, 0.8635303, 1.0468464, 0.10373614, 0.4509778, -0.5181635, 0.059868157, 0.19541834, -0.25151467, 0.77375585, -1.2270973, 0.043475877, 0.1801568, -0.24104616, 1.4657254, -0.83110875, -2.200087, -0.6295781, 0.23697604, 1.3182279, 0.21964897, 0.15949953, -0.9891008, 0.9379275, -1.9905164, 0.44799554, -1.3846635, -0.19984552, -0.083180286, -0.039368667, 0.4008065, -0.61750543, 0.10130224, 0.9562229, -0.70481336, -0.3071887, 0.33269727, 1.0820564, -0.87225217, 0.27949533, -1.3149992, 0.47651798, -0.13200113, 0.09143972, -1.6568116, -0.6010761, 0.44805658, 0.3558244, -0.6143907, -1.4535742, -1.1309178, 0.9347337, 0.4499147, -0.4877949, -0.19307737, -0.7422137, 0.021937797, 0.73073286, -0.6531677, -0.08348158, -0.123009026, 0.24812672, 0.8222479, -0.67936754, -0.03646995, -1.554987, -0.3448236, -1.3641866, -0.15069401, 0.47042373, -0.15818714, -0.8095279, -1.7079018, -1.0239773, 0.9145745, 0.30216333, 0.902673, 1.0572184, -0.46377444, 0.32772428, 0.23604411, 1.9226086, -0.7618908, -1.3991809, 0.737344, -0.12671286, 0.48509648, 0.07830081, 0.27150762, 0.22425711, -0.21474972, -0.5408262, 0.7366384, 0.50044435, -0.63421607, 0.49235535, -0.08361161, 0.73284197, 0.56537354, 0.29506716, -0.35788256, 0.06534785, -0.7014859, 0.71503305, 0.120099135, -1.330863, -0.7765596, 0.3203357, -0.13024248, -0.09789792, -2.3026602, -0.0858212, 0.66458344, -0.021387637, 0.58201754, -0.009115983, 0.298184, 0.32770485, -0.61347353, -0.75143254, 0.87891984, -0.27242887, -0.5693034, -1.609992, 0.7793802, 0.27315345, -0.9728371, 0.28884837, 0.2025704, -2.477388, -0.020010546, 0.83945143, 0.85488814, 0.9491195, -0.46582282, 0.046570983, 0.61150867, 0.73044956, -0.6849804, 0.3089575, -0.007435011, 0.9408022, -0.03378077, 0.558615, 1.5931356, -0.7759814, -0.26532048, 0.25314063, -0.9441165, -0.21081485, 0.26620007, -1.0742505, 1.3613102, 0.078597054, -1.2300674, 0.7206876, 0.26140505, -1.526103, 0.23775779, 0.64223135, 0.016841117, -0.88364244, -0.6217705, 0.5679086, -0.9033274, 0.22036605, 0.4990582, -0.3167125, -0.2735901, 0.45441318, 1.0462785, 0.5965737, -0.51666975, 0.13218132, 0.86853504, 1.4398018, 0.2740142, 0.25748384, -0.2697761, 0.9442301, -0.6055231, 1.4126737, -0.37891597, 0.44880587, 0.009073608, 0.55473304, -0.940343, 0.76804966, 0.007302016, -0.6087611, 1.9026344, -0.9128766, 0.49901897, -1.739627, 0.63540435, 1.0969976, -0.9523704, 1.2098573, -0.7476375, 0.15766582, -0.5439461, 0.25082344, -0.2778067, -0.36674824, 0.15191577, 0.043076668, 0.12959833, 1.2371875, -0.9153108, -0.5982144, 0.8991062, -0.68823564, -1.0923796, -0.4667385, 0.42403018, -0.092934765, -0.7243048, -0.24628131, 0.6174033, -0.4420874, -1.2254388, 0.4835207, -1.472428, -0.8337005, 0.3752591, 0.09067496, 0.20247367, 0.67010164, -0.42029774, -0.19049674, -0.036616683, -1.2667367, 0.083018444, -0.3246094, -0.4982129, 0.08512106, 1.3069637, 0.20103371, -0.43868005, 0.5997726, -1.2203635, 0.8010277, -0.8870668, -0.4821588, 0.96218216, 0.7868718, -0.16525218, 0.676179, -0.74215317, -1.3567107, -0.94742537, 0.047654904, 0.49878678, 0.35454756, -0.1907419, -0.90546227, 0.85048485, 0.2895674, 1.571914, 0.14620215, -0.12404446, -0.867313, -0.060033634, 0.44291633, 0.07850296, -0.9937012, -0.42930907, -0.27286363, 1.4816728, -0.72327566, -1.5155468, 0.8064755, 0.40818477, -0.3107036, -0.6801929, 0.6732247, -0.785411, -1.0793654, 0.25419238, -1.033352, -1.4140966, -0.12748173, -0.51928407, 0.9674828, -0.4260699, -1.8771431, -0.5483761, 0.1457074, -0.34260362, 0.6637039, 0.9944253, -0.5000242, -1.3535061, -1.310367, 1.5137237, 2.063428, -1.3627845, -0.27402556, 1.1285495, 1.1935613, 1.751186, -0.37801105, -0.122339204, 0.8697192, 0.27019858, -0.12050609, 0.2514539, -0.213761, -0.98292476, 0.11860454, -0.5062104, 0.62642336, 0.3798346, -0.62619644, -0.63480115, 0.6145363, 0.15920027, 1.0346069, 1.271251, -0.6315867, -1.2108663, 0.57045764, -0.98329866, -0.44920292, 1.380247, 0.35425717, 0.27412575, -0.02665345, -0.11358385, 0.13299659, 0.6706591, 1.0233626, 0.30069482, 1.0516145, 0.49312574, -0.1867784, 0.7270293, -0.12394741, 1.8350058, 1.2667503, 0.10459163, -1.8404145, 0.6131768, 0.5385249, -1.2097481, -0.47319415, 0.99665093, -0.33528334, -0.5780113, -0.63881385, 1.3345476, -0.8225533, -0.29508072, -0.9633077, 0.6717504, -0.37589532, 0.41233152, -0.9535013, 0.44962043, -0.32348135, -0.6494948, -1.621438, -0.861135, 0.72502595, -0.16735303, 0.53825355, -0.7820088, 1.1519063, -1.6217488, 1.2226305, 1.7184869, 1.0873506, 0.50660473, -0.14451276, -1.316133, 0.22886144, 0.2928434, -0.6059728, -0.9244834, 1.4458643, 0.0726026, -0.290402, 0.19633989, 1.7361274, 0.17787883, -0.4467752, -1.4312496, -1.1792798, 0.6568744, -0.2353363, -0.43751293, 1.3941282, 0.13718261, 1.512648, 0.31169397, 0.46392092, -0.5926447, -0.6842333, -0.5387211, 0.77969503, -0.46056294, -1.5320672, -0.34597594, -0.93019044, 1.5610873, 0.9481094, -1.0305047, -0.643654, -0.3292068, 0.6716702, 0.28661674, -1.0739019, -0.8592372, 0.36975443, 0.52222073, -0.92491925, -0.5809222, -0.31007364, 1.1856496, -0.15246537, -1.2889712, -0.099821165, -0.451029, 0.67319465, 0.026079986, 0.21622309, 0.58434415, -0.7867023, -0.12249255, 0.55289614, -0.5527456, -0.7476981, -1.0137091, -0.39940232, -1.0391762, -0.26460105, -1.018, -0.6657628, 0.2771317, 1.6550211, -0.118730865, 0.7123262, 0.7159514, 2.2209885, -0.25631192, -0.5385833, -0.07816545, 0.28483626, 0.47128195, -0.38693196, -0.6112689, -0.64106905, -0.42099273, -0.12876903, 0.44979295, 0.7464367, 0.6973219, -0.1797649, -0.18809746, 0.43090156, -1.4821783, 1.34643, 0.35895646, 0.9858886, -0.21840629, 0.83632267, 0.21678904, -1.326447, -0.8749625, -0.05766727, -1.2058804, -0.39778537, -1.2160727, -0.3623044, -0.10821333, 0.14467457, -0.119695514, -1.813981, 0.10079424, -0.7227745, -0.14535382, 0.7652277, -1.4579706, 0.8609652, 1.5378069, 0.2689542, 0.10737092, -0.25211644, 0.29032755, 0.77771676, 0.5768255, -0.8756928, -0.025279216, 0.18640393, 0.2501849, 1.364368, -1.1608008, -0.04601893, -0.4650988, 0.83899444, -0.8602426, 0.35561138, -0.48252815, 1.8101969, -0.023175277, -0.66762316, -0.42764288, -1.5676253, 2.3226027, -0.09377426, -0.16859365, -1.4371145, -0.6393635, -0.9447458, 0.027634865, 0.4528245, 0.1686412, -0.778491, 0.061765596, 0.71673065, -0.32658687, -0.4728215, 0.029112073, -0.4678592, -0.0921162, 0.22657143, 1.3558004, 1.6546444, -0.12860808, 0.8169991, 0.7899942, -1.502492, 0.2674548, 0.24938937, -0.33215982, 0.89597845, -1.6372502, -0.5381735, -0.025380246, -0.67614436, -0.17595783, -0.5781694, 0.672135, 0.51085097, 0.021413513, 0.9697366, 0.94628876, -1.3035153, -0.8715645, -0.74060833, -0.0239655, -0.19893123, -0.7059541, -0.73656887, 0.51716363, -0.24957384, 0.872674, 1.9182698, -0.0033155065, -0.13848448, -1.4782411, -1.1526804, -0.22785807, 0.5279206, 0.7109796, 1.1145124, -0.35244876, -0.53137267, -0.94762254, 0.084059104, -1.1370723, 1.1415067, 0.15631561, -0.7199638, 0.004645355, 0.055946425, 0.65377164, -0.22102919, -0.048567824, -0.40843403, 0.3958103, 1.095192, 0.9156384, -1.9879025, -0.4727384, -0.2508134, -0.09153135, 0.400618, -0.42410254, 1.2957652, 0.58694494, 1.3448864, 0.015747169, 0.8198134, 0.103170745, -0.27809107, -0.031206653, -0.15289953, -0.29097688, 0.8450918, 0.06475736, -0.7763581, -0.8955928, -0.5786203, 1.9191916, 1.2143033, -1.5647135, 0.28472683, -0.3495334, 1.2365811, 0.3950894, -0.15681648, -0.038230617, -0.60112876, 1.4845955, -0.17692092, 0.4165346, -0.52610874, 1.036766, -0.17338172, -1.8616784, 0.66563106, 0.8475896, -0.57766837, -0.28701025, 0.81126726, 0.66136044, 1.0384926, 0.61120534, -1.1848512, 0.16732666, -1.3054559, 0.47450057, -0.4237256, 0.46354043, -0.21352848, -0.32955757, 1.5729322, 0.32031077, 0.65078527, -0.32270432, 0.66611147, -0.32647175, -0.9348624, -0.26661325, -0.47289985, 0.14995739},
			v2:     []float32{0.79259473, -0.2337938, -3.5238268, -0.6797824, 0.37302414, 0.7837593, -0.54397726, -0.04418074, 0.5285727, -0.8084628, -1.462792, 0.032606587, 1.4740391, -1.5470364, 0.09983079, -1.1446294, -0.22862977, -0.6063943, 0.08504771, 1.1193966, -0.16909716, 0.21482648, 0.68661225, 0.66543823, 4.175071, 0.32386184, 0.97810113, -0.47556987, -0.7861273, 0.16839442, 0.776994, -0.000322178, -0.56479204, -0.32903662, 1.5316194, -1.1464111, 0.72433597, -0.006980326, 0.17462811, -0.8687875, 0.19160295, -0.48662597, 0.57348645, -0.516348, 2.0477033, -0.975414, 0.37772718, 0.10701509, 0.99365544, -0.07417056, -0.69721276, -0.023380794, 1.4645977, -0.1848405, 1.7777071, 0.31387722, 0.2801861, -1.3942807, 1.5425565, 0.8856228, 0.44523108, 2.4371812, -0.3117515, 1.3454062, -0.75566435, -0.67002755, -1.5319333, 1.4163513, 0.4962346, -0.3042049, 1.0017487, -0.12295142, 0.8306027, -1.4532456, -0.76165426, 0.015192937, -0.84459233, -0.08137872, 0.9473733, 0.82436687, 0.83364296, -0.44278985, -0.062474474, 1.0493901, -0.1465756, -0.024196614, 0.769622, -0.64974594, -0.4914698, 0.41185072, -0.12660448, 0.4011047, 1.0763884, 1.3638816, -0.51565164, 0.027972791, -0.27883178, 0.768155, -1.1540868, -1.0040126, 1.4144144, 0.45689034, -0.15397857, 0.16674007, 1.929506, 1.1407464, 0.5259268, -0.21019599, -1.4113661, -0.9726001, -0.2767439, 2.3485737, 0.3774515, 0.24730943, 0.33357936, -0.076547265, 0.41992, -1.223125, 0.9738513, 1.2542086, 0.0599437, -0.701916, 0.4834637, 1.0551424, 0.0807187, -0.70321053, -0.7667097, 0.24126811, 0.58317804, -0.55555224, 0.5687615, -1.096735, -0.547077, -0.5660306, -0.7369826, 1.015132, -0.65624905, -2.4575853, 0.59058136, 0.4371084, 0.13294399, 0.44510996, -0.05926709, -1.4080355, 1.3620367, -1.4647382, 0.31473136, -1.0787888, -0.292637, -0.94722384, -0.09846927, 0.06259295, -0.15852596, 0.61957467, 0.62757134, -1.1241978, -0.67986566, -0.34749135, 0.94051415, 0.5027368, 0.95739275, -0.75989723, 0.070335194, 0.8248467, 0.021537721, -0.73434675, 0.5223651, 1.3834717, 0.8527136, -0.0678476, -1.3864589, -1.0259968, 1.071432, -0.34804028, 0.11871871, -0.8198968, -0.087122135, -0.7840733, 0.4854321, -0.2295621, -0.7978474, -0.38659084, -0.26538068, 1.6493177, -1.1841247, -0.118537106, -1.4279566, -0.56018656, -0.6930443, -0.53695136, -0.39738724, 0.060429707, -0.87793756, -1.9683174, -0.27140087, 0.09022461, 0.50436735, 0.99474746, 1.123274, -1.5480992, 0.57115996, -0.56822515, 0.2649669, -0.71640754, -0.68640673, -0.01283123, 0.03389179, 1.3644886, 0.6851406, 0.70703757, 0.6841513, -0.61704266, 0.063180074, 0.5692284, 0.4626988, 0.029366463, 0.061054513, -0.19129741, 0.22485013, 0.6833995, 0.2719751, 1.0430965, -0.14861543, 0.31252724, -0.49218792, 0.4316146, -2.138631, -0.07966708, 0.67678887, -0.8454585, -0.32980067, -2.5863874, 0.7372738, 0.88028157, -0.10773681, 0.3806514, -0.13114882, 0.6160076, 0.17461677, -0.3615304, -0.34945622, 0.12647668, -0.010626564, 0.43860996, -1.1798518, 0.98884, -0.07919021, -0.41272235, 0.10237583, 0.4426036, -0.7842716, 0.18869501, 1.3225101, 0.7176156, 1.0428587, -0.76262224, -0.31713873, 0.74242735, 0.87560624, 0.012118086, -0.2708386, -0.34772107, 0.37885955, -1.1992292, -1.0159256, 0.16122976, -1.2972596, -0.40692312, 0.37457746, -0.2191833, -0.72725224, -0.20249811, -1.1844395, 1.3058139, 0.06977103, -1.5094562, 0.5771984, -0.7622927, -0.18764856, 0.46091706, 0.10278028, -0.038719043, 0.1622108, -1.0490339, 0.8658006, 0.51436555, 0.47736093, 0.47897327, -0.2221641, 0.046775084, -0.41835356, 0.3680637, 0.46472815, -0.12863676, 0.73643124, 1.7678769, 1.6366373, -1.1026446, 0.5764164, -1.139793, 1.0953741, -0.82988834, 1.813986, -0.17699656, -0.08755432, -0.95528436, 0.7660612, -0.44723696, 1.744108, 0.4119574, -0.92639464, 0.9620396, -0.26673478, 0.84739727, -1.6512128, 1.5982218, 0.41592208, -0.5243386, 1.7375863, -1.5064986, 0.78044707, -0.91871566, 0.623746, -0.33386642, -0.43772608, 0.027495548, 0.42459217, 0.84521854, -0.0875709, -0.5709031, -0.33665162, 1.167261, -0.6088976, -0.9499998, -0.04595554, -0.06785661, -1.3939023, -0.86531746, -0.14141652, 1.0483402, 0.5232639, -0.8295129, 0.34536532, -1.104911, -0.4254458, 0.4937253, -0.090891995, -0.7473824, 1.1017886, 0.20301044, -0.29199576, 0.60667205, -1.2274352, 0.7443015, 0.81413686, -0.0089604035, 0.78074574, 0.62111974, 0.017957307, -0.5469049, 2.0787492, -0.8429691, 1.0651317, -0.65180665, -0.86210597, -0.026365865, 1.0656407, -0.8623757, 0.36477083, 0.08925198, -1.5086604, -0.9842552, 0.2501043, 0.5249778, 0.65732145, -0.54298663, -0.96471, -0.0760082, -1.2182826, 0.60394835, 0.049937043, -1.0474427, -1.0940294, 0.62707865, 0.044008173, -0.53940886, -0.39740103, -0.61077714, -0.26663193, 0.7056042, -0.3807878, -2.083685, 0.009954967, 0.23637737, 0.27538517, -0.017982796, 0.04779286, -0.19257301, 0.40488648, 0.07594559, -0.9073357, -1.8061352, 0.09890094, -0.54160416, 1.2092317, 0.5943776, -1.770674, -0.70543295, 0.7339517, -0.51125157, 1.1045377, 0.6358604, 0.18588883, -1.7117001, 0.102802224, 0.4521395, 2.2323716, 0.23505792, -0.42742008, 0.9391152, 0.58735955, 1.7170625, -0.35504377, -0.51361656, 0.004195521, 0.45498013, -0.64505285, -0.49341464, -0.8722012, -1.5052272, -0.038168307, -0.043984205, 0.3886037, -0.25403276, -0.3820205, 0.27881664, -0.37273237, 0.31087047, 0.8320368, 1.915541, -0.75911987, -0.3013711, -0.38916078, -0.052075706, -0.10864355, 1.0382979, 0.64845616, -0.20454343, -0.4593991, -1.0038705, 0.29872167, -0.3643093, 0.68791354, 0.13818401, 1.240767, 0.70042944, -0.22464147, 0.72171354, -0.9254472, 1.1596119, 1.4638703, -0.46220925, -1.0792314, -0.3386259, -0.8519223, -0.9377905, 0.032379206, -0.24437489, 0.21810979, 0.82704586, -0.6294444, 1.0167073, -0.22322115, -0.48485464, -0.8876802, 1.4977512, 0.19985855, 0.80396897, 1.3215015, -0.25615168, -0.75053334, -0.6676252, -2.0982199, -0.8604602, 0.4162949, 0.3576231, 0.31141865, -0.5857935, 1.2024496, -0.10103811, 0.94472384, 0.92938566, 0.6171248, -0.43201095, -0.10707618, -0.76827717, -1.4111861, 0.069235474, 0.49377853, -0.6700959, 1.2746884, 0.17247766, 0.21311824, 0.46552896, 1.4299009, 0.7953531, 0.6155269, -1.0366261, -1.1062882, 1.0172001, -0.3770462, -0.63856363, 0.48476207, 0.753886, 0.9255824, 0.26699892, -0.7560532, -0.42972696, -0.6308938, -0.27333987, 0.82377183, -0.76752865, -0.8241781, -0.80687994, -0.86332357, 1.3512888, 0.47235, -0.062901914, -0.50133914, -0.49218607, 1.0891556, 0.10795422, -1.0144705, -1.0423889, 0.40363264, -0.08698335, -1.8367128, -1.3475386, 1.1504307, 0.58302814, -0.6506213, -1.2502488, 0.23788826, -0.8388028, 0.10572763, 1.3849111, -0.030442767, 0.7717763, -0.83239496, -1.9374269, 1.4642407, -1.3262877, -0.23713472, -1.3840811, -0.54843813, -1.09361, 0.65765923, -1.2376026, -0.9401457, -0.029645868, 0.9534676, -0.04691778, 0.8622837, 0.26962355, 0.93310094, -1.6290861, -0.38343212, -0.53144443, -0.14405118, 0.59138983, -0.6948143, -0.5306547, -0.35516995, 0.65113044, -0.002249714, -0.283839, 1.2523994, 0.19689539, 0.005296655, -0.38785636, 1.2713007, -0.643932, 0.5808283, 0.72826713, 0.61490166, 0.3451336, 0.43224746, -0.007995486, -0.4000072, -0.19717887, -0.17058994, 0.12491931, -0.5395574, -0.7534895, 0.20900846, -0.5765571, -0.33123407, 0.26580748, -0.73630923, -0.4886188, -0.40352848, -0.024798952, 1.1167312, -1.8121498, 0.121740885, 1.3917418, -0.36938718, 0.021185733, -0.3161856, -0.28234977, 1.5324662, 0.5171375, -1.091852, 0.38685986, 0.71577173, -0.49378395, 1.4271356, -0.10178216, -0.11572403, -0.9111893, 0.5343625, -0.88408387, -0.3173598, -1.1050034, 1.9068451, -0.026346672, -1.577805, -1.2054079, -0.6667847, 2.2555547, 0.17100832, 0.4421487, -1.4646888, -0.6964792, -0.8254282, 0.20609814, -0.15114594, -0.25856236, -0.16574994, 1.060196, 1.0767117, -0.7768328, -0.5820332, -0.044080924, -1.0714506, -0.38727042, 0.16535659, 0.54267305, 0.69908077, -0.20411104, 1.5637348, 1.2224724, -0.57318246, -0.8794975, 0.4183081, -0.4004147, 1.2175434, -1.2535589, -1.1905729, 0.26656315, -0.23623033, 0.46363038, -0.5532423, -0.029677011, 0.52004635, -0.27674642, -0.18623446, 0.1482059, -1.0052068, 0.06828004, -0.46109632, -0.092749745, -0.1461868, -0.41643947, -0.25797963, 1.079101, -0.48877788, 0.5428395, 1.24681, 0.37028635, -0.21724439, -1.2412205, -0.9021638, -0.39376855, -0.3008405, -0.4574111, 1.0381953, -0.26044142, -0.07045314, -0.74299556, -0.42871034, -0.89191866, 0.45772702, 0.051298574, -0.83936286, -0.9838704, 0.0026706588, -0.33328205, -0.028797695, 0.5782403, 0.39765182, 1.0093777, -0.036723502, 0.81237346, -1.4724786, 0.69834596, 0.12422679, 0.037060402, 0.60380626, 0.14557908, 0.07907937, 1.0044162, 1.2319636, -0.23354784, 1.9470067, 0.40222007, 0.27360803, -1.16805, -0.3078477, 0.026089173, 0.9148015, -0.15032469, -0.7708012, -1.000283, 0.2544599, 1.8954918, 0.25876352, -1.4745554, 0.5720192, 0.25725845, 0.88494, -0.2649834, -0.7182274, 0.87877566, -0.8262619, 0.8974995, -1.4166473, 0.8281653, -0.05496105, 0.94833344, -0.3545212, -1.9316326, 0.6161145, 0.6759596, 0.14657244, -0.713806, 1.0426415, 0.7153422, 1.2666845, 0.13204168, -0.9634924, -0.84210515, -1.3326854, 0.75225854, -0.20051974, 1.1867173, -0.15015687, 0.079517074, 2.2878475, 0.2789643, 0.33755586, -0.27782178, 0.9329344, 0.4237718, -0.787215, 0.19184834, -0.9254514, -0.10571754},
			output: 0.26687056,
		},
	}
	for _, test := range tests {
		v := Float32s(test.v1)
		got := v.cosineDistance(test.v2)
		if got != test.output {
			t.Errorf("cosineDistance(%v, %v) = %v, want %v", test.v1, test.v2, got, test.output)
		}
	}
}

// BenchmarkVector_CosineDistance-10    	 1274827	       920.7 ns/op
func BenchmarkVector_CosineDistance(b *testing.B) {
	v1 := []float32{-0.48152837, -0.9233136, -3.2174091, 0.18817312, 1.0239712, 0.537516, -0.1355207, 1.2107713, -0.16657588, -0.75681144, -0.5795268, -0.7529298, 0.7570292, -1.4908372, -0.92496884, -2.0757003, -1.0849329, -0.06319283, 0.89413524, 0.15037605, -0.8991889, -0.1384988, 0.3735187, 0.9532701, 3.7232246, 0.6240758, 1.2555892, -0.40052778, -0.78026164, 0.78998053, 0.8411393, -0.7796171, -0.081749454, 0.20354177, 1.6018617, -1.2285175, -0.52918386, 0.6926404, 0.6901233, -0.5641664, 0.52443236, -0.8759955, 0.4798163, 0.024029626, 2.5422654, -1.6009299, -0.2767235, 0.40043604, 0.9862171, -0.15648642, -0.62372494, 0.53350985, 1.3108852, 0.11291069, 2.488056, -0.27698934, 0.69433826, -1.7402563, 0.33753332, 0.8440583, 1.0157287, 1.664056, 0.5567892, 1.3590034, 0.5874456, -0.1302745, -0.18488652, 1.0962113, -0.07639228, -1.5551255, 0.90824217, 0.9435294, 0.24992238, 0.046190158, 0.9796589, 0.15338026, 0.33330208, -0.595007, 0.3785809, 0.755604, 0.60600674, -0.7329147, 0.63378763, 0.4955551, 0.30737805, 0.033930197, 0.05713391, -0.47152558, -0.24528675, -0.11068135, 0.649512, 1.21314, 1.2610705, 0.20822904, -1.4788649, 0.13393652, 0.72075874, 0.3123695, -0.12744737, -0.9181958, 0.68768495, 0.98774797, -0.18158372, -0.24322316, 1.4179947, 0.41308767, 0.22775759, -1.1111484, -0.9217407, 0.40848225, -0.45400017, 1.1362616, 0.29854506, -0.8821255, 0.19714928, 0.8859011, -0.4426835, -1.6458354, 0.33348596, 0.8533795, -0.70031506, -1.4923981, 0.8635303, 1.0468464, 0.10373614, 0.4509778, -0.5181635, 0.059868157, 0.19541834, -0.25151467, 0.77375585, -1.2270973, 0.043475877, 0.1801568, -0.24104616, 1.4657254, -0.83110875, -2.200087, -0.6295781, 0.23697604, 1.3182279, 0.21964897, 0.15949953, -0.9891008, 0.9379275, -1.9905164, 0.44799554, -1.3846635, -0.19984552, -0.083180286, -0.039368667, 0.4008065, -0.61750543, 0.10130224, 0.9562229, -0.70481336, -0.3071887, 0.33269727, 1.0820564, -0.87225217, 0.27949533, -1.3149992, 0.47651798, -0.13200113, 0.09143972, -1.6568116, -0.6010761, 0.44805658, 0.3558244, -0.6143907, -1.4535742, -1.1309178, 0.9347337, 0.4499147, -0.4877949, -0.19307737, -0.7422137, 0.021937797, 0.73073286, -0.6531677, -0.08348158, -0.123009026, 0.24812672, 0.8222479, -0.67936754, -0.03646995, -1.554987, -0.3448236, -1.3641866, -0.15069401, 0.47042373, -0.15818714, -0.8095279, -1.7079018, -1.0239773, 0.9145745, 0.30216333, 0.902673, 1.0572184, -0.46377444, 0.32772428, 0.23604411, 1.9226086, -0.7618908, -1.3991809, 0.737344, -0.12671286, 0.48509648, 0.07830081, 0.27150762, 0.22425711, -0.21474972, -0.5408262, 0.7366384, 0.50044435, -0.63421607, 0.49235535, -0.08361161, 0.73284197, 0.56537354, 0.29506716, -0.35788256, 0.06534785, -0.7014859, 0.71503305, 0.120099135, -1.330863, -0.7765596, 0.3203357, -0.13024248, -0.09789792, -2.3026602, -0.0858212, 0.66458344, -0.021387637, 0.58201754, -0.009115983, 0.298184, 0.32770485, -0.61347353, -0.75143254, 0.87891984, -0.27242887, -0.5693034, -1.609992, 0.7793802, 0.27315345, -0.9728371, 0.28884837, 0.2025704, -2.477388, -0.020010546, 0.83945143, 0.85488814, 0.9491195, -0.46582282, 0.046570983, 0.61150867, 0.73044956, -0.6849804, 0.3089575, -0.007435011, 0.9408022, -0.03378077, 0.558615, 1.5931356, -0.7759814, -0.26532048, 0.25314063, -0.9441165, -0.21081485, 0.26620007, -1.0742505, 1.3613102, 0.078597054, -1.2300674, 0.7206876, 0.26140505, -1.526103, 0.23775779, 0.64223135, 0.016841117, -0.88364244, -0.6217705, 0.5679086, -0.9033274, 0.22036605, 0.4990582, -0.3167125, -0.2735901, 0.45441318, 1.0462785, 0.5965737, -0.51666975, 0.13218132, 0.86853504, 1.4398018, 0.2740142, 0.25748384, -0.2697761, 0.9442301, -0.6055231, 1.4126737, -0.37891597, 0.44880587, 0.009073608, 0.55473304, -0.940343, 0.76804966, 0.007302016, -0.6087611, 1.9026344, -0.9128766, 0.49901897, -1.739627, 0.63540435, 1.0969976, -0.9523704, 1.2098573, -0.7476375, 0.15766582, -0.5439461, 0.25082344, -0.2778067, -0.36674824, 0.15191577, 0.043076668, 0.12959833, 1.2371875, -0.9153108, -0.5982144, 0.8991062, -0.68823564, -1.0923796, -0.4667385, 0.42403018, -0.092934765, -0.7243048, -0.24628131, 0.6174033, -0.4420874, -1.2254388, 0.4835207, -1.472428, -0.8337005, 0.3752591, 0.09067496, 0.20247367, 0.67010164, -0.42029774, -0.19049674, -0.036616683, -1.2667367, 0.083018444, -0.3246094, -0.4982129, 0.08512106, 1.3069637, 0.20103371, -0.43868005, 0.5997726, -1.2203635, 0.8010277, -0.8870668, -0.4821588, 0.96218216, 0.7868718, -0.16525218, 0.676179, -0.74215317, -1.3567107, -0.94742537, 0.047654904, 0.49878678, 0.35454756, -0.1907419, -0.90546227, 0.85048485, 0.2895674, 1.571914, 0.14620215, -0.12404446, -0.867313, -0.060033634, 0.44291633, 0.07850296, -0.9937012, -0.42930907, -0.27286363, 1.4816728, -0.72327566, -1.5155468, 0.8064755, 0.40818477, -0.3107036, -0.6801929, 0.6732247, -0.785411, -1.0793654, 0.25419238, -1.033352, -1.4140966, -0.12748173, -0.51928407, 0.9674828, -0.4260699, -1.8771431, -0.5483761, 0.1457074, -0.34260362, 0.6637039, 0.9944253, -0.5000242, -1.3535061, -1.310367, 1.5137237, 2.063428, -1.3627845, -0.27402556, 1.1285495, 1.1935613, 1.751186, -0.37801105, -0.122339204, 0.8697192, 0.27019858, -0.12050609, 0.2514539, -0.213761, -0.98292476, 0.11860454, -0.5062104, 0.62642336, 0.3798346, -0.62619644, -0.63480115, 0.6145363, 0.15920027, 1.0346069, 1.271251, -0.6315867, -1.2108663, 0.57045764, -0.98329866, -0.44920292, 1.380247, 0.35425717, 0.27412575, -0.02665345, -0.11358385, 0.13299659, 0.6706591, 1.0233626, 0.30069482, 1.0516145, 0.49312574, -0.1867784, 0.7270293, -0.12394741, 1.8350058, 1.2667503, 0.10459163, -1.8404145, 0.6131768, 0.5385249, -1.2097481, -0.47319415, 0.99665093, -0.33528334, -0.5780113, -0.63881385, 1.3345476, -0.8225533, -0.29508072, -0.9633077, 0.6717504, -0.37589532, 0.41233152, -0.9535013, 0.44962043, -0.32348135, -0.6494948, -1.621438, -0.861135, 0.72502595, -0.16735303, 0.53825355, -0.7820088, 1.1519063, -1.6217488, 1.2226305, 1.7184869, 1.0873506, 0.50660473, -0.14451276, -1.316133, 0.22886144, 0.2928434, -0.6059728, -0.9244834, 1.4458643, 0.0726026, -0.290402, 0.19633989, 1.7361274, 0.17787883, -0.4467752, -1.4312496, -1.1792798, 0.6568744, -0.2353363, -0.43751293, 1.3941282, 0.13718261, 1.512648, 0.31169397, 0.46392092, -0.5926447, -0.6842333, -0.5387211, 0.77969503, -0.46056294, -1.5320672, -0.34597594, -0.93019044, 1.5610873, 0.9481094, -1.0305047, -0.643654, -0.3292068, 0.6716702, 0.28661674, -1.0739019, -0.8592372, 0.36975443, 0.52222073, -0.92491925, -0.5809222, -0.31007364, 1.1856496, -0.15246537, -1.2889712, -0.099821165, -0.451029, 0.67319465, 0.026079986, 0.21622309, 0.58434415, -0.7867023, -0.12249255, 0.55289614, -0.5527456, -0.7476981, -1.0137091, -0.39940232, -1.0391762, -0.26460105, -1.018, -0.6657628, 0.2771317, 1.6550211, -0.118730865, 0.7123262, 0.7159514, 2.2209885, -0.25631192, -0.5385833, -0.07816545, 0.28483626, 0.47128195, -0.38693196, -0.6112689, -0.64106905, -0.42099273, -0.12876903, 0.44979295, 0.7464367, 0.6973219, -0.1797649, -0.18809746, 0.43090156, -1.4821783, 1.34643, 0.35895646, 0.9858886, -0.21840629, 0.83632267, 0.21678904, -1.326447, -0.8749625, -0.05766727, -1.2058804, -0.39778537, -1.2160727, -0.3623044, -0.10821333, 0.14467457, -0.119695514, -1.813981, 0.10079424, -0.7227745, -0.14535382, 0.7652277, -1.4579706, 0.8609652, 1.5378069, 0.2689542, 0.10737092, -0.25211644, 0.29032755, 0.77771676, 0.5768255, -0.8756928, -0.025279216, 0.18640393, 0.2501849, 1.364368, -1.1608008, -0.04601893, -0.4650988, 0.83899444, -0.8602426, 0.35561138, -0.48252815, 1.8101969, -0.023175277, -0.66762316, -0.42764288, -1.5676253, 2.3226027, -0.09377426, -0.16859365, -1.4371145, -0.6393635, -0.9447458, 0.027634865, 0.4528245, 0.1686412, -0.778491, 0.061765596, 0.71673065, -0.32658687, -0.4728215, 0.029112073, -0.4678592, -0.0921162, 0.22657143, 1.3558004, 1.6546444, -0.12860808, 0.8169991, 0.7899942, -1.502492, 0.2674548, 0.24938937, -0.33215982, 0.89597845, -1.6372502, -0.5381735, -0.025380246, -0.67614436, -0.17595783, -0.5781694, 0.672135, 0.51085097, 0.021413513, 0.9697366, 0.94628876, -1.3035153, -0.8715645, -0.74060833, -0.0239655, -0.19893123, -0.7059541, -0.73656887, 0.51716363, -0.24957384, 0.872674, 1.9182698, -0.0033155065, -0.13848448, -1.4782411, -1.1526804, -0.22785807, 0.5279206, 0.7109796, 1.1145124, -0.35244876, -0.53137267, -0.94762254, 0.084059104, -1.1370723, 1.1415067, 0.15631561, -0.7199638, 0.004645355, 0.055946425, 0.65377164, -0.22102919, -0.048567824, -0.40843403, 0.3958103, 1.095192, 0.9156384, -1.9879025, -0.4727384, -0.2508134, -0.09153135, 0.400618, -0.42410254, 1.2957652, 0.58694494, 1.3448864, 0.015747169, 0.8198134, 0.103170745, -0.27809107, -0.031206653, -0.15289953, -0.29097688, 0.8450918, 0.06475736, -0.7763581, -0.8955928, -0.5786203, 1.9191916, 1.2143033, -1.5647135, 0.28472683, -0.3495334, 1.2365811, 0.3950894, -0.15681648, -0.038230617, -0.60112876, 1.4845955, -0.17692092, 0.4165346, -0.52610874, 1.036766, -0.17338172, -1.8616784, 0.66563106, 0.8475896, -0.57766837, -0.28701025, 0.81126726, 0.66136044, 1.0384926, 0.61120534, -1.1848512, 0.16732666, -1.3054559, 0.47450057, -0.4237256, 0.46354043, -0.21352848, -0.32955757, 1.5729322, 0.32031077, 0.65078527, -0.32270432, 0.66611147, -0.32647175, -0.9348624, -0.26661325, -0.47289985, 0.14995739}
	v2 := []float32{0.79259473, -0.2337938, -3.5238268, -0.6797824, 0.37302414, 0.7837593, -0.54397726, -0.04418074, 0.5285727, -0.8084628, -1.462792, 0.032606587, 1.4740391, -1.5470364, 0.09983079, -1.1446294, -0.22862977, -0.6063943, 0.08504771, 1.1193966, -0.16909716, 0.21482648, 0.68661225, 0.66543823, 4.175071, 0.32386184, 0.97810113, -0.47556987, -0.7861273, 0.16839442, 0.776994, -0.000322178, -0.56479204, -0.32903662, 1.5316194, -1.1464111, 0.72433597, -0.006980326, 0.17462811, -0.8687875, 0.19160295, -0.48662597, 0.57348645, -0.516348, 2.0477033, -0.975414, 0.37772718, 0.10701509, 0.99365544, -0.07417056, -0.69721276, -0.023380794, 1.4645977, -0.1848405, 1.7777071, 0.31387722, 0.2801861, -1.3942807, 1.5425565, 0.8856228, 0.44523108, 2.4371812, -0.3117515, 1.3454062, -0.75566435, -0.67002755, -1.5319333, 1.4163513, 0.4962346, -0.3042049, 1.0017487, -0.12295142, 0.8306027, -1.4532456, -0.76165426, 0.015192937, -0.84459233, -0.08137872, 0.9473733, 0.82436687, 0.83364296, -0.44278985, -0.062474474, 1.0493901, -0.1465756, -0.024196614, 0.769622, -0.64974594, -0.4914698, 0.41185072, -0.12660448, 0.4011047, 1.0763884, 1.3638816, -0.51565164, 0.027972791, -0.27883178, 0.768155, -1.1540868, -1.0040126, 1.4144144, 0.45689034, -0.15397857, 0.16674007, 1.929506, 1.1407464, 0.5259268, -0.21019599, -1.4113661, -0.9726001, -0.2767439, 2.3485737, 0.3774515, 0.24730943, 0.33357936, -0.076547265, 0.41992, -1.223125, 0.9738513, 1.2542086, 0.0599437, -0.701916, 0.4834637, 1.0551424, 0.0807187, -0.70321053, -0.7667097, 0.24126811, 0.58317804, -0.55555224, 0.5687615, -1.096735, -0.547077, -0.5660306, -0.7369826, 1.015132, -0.65624905, -2.4575853, 0.59058136, 0.4371084, 0.13294399, 0.44510996, -0.05926709, -1.4080355, 1.3620367, -1.4647382, 0.31473136, -1.0787888, -0.292637, -0.94722384, -0.09846927, 0.06259295, -0.15852596, 0.61957467, 0.62757134, -1.1241978, -0.67986566, -0.34749135, 0.94051415, 0.5027368, 0.95739275, -0.75989723, 0.070335194, 0.8248467, 0.021537721, -0.73434675, 0.5223651, 1.3834717, 0.8527136, -0.0678476, -1.3864589, -1.0259968, 1.071432, -0.34804028, 0.11871871, -0.8198968, -0.087122135, -0.7840733, 0.4854321, -0.2295621, -0.7978474, -0.38659084, -0.26538068, 1.6493177, -1.1841247, -0.118537106, -1.4279566, -0.56018656, -0.6930443, -0.53695136, -0.39738724, 0.060429707, -0.87793756, -1.9683174, -0.27140087, 0.09022461, 0.50436735, 0.99474746, 1.123274, -1.5480992, 0.57115996, -0.56822515, 0.2649669, -0.71640754, -0.68640673, -0.01283123, 0.03389179, 1.3644886, 0.6851406, 0.70703757, 0.6841513, -0.61704266, 0.063180074, 0.5692284, 0.4626988, 0.029366463, 0.061054513, -0.19129741, 0.22485013, 0.6833995, 0.2719751, 1.0430965, -0.14861543, 0.31252724, -0.49218792, 0.4316146, -2.138631, -0.07966708, 0.67678887, -0.8454585, -0.32980067, -2.5863874, 0.7372738, 0.88028157, -0.10773681, 0.3806514, -0.13114882, 0.6160076, 0.17461677, -0.3615304, -0.34945622, 0.12647668, -0.010626564, 0.43860996, -1.1798518, 0.98884, -0.07919021, -0.41272235, 0.10237583, 0.4426036, -0.7842716, 0.18869501, 1.3225101, 0.7176156, 1.0428587, -0.76262224, -0.31713873, 0.74242735, 0.87560624, 0.012118086, -0.2708386, -0.34772107, 0.37885955, -1.1992292, -1.0159256, 0.16122976, -1.2972596, -0.40692312, 0.37457746, -0.2191833, -0.72725224, -0.20249811, -1.1844395, 1.3058139, 0.06977103, -1.5094562, 0.5771984, -0.7622927, -0.18764856, 0.46091706, 0.10278028, -0.038719043, 0.1622108, -1.0490339, 0.8658006, 0.51436555, 0.47736093, 0.47897327, -0.2221641, 0.046775084, -0.41835356, 0.3680637, 0.46472815, -0.12863676, 0.73643124, 1.7678769, 1.6366373, -1.1026446, 0.5764164, -1.139793, 1.0953741, -0.82988834, 1.813986, -0.17699656, -0.08755432, -0.95528436, 0.7660612, -0.44723696, 1.744108, 0.4119574, -0.92639464, 0.9620396, -0.26673478, 0.84739727, -1.6512128, 1.5982218, 0.41592208, -0.5243386, 1.7375863, -1.5064986, 0.78044707, -0.91871566, 0.623746, -0.33386642, -0.43772608, 0.027495548, 0.42459217, 0.84521854, -0.0875709, -0.5709031, -0.33665162, 1.167261, -0.6088976, -0.9499998, -0.04595554, -0.06785661, -1.3939023, -0.86531746, -0.14141652, 1.0483402, 0.5232639, -0.8295129, 0.34536532, -1.104911, -0.4254458, 0.4937253, -0.090891995, -0.7473824, 1.1017886, 0.20301044, -0.29199576, 0.60667205, -1.2274352, 0.7443015, 0.81413686, -0.0089604035, 0.78074574, 0.62111974, 0.017957307, -0.5469049, 2.0787492, -0.8429691, 1.0651317, -0.65180665, -0.86210597, -0.026365865, 1.0656407, -0.8623757, 0.36477083, 0.08925198, -1.5086604, -0.9842552, 0.2501043, 0.5249778, 0.65732145, -0.54298663, -0.96471, -0.0760082, -1.2182826, 0.60394835, 0.049937043, -1.0474427, -1.0940294, 0.62707865, 0.044008173, -0.53940886, -0.39740103, -0.61077714, -0.26663193, 0.7056042, -0.3807878, -2.083685, 0.009954967, 0.23637737, 0.27538517, -0.017982796, 0.04779286, -0.19257301, 0.40488648, 0.07594559, -0.9073357, -1.8061352, 0.09890094, -0.54160416, 1.2092317, 0.5943776, -1.770674, -0.70543295, 0.7339517, -0.51125157, 1.1045377, 0.6358604, 0.18588883, -1.7117001, 0.102802224, 0.4521395, 2.2323716, 0.23505792, -0.42742008, 0.9391152, 0.58735955, 1.7170625, -0.35504377, -0.51361656, 0.004195521, 0.45498013, -0.64505285, -0.49341464, -0.8722012, -1.5052272, -0.038168307, -0.043984205, 0.3886037, -0.25403276, -0.3820205, 0.27881664, -0.37273237, 0.31087047, 0.8320368, 1.915541, -0.75911987, -0.3013711, -0.38916078, -0.052075706, -0.10864355, 1.0382979, 0.64845616, -0.20454343, -0.4593991, -1.0038705, 0.29872167, -0.3643093, 0.68791354, 0.13818401, 1.240767, 0.70042944, -0.22464147, 0.72171354, -0.9254472, 1.1596119, 1.4638703, -0.46220925, -1.0792314, -0.3386259, -0.8519223, -0.9377905, 0.032379206, -0.24437489, 0.21810979, 0.82704586, -0.6294444, 1.0167073, -0.22322115, -0.48485464, -0.8876802, 1.4977512, 0.19985855, 0.80396897, 1.3215015, -0.25615168, -0.75053334, -0.6676252, -2.0982199, -0.8604602, 0.4162949, 0.3576231, 0.31141865, -0.5857935, 1.2024496, -0.10103811, 0.94472384, 0.92938566, 0.6171248, -0.43201095, -0.10707618, -0.76827717, -1.4111861, 0.069235474, 0.49377853, -0.6700959, 1.2746884, 0.17247766, 0.21311824, 0.46552896, 1.4299009, 0.7953531, 0.6155269, -1.0366261, -1.1062882, 1.0172001, -0.3770462, -0.63856363, 0.48476207, 0.753886, 0.9255824, 0.26699892, -0.7560532, -0.42972696, -0.6308938, -0.27333987, 0.82377183, -0.76752865, -0.8241781, -0.80687994, -0.86332357, 1.3512888, 0.47235, -0.062901914, -0.50133914, -0.49218607, 1.0891556, 0.10795422, -1.0144705, -1.0423889, 0.40363264, -0.08698335, -1.8367128, -1.3475386, 1.1504307, 0.58302814, -0.6506213, -1.2502488, 0.23788826, -0.8388028, 0.10572763, 1.3849111, -0.030442767, 0.7717763, -0.83239496, -1.9374269, 1.4642407, -1.3262877, -0.23713472, -1.3840811, -0.54843813, -1.09361, 0.65765923, -1.2376026, -0.9401457, -0.029645868, 0.9534676, -0.04691778, 0.8622837, 0.26962355, 0.93310094, -1.6290861, -0.38343212, -0.53144443, -0.14405118, 0.59138983, -0.6948143, -0.5306547, -0.35516995, 0.65113044, -0.002249714, -0.283839, 1.2523994, 0.19689539, 0.005296655, -0.38785636, 1.2713007, -0.643932, 0.5808283, 0.72826713, 0.61490166, 0.3451336, 0.43224746, -0.007995486, -0.4000072, -0.19717887, -0.17058994, 0.12491931, -0.5395574, -0.7534895, 0.20900846, -0.5765571, -0.33123407, 0.26580748, -0.73630923, -0.4886188, -0.40352848, -0.024798952, 1.1167312, -1.8121498, 0.121740885, 1.3917418, -0.36938718, 0.021185733, -0.3161856, -0.28234977, 1.5324662, 0.5171375, -1.091852, 0.38685986, 0.71577173, -0.49378395, 1.4271356, -0.10178216, -0.11572403, -0.9111893, 0.5343625, -0.88408387, -0.3173598, -1.1050034, 1.9068451, -0.026346672, -1.577805, -1.2054079, -0.6667847, 2.2555547, 0.17100832, 0.4421487, -1.4646888, -0.6964792, -0.8254282, 0.20609814, -0.15114594, -0.25856236, -0.16574994, 1.060196, 1.0767117, -0.7768328, -0.5820332, -0.044080924, -1.0714506, -0.38727042, 0.16535659, 0.54267305, 0.69908077, -0.20411104, 1.5637348, 1.2224724, -0.57318246, -0.8794975, 0.4183081, -0.4004147, 1.2175434, -1.2535589, -1.1905729, 0.26656315, -0.23623033, 0.46363038, -0.5532423, -0.029677011, 0.52004635, -0.27674642, -0.18623446, 0.1482059, -1.0052068, 0.06828004, -0.46109632, -0.092749745, -0.1461868, -0.41643947, -0.25797963, 1.079101, -0.48877788, 0.5428395, 1.24681, 0.37028635, -0.21724439, -1.2412205, -0.9021638, -0.39376855, -0.3008405, -0.4574111, 1.0381953, -0.26044142, -0.07045314, -0.74299556, -0.42871034, -0.89191866, 0.45772702, 0.051298574, -0.83936286, -0.9838704, 0.0026706588, -0.33328205, -0.028797695, 0.5782403, 0.39765182, 1.0093777, -0.036723502, 0.81237346, -1.4724786, 0.69834596, 0.12422679, 0.037060402, 0.60380626, 0.14557908, 0.07907937, 1.0044162, 1.2319636, -0.23354784, 1.9470067, 0.40222007, 0.27360803, -1.16805, -0.3078477, 0.026089173, 0.9148015, -0.15032469, -0.7708012, -1.000283, 0.2544599, 1.8954918, 0.25876352, -1.4745554, 0.5720192, 0.25725845, 0.88494, -0.2649834, -0.7182274, 0.87877566, -0.8262619, 0.8974995, -1.4166473, 0.8281653, -0.05496105, 0.94833344, -0.3545212, -1.9316326, 0.6161145, 0.6759596, 0.14657244, -0.713806, 1.0426415, 0.7153422, 1.2666845, 0.13204168, -0.9634924, -0.84210515, -1.3326854, 0.75225854, -0.20051974, 1.1867173, -0.15015687, 0.079517074, 2.2878475, 0.2789643, 0.33755586, -0.27782178, 0.9329344, 0.4237718, -0.787215, 0.19184834, -0.9254514, -0.10571754}
	for i := 0; i < b.N; i++ {
		Float32s(v1).cosineDistance(v2)
	}
}
