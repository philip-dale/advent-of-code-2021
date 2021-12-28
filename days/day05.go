package days

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

type lines struct {
	data []line
}

type vent_map struct {
	x_size int
	y_size int
	data   [][]int
}

var day05_sample = lines{
	[]line{
		{point{0, 9}, point{5, 9}},
		{point{8, 0}, point{0, 8}},
		{point{9, 4}, point{3, 4}},
		{point{2, 2}, point{2, 1}},
		{point{7, 0}, point{7, 4}},
		{point{6, 4}, point{2, 0}},
		{point{0, 9}, point{2, 9}},
		{point{3, 4}, point{1, 4}},
		{point{0, 0}, point{8, 8}},
		{point{5, 5}, point{8, 2}},
	},
}

var day05_data = lines{
	[]line{
		{point{720, 475}, point{720, 669}},
		{point{268, 784}, point{595, 784}},
		{point{163, 778}, point{149, 778}},
		{point{856, 917}, point{856, 114}},
		{point{929, 964}, point{433, 964}},
		{point{344, 924}, point{332, 912}},
		{point{73, 923}, point{73, 568}},
		{point{714, 501}, point{714, 903}},
		{point{95, 968}, point{95, 696}},
		{point{575, 615}, point{575, 218}},
		{point{233, 182}, point{233, 295}},
		{point{90, 752}, point{90, 278}},
		{point{683, 21}, point{419, 21}},
		{point{714, 172}, point{60, 172}},
		{point{836, 248}, point{836, 39}},
		{point{770, 53}, point{284, 539}},
		{point{528, 682}, point{528, 583}},
		{point{407, 360}, point{708, 59}},
		{point{451, 671}, point{293, 671}},
		{point{418, 446}, point{418, 113}},
		{point{948, 283}, point{504, 727}},
		{point{975, 494}, point{669, 494}},
		{point{911, 896}, point{911, 963}},
		{point{843, 855}, point{14, 26}},
		{point{558, 340}, point{984, 766}},
		{point{260, 909}, point{333, 909}},
		{point{127, 864}, point{903, 864}},
		{point{896, 174}, point{896, 773}},
		{point{989, 981}, point{27, 19}},
		{point{525, 907}, point{525, 36}},
		{point{226, 276}, point{938, 988}},
		{point{285, 694}, point{422, 831}},
		{point{985, 857}, point{674, 546}},
		{point{559, 436}, point{133, 862}},
		{point{12, 70}, point{12, 961}},
		{point{412, 185}, point{412, 922}},
		{point{94, 174}, point{874, 954}},
		{point{435, 840}, point{688, 587}},
		{point{43, 405}, point{43, 724}},
		{point{494, 826}, point{494, 350}},
		{point{591, 985}, point{591, 936}},
		{point{440, 251}, point{401, 251}},
		{point{341, 816}, point{920, 237}},
		{point{80, 901}, point{80, 361}},
		{point{962, 544}, point{962, 252}},
		{point{26, 732}, point{105, 653}},
		{point{821, 54}, point{15, 860}},
		{point{519, 731}, point{952, 731}},
		{point{723, 869}, point{826, 869}},
		{point{752, 176}, point{110, 818}},
		{point{849, 163}, point{71, 163}},
		{point{136, 748}, point{136, 796}},
		{point{301, 390}, point{234, 390}},
		{point{87, 896}, point{968, 15}},
		{point{603, 388}, point{515, 476}},
		{point{866, 345}, point{866, 742}},
		{point{477, 263}, point{477, 635}},
		{point{391, 675}, point{391, 613}},
		{point{460, 925}, point{162, 627}},
		{point{12, 841}, point{12, 246}},
		{point{712, 964}, point{712, 887}},
		{point{18, 984}, point{972, 30}},
		{point{931, 85}, point{217, 799}},
		{point{591, 848}, point{978, 848}},
		{point{87, 210}, point{840, 963}},
		{point{816, 645}, point{816, 619}},
		{point{183, 710}, point{348, 710}},
		{point{39, 861}, point{877, 23}},
		{point{233, 529}, point{233, 717}},
		{point{20, 43}, point{935, 958}},
		{point{979, 988}, point{12, 21}},
		{point{941, 241}, point{855, 241}},
		{point{416, 507}, point{422, 501}},
		{point{410, 978}, point{971, 978}},
		{point{883, 574}, point{883, 440}},
		{point{458, 865}, point{912, 865}},
		{point{113, 814}, point{952, 814}},
		{point{942, 654}, point{545, 654}},
		{point{512, 78}, point{582, 78}},
		{point{102, 927}, point{910, 119}},
		{point{157, 874}, point{76, 793}},
		{point{309, 987}, point{760, 536}},
		{point{74, 956}, point{981, 49}},
		{point{733, 913}, point{496, 676}},
		{point{32, 776}, point{32, 556}},
		{point{676, 890}, point{676, 709}},
		{point{18, 568}, point{18, 725}},
		{point{708, 531}, point{167, 531}},
		{point{240, 33}, point{240, 356}},
		{point{804, 380}, point{804, 735}},
		{point{925, 85}, point{925, 74}},
		{point{692, 287}, point{692, 526}},
		{point{570, 451}, point{570, 379}},
		{point{521, 13}, point{521, 226}},
		{point{249, 598}, point{677, 170}},
		{point{739, 804}, point{353, 418}},
		{point{785, 138}, point{375, 548}},
		{point{913, 161}, point{286, 788}},
		{point{179, 914}, point{158, 935}},
		{point{208, 250}, point{208, 29}},
		{point{953, 506}, point{953, 84}},
		{point{904, 328}, point{347, 885}},
		{point{743, 311}, point{980, 311}},
		{point{897, 988}, point{981, 988}},
		{point{269, 660}, point{534, 660}},
		{point{180, 443}, point{125, 443}},
		{point{101, 713}, point{513, 713}},
		{point{491, 523}, point{265, 749}},
		{point{33, 165}, point{140, 272}},
		{point{813, 544}, point{305, 36}},
		{point{735, 556}, point{605, 556}},
		{point{227, 252}, point{942, 967}},
		{point{67, 69}, point{942, 944}},
		{point{623, 118}, point{920, 118}},
		{point{681, 764}, point{745, 764}},
		{point{182, 499}, point{884, 499}},
		{point{51, 921}, point{787, 185}},
		{point{632, 232}, point{577, 232}},
		{point{887, 112}, point{12, 987}},
		{point{943, 881}, point{255, 881}},
		{point{655, 293}, point{163, 785}},
		{point{453, 96}, point{517, 160}},
		{point{581, 870}, point{581, 481}},
		{point{473, 936}, point{638, 771}},
		{point{264, 719}, point{137, 592}},
		{point{659, 336}, point{659, 75}},
		{point{781, 29}, point{365, 445}},
		{point{93, 677}, point{240, 530}},
		{point{798, 566}, point{594, 566}},
		{point{108, 244}, point{108, 634}},
		{point{69, 879}, point{592, 879}},
		{point{763, 700}, point{763, 351}},
		{point{227, 108}, point{916, 797}},
		{point{149, 328}, point{149, 641}},
		{point{809, 586}, point{776, 619}},
		{point{177, 308}, point{319, 166}},
		{point{970, 880}, point{156, 66}},
		{point{427, 765}, point{427, 433}},
		{point{483, 771}, point{39, 771}},
		{point{813, 895}, point{416, 895}},
		{point{808, 151}, point{808, 823}},
		{point{412, 797}, point{412, 673}},
		{point{368, 965}, point{368, 191}},
		{point{269, 776}, point{409, 636}},
		{point{509, 895}, point{509, 281}},
		{point{873, 863}, point{873, 746}},
		{point{622, 788}, point{623, 788}},
		{point{274, 212}, point{301, 212}},
		{point{335, 804}, point{172, 641}},
		{point{661, 864}, point{620, 864}},
		{point{506, 951}, point{178, 951}},
		{point{611, 609}, point{611, 508}},
		{point{255, 79}, point{255, 479}},
		{point{443, 788}, point{441, 790}},
		{point{548, 738}, point{616, 738}},
		{point{535, 407}, point{535, 845}},
		{point{365, 469}, point{622, 469}},
		{point{541, 670}, point{534, 670}},
		{point{49, 52}, point{922, 925}},
		{point{433, 412}, point{917, 412}},
		{point{475, 494}, point{78, 891}},
		{point{377, 47}, point{377, 384}},
		{point{401, 22}, point{479, 22}},
		{point{642, 889}, point{642, 988}},
		{point{788, 334}, point{788, 375}},
		{point{533, 327}, point{217, 327}},
		{point{529, 934}, point{529, 368}},
		{point{917, 491}, point{280, 491}},
		{point{922, 510}, point{922, 483}},
		{point{695, 104}, point{695, 783}},
		{point{884, 197}, point{192, 889}},
		{point{175, 956}, point{956, 175}},
		{point{384, 711}, point{384, 181}},
		{point{28, 931}, point{28, 732}},
		{point{768, 522}, point{762, 522}},
		{point{816, 964}, point{541, 689}},
		{point{192, 423}, point{668, 899}},
		{point{742, 133}, point{139, 133}},
		{point{829, 708}, point{915, 708}},
		{point{927, 989}, point{72, 134}},
		{point{819, 851}, point{819, 470}},
		{point{326, 699}, point{112, 699}},
		{point{166, 82}, point{370, 286}},
		{point{801, 621}, point{219, 39}},
		{point{392, 332}, point{392, 375}},
		{point{170, 526}, point{549, 526}},
		{point{296, 907}, point{296, 378}},
		{point{912, 456}, point{912, 814}},
		{point{869, 346}, point{648, 346}},
		{point{545, 224}, point{622, 224}},
		{point{626, 657}, point{221, 657}},
		{point{829, 313}, point{829, 626}},
		{point{565, 758}, point{565, 298}},
		{point{113, 810}, point{89, 834}},
		{point{729, 418}, point{679, 418}},
		{point{626, 794}, point{805, 794}},
		{point{811, 568}, point{811, 564}},
		{point{902, 600}, point{735, 600}},
		{point{776, 519}, point{448, 519}},
		{point{874, 890}, point{321, 337}},
		{point{479, 96}, point{479, 153}},
		{point{331, 396}, point{158, 396}},
		{point{420, 413}, point{73, 413}},
		{point{845, 949}, point{49, 153}},
		{point{55, 624}, point{55, 413}},
		{point{349, 761}, point{116, 761}},
		{point{429, 252}, point{429, 384}},
		{point{310, 340}, point{208, 340}},
		{point{208, 692}, point{726, 174}},
		{point{648, 66}, point{648, 685}},
		{point{567, 580}, point{858, 871}},
		{point{747, 57}, point{430, 57}},
		{point{97, 951}, point{850, 198}},
		{point{420, 670}, point{420, 518}},
		{point{583, 308}, point{367, 308}},
		{point{240, 983}, point{219, 983}},
		{point{404, 901}, point{135, 901}},
		{point{118, 126}, point{118, 166}},
		{point{981, 316}, point{727, 62}},
		{point{512, 262}, point{512, 761}},
		{point{445, 758}, point{747, 758}},
		{point{320, 505}, point{252, 437}},
		{point{739, 379}, point{556, 562}},
		{point{509, 791}, point{587, 713}},
		{point{747, 271}, point{196, 822}},
		{point{70, 132}, point{906, 968}},
		{point{90, 580}, point{90, 881}},
		{point{273, 529}, point{273, 886}},
		{point{786, 443}, point{830, 443}},
		{point{116, 903}, point{116, 130}},
		{point{822, 597}, point{822, 430}},
		{point{585, 875}, point{736, 875}},
		{point{470, 649}, point{740, 649}},
		{point{814, 533}, point{814, 20}},
		{point{235, 468}, point{921, 468}},
		{point{413, 262}, point{413, 37}},
		{point{963, 761}, point{963, 389}},
		{point{919, 445}, point{919, 112}},
		{point{788, 800}, point{513, 525}},
		{point{770, 783}, point{262, 275}},
		{point{601, 330}, point{504, 330}},
		{point{882, 668}, point{760, 790}},
		{point{450, 431}, point{675, 431}},
		{point{599, 400}, point{344, 400}},
		{point{887, 19}, point{17, 889}},
		{point{420, 86}, point{420, 202}},
		{point{95, 871}, point{726, 240}},
		{point{337, 558}, point{571, 558}},
		{point{493, 555}, point{934, 114}},
		{point{804, 112}, point{88, 828}},
		{point{785, 673}, point{785, 304}},
		{point{27, 285}, point{27, 865}},
		{point{200, 379}, point{238, 341}},
		{point{303, 383}, point{140, 546}},
		{point{245, 757}, point{33, 757}},
		{point{960, 526}, point{516, 526}},
		{point{303, 933}, point{986, 250}},
		{point{571, 848}, point{571, 781}},
		{point{812, 804}, point{28, 804}},
		{point{752, 877}, point{752, 278}},
		{point{855, 847}, point{63, 55}},
		{point{720, 740}, point{645, 740}},
		{point{872, 789}, point{206, 123}},
		{point{189, 618}, point{189, 567}},
		{point{952, 141}, point{952, 235}},
		{point{81, 523}, point{81, 577}},
		{point{859, 139}, point{449, 139}},
		{point{977, 978}, point{551, 978}},
		{point{30, 15}, point{855, 840}},
		{point{344, 65}, point{842, 65}},
		{point{435, 414}, point{318, 414}},
		{point{324, 813}, point{902, 235}},
		{point{18, 20}, point{977, 979}},
		{point{606, 373}, point{827, 152}},
		{point{678, 881}, point{277, 480}},
		{point{81, 196}, point{81, 176}},
		{point{652, 528}, point{546, 528}},
		{point{673, 310}, point{673, 503}},
		{point{413, 494}, point{413, 859}},
		{point{393, 310}, point{139, 310}},
		{point{485, 157}, point{368, 157}},
		{point{152, 611}, point{362, 401}},
		{point{929, 346}, point{929, 220}},
		{point{577, 102}, point{577, 589}},
		{point{20, 843}, point{882, 843}},
		{point{261, 69}, point{988, 796}},
		{point{417, 37}, point{90, 37}},
		{point{368, 469}, point{149, 250}},
		{point{651, 785}, point{179, 313}},
		{point{953, 759}, point{953, 31}},
		{point{534, 215}, point{534, 199}},
		{point{375, 908}, point{375, 11}},
		{point{408, 571}, point{408, 224}},
		{point{146, 88}, point{146, 16}},
		{point{923, 843}, point{923, 11}},
		{point{885, 605}, point{908, 605}},
		{point{383, 288}, point{698, 288}},
		{point{955, 409}, point{379, 409}},
		{point{10, 579}, point{10, 143}},
		{point{487, 277}, point{918, 708}},
		{point{240, 800}, point{508, 532}},
		{point{655, 121}, point{655, 956}},
		{point{277, 208}, point{277, 395}},
		{point{242, 430}, point{654, 430}},
		{point{518, 982}, point{143, 982}},
		{point{626, 758}, point{626, 125}},
		{point{90, 67}, point{963, 940}},
		{point{57, 11}, point{979, 933}},
		{point{777, 29}, point{777, 669}},
		{point{747, 672}, point{516, 672}},
		{point{915, 878}, point{915, 964}},
		{point{678, 941}, point{678, 738}},
		{point{967, 645}, point{967, 814}},
		{point{356, 293}, point{356, 143}},
		{point{282, 710}, point{585, 710}},
		{point{210, 126}, point{210, 836}},
		{point{77, 122}, point{270, 122}},
		{point{976, 956}, point{655, 956}},
		{point{129, 394}, point{697, 962}},
		{point{621, 94}, point{621, 265}},
		{point{431, 619}, point{431, 349}},
		{point{551, 990}, point{551, 666}},
		{point{379, 681}, point{109, 411}},
		{point{713, 94}, point{832, 94}},
		{point{800, 776}, point{437, 413}},
		{point{261, 517}, point{261, 551}},
		{point{245, 963}, point{245, 807}},
		{point{890, 888}, point{890, 269}},
		{point{591, 455}, point{591, 494}},
		{point{243, 756}, point{166, 756}},
		{point{372, 303}, point{805, 303}},
		{point{22, 459}, point{22, 138}},
		{point{44, 248}, point{44, 689}},
		{point{219, 629}, point{219, 646}},
		{point{444, 633}, point{777, 300}},
		{point{679, 102}, point{679, 945}},
		{point{950, 123}, point{250, 823}},
		{point{485, 264}, point{485, 963}},
		{point{313, 355}, point{445, 355}},
		{point{791, 823}, point{895, 823}},
		{point{914, 193}, point{495, 612}},
		{point{174, 932}, point{351, 932}},
		{point{509, 989}, point{813, 685}},
		{point{830, 559}, point{830, 65}},
		{point{762, 779}, point{84, 101}},
		{point{66, 432}, point{129, 432}},
		{point{224, 705}, point{777, 152}},
		{point{377, 280}, point{93, 280}},
		{point{799, 313}, point{713, 313}},
		{point{972, 496}, point{972, 250}},
		{point{321, 978}, point{488, 978}},
		{point{74, 227}, point{682, 835}},
		{point{434, 827}, point{122, 827}},
		{point{206, 329}, point{536, 659}},
		{point{591, 355}, point{591, 967}},
		{point{838, 436}, point{892, 382}},
		{point{62, 889}, point{878, 73}},
		{point{146, 948}, point{530, 564}},
		{point{308, 510}, point{308, 783}},
		{point{158, 355}, point{158, 469}},
		{point{375, 239}, point{375, 961}},
		{point{754, 281}, point{774, 281}},
		{point{818, 554}, point{818, 922}},
		{point{204, 38}, point{344, 38}},
		{point{689, 283}, point{587, 283}},
		{point{642, 820}, point{862, 600}},
		{point{865, 821}, point{865, 335}},
		{point{51, 870}, point{51, 117}},
		{point{628, 769}, point{129, 270}},
		{point{860, 300}, point{860, 666}},
		{point{626, 47}, point{626, 915}},
		{point{222, 733}, point{449, 733}},
		{point{284, 842}, point{284, 717}},
		{point{965, 834}, point{697, 834}},
		{point{159, 826}, point{159, 732}},
		{point{710, 679}, point{907, 482}},
		{point{356, 355}, point{36, 355}},
		{point{736, 289}, point{577, 130}},
		{point{589, 319}, point{96, 319}},
		{point{967, 20}, point{31, 956}},
		{point{138, 871}, point{138, 446}},
		{point{272, 174}, point{272, 753}},
		{point{958, 228}, point{958, 842}},
		{point{159, 115}, point{576, 532}},
		{point{244, 986}, point{244, 662}},
		{point{191, 160}, point{191, 376}},
		{point{979, 987}, point{11, 19}},
		{point{588, 700}, point{588, 819}},
		{point{76, 933}, point{404, 605}},
		{point{142, 656}, point{142, 906}},
		{point{390, 429}, point{859, 898}},
		{point{802, 499}, point{380, 499}},
		{point{18, 988}, point{969, 37}},
		{point{61, 154}, point{61, 569}},
		{point{803, 47}, point{803, 963}},
		{point{148, 597}, point{148, 454}},
		{point{368, 221}, point{174, 27}},
		{point{416, 351}, point{510, 351}},
		{point{27, 979}, point{950, 979}},
		{point{861, 937}, point{103, 937}},
		{point{49, 889}, point{876, 62}},
		{point{429, 737}, point{555, 737}},
		{point{940, 704}, point{940, 663}},
		{point{644, 379}, point{661, 379}},
		{point{956, 589}, point{956, 423}},
		{point{836, 119}, point{237, 119}},
		{point{841, 211}, point{154, 898}},
		{point{743, 438}, point{743, 370}},
		{point{691, 155}, point{990, 155}},
		{point{531, 768}, point{56, 293}},
		{point{853, 103}, point{110, 846}},
		{point{411, 673}, point{411, 126}},
		{point{824, 876}, point{406, 876}},
		{point{516, 639}, point{98, 639}},
		{point{583, 353}, point{901, 353}},
		{point{64, 898}, point{858, 898}},
		{point{874, 233}, point{874, 551}},
		{point{294, 830}, point{294, 244}},
		{point{921, 27}, point{86, 862}},
		{point{695, 747}, point{56, 108}},
		{point{442, 552}, point{442, 624}},
		{point{731, 431}, point{905, 431}},
		{point{320, 56}, point{139, 56}},
		{point{450, 96}, point{510, 156}},
		{point{628, 803}, point{65, 803}},
		{point{102, 855}, point{782, 175}},
		{point{399, 877}, point{940, 877}},
		{point{766, 664}, point{757, 664}},
		{point{705, 455}, point{407, 753}},
		{point{924, 458}, point{408, 458}},
		{point{302, 272}, point{536, 38}},
		{point{269, 274}, point{269, 196}},
		{point{368, 124}, point{935, 124}},
		{point{151, 686}, point{151, 629}},
		{point{171, 857}, point{171, 62}},
		{point{81, 895}, point{953, 23}},
		{point{150, 95}, point{150, 955}},
		{point{325, 16}, point{102, 16}},
		{point{148, 775}, point{144, 775}},
		{point{742, 449}, point{742, 715}},
		{point{706, 284}, point{706, 865}},
		{point{857, 309}, point{754, 206}},
		{point{866, 818}, point{140, 92}},
		{point{13, 768}, point{956, 768}},
		{point{775, 514}, point{132, 514}},
		{point{868, 407}, point{868, 516}},
		{point{13, 11}, point{989, 987}},
		{point{428, 296}, point{19, 705}},
		{point{544, 665}, point{376, 497}},
		{point{371, 619}, point{371, 512}},
		{point{657, 551}, point{907, 551}},
		{point{571, 324}, point{765, 324}},
		{point{555, 539}, point{295, 799}},
		{point{854, 956}, point{854, 85}},
		{point{152, 255}, point{719, 822}},
		{point{162, 519}, point{868, 519}},
		{point{276, 235}, point{276, 693}},
		{point{58, 870}, point{739, 189}},
		{point{731, 229}, point{613, 347}},
		{point{469, 378}, point{469, 44}},
		{point{30, 909}, point{929, 10}},
		{point{474, 423}, point{256, 205}},
		{point{810, 263}, point{810, 545}},
		{point{230, 244}, point{741, 244}},
		{point{892, 709}, point{394, 709}},
		{point{141, 87}, point{912, 858}},
		{point{280, 820}, point{907, 193}},
		{point{935, 897}, point{101, 63}},
		{point{283, 15}, point{283, 170}},
		{point{122, 749}, point{761, 110}},
		{point{265, 475}, point{265, 390}},
		{point{286, 464}, point{891, 464}},
		{point{793, 819}, point{793, 659}},
		{point{372, 96}, point{372, 591}},
		{point{436, 732}, point{436, 527}},
		{point{48, 301}, point{429, 301}},
		{point{696, 298}, point{696, 294}},
		{point{757, 553}, point{438, 872}},
		{point{635, 856}, point{239, 856}},
		{point{312, 987}, point{989, 987}},
		{point{388, 599}, point{546, 599}},
		{point{235, 391}, point{581, 391}},
		{point{696, 340}, point{696, 215}},
		{point{852, 485}, point{105, 485}},
		{point{931, 586}, point{803, 714}},
		{point{956, 164}, point{198, 922}},
		{point{305, 274}, point{305, 979}},
		{point{363, 774}, point{363, 81}},
		{point{431, 957}, point{431, 460}},
		{point{118, 439}, point{512, 833}},
		{point{193, 318}, point{494, 619}},
		{point{796, 819}, point{332, 819}},
		{point{911, 764}, point{865, 810}},
		{point{282, 482}, point{128, 482}},
		{point{551, 244}, point{694, 244}},
		{point{947, 142}, point{45, 142}},
		{point{653, 57}, point{131, 57}},
		{point{822, 720}, point{721, 720}},
		{point{861, 871}, point{125, 135}},
		{point{950, 698}, point{597, 698}},
	},
}

func (l *lines) filter_angled() {
	to_remove := []int{}
	for index, line := range l.data {
		if line.start.x != line.end.x && line.start.y != line.end.y {
			to_remove = append(to_remove, index)
		}
	}

	for i := len(to_remove) - 1; i >= 0; i-- {
		l.data = append(l.data[:to_remove[i]], l.data[to_remove[i]+1:]...)
	}
}

func (l *lines) get_range() (int, int) {
	x_max := 0
	y_max := 0

	for _, line := range l.data {
		if line.start.x > x_max {
			x_max = line.start.x
		}
		if line.end.x > x_max {
			x_max = line.end.x
		}
		if line.start.y > y_max {
			y_max = line.start.y
		}
		if line.end.y > y_max {
			y_max = line.end.y
		}
	}
	return x_max, y_max
}

func create_map(x int, y int) vent_map {
	v_map := vent_map{
		x + 1,
		y + 1,
		make([][]int, x+1),
	}
	for i := range v_map.data {
		v_map.data[i] = make([]int, y+1)
	}
	return v_map
}

func (v_map *vent_map) populate(l lines) {
	for _, line := range l.data {
		x_delta := line.end.x - line.start.x
		y_delta := line.end.y - line.start.y

		x_step := 0
		if x_delta > 0 {
			x_step = 1
		} else if x_delta < 0 {
			x_step = -1
		}

		y_step := 0
		if y_delta > 0 {
			y_step = 1
		} else if y_delta < 0 {
			y_step = -1
		}

		x := line.start.x
		y := line.start.y
		for {
			v_map.data[x][y]++
			if x == line.end.x && y == line.end.y {
				break
			}
			x += x_step
			y += y_step

		}

	}
}

func (v_map vent_map) count_overlaps() int {
	count := 0
	for x := 0; x < v_map.x_size; x++ {
		for y := 0; y < v_map.y_size; y++ {
			if v_map.data[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func Day05A() int {

	data := day05_data

	data.filter_angled()
	v_map := create_map(data.get_range())
	v_map.populate(data)
	return v_map.count_overlaps()
}

func Day05B() int {

	data := day05_data

	v_map := create_map(data.get_range())
	v_map.populate(data)
	return v_map.count_overlaps()
}
