package days

import (
	"strings"
)

type display struct {
	signal_pattern  [10]string
	display_pattern [4]string
}

type display_results struct {
	input         display
	signal_map    [10]int
	display_value [4]int
	segment_map   [7]string
}

var day08_sample = []display{
	{[10]string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, [4]string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
	{[10]string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, [4]string{"fcgedb", "cgb", "dgebacf", "gc"}},
	{[10]string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, [4]string{"cg", "cg", "fdcagb", "cbg"}},
	{[10]string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, [4]string{"efabcd", "cedba", "gadfec", "cb"}},
	{[10]string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, [4]string{"gecf", "egdcabf", "bgf", "bfgea"}},
	{[10]string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, [4]string{"gebdcfa", "ecba", "ca", "fadegcb"}},
	{[10]string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, [4]string{"cefg", "dcbef", "fcge", "gbcadfe"}},
	{[10]string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, [4]string{"ed", "bcgafe", "cdgba", "cbgef"}},
	{[10]string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, [4]string{"gbdfcae", "bgc", "cg", "cgb"}},
	{[10]string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, [4]string{"fgae", "cfgab", "fg", "bagce"}},
}
var day08_data = []display{
	{[10]string{"dbc", "gfecab", "afcdg", "dfebcag", "bd", "dgbe", "bcaeg", "dcefab", "ecgadb", "agcbd"}, [4]string{"acdgb", "gbcda", "gdecfba", "bacge"}},
	{[10]string{"bacdegf", "aefbdc", "ebf", "fdbcag", "edbfa", "gdaeb", "acfdb", "cdegbf", "face", "fe"}, [4]string{"ebf", "ecdabf", "fcbad", "afcdbg"}},
	{[10]string{"cabgde", "gd", "becgfd", "dgfe", "cebgf", "gfdeacb", "fdbac", "bcgaef", "bgdfc", "gdc"}, [4]string{"cbfad", "dg", "dgef", "ecfbdg"}},
	{[10]string{"adecgf", "egfdc", "cgeadb", "adbfce", "dafg", "bcfeg", "dge", "dg", "fadec", "dcbegaf"}, [4]string{"dg", "dcegf", "cdgafe", "gbacfed"}},
	{[10]string{"badeg", "gdbfcea", "acgef", "fgdc", "bgcaef", "cd", "ebcfda", "gadce", "edcgaf", "edc"}, [4]string{"dce", "efcga", "dbgea", "dc"}},
	{[10]string{"bgfec", "dbfgec", "ebgafc", "afegcdb", "fca", "af", "cgdabf", "aegf", "badce", "cfabe"}, [4]string{"caf", "ebacf", "bafgec", "cbfeag"}},
	{[10]string{"begacd", "bd", "cfadbeg", "fbecg", "begfd", "bgd", "cdfeag", "gbaedf", "fdba", "agefd"}, [4]string{"cbgfe", "fdab", "efbdg", "bd"}},
	{[10]string{"bg", "gfdebc", "beg", "gabd", "fcbae", "fgaeb", "afegd", "gafdce", "dacbfge", "efgbda"}, [4]string{"gdba", "dfgae", "dgba", "gb"}},
	{[10]string{"cfd", "edcfbg", "abfcg", "gfdcb", "gbcefda", "dgcbe", "defg", "df", "acebfd", "caedbg"}, [4]string{"agedcbf", "dgbfaec", "gbfecd", "cdf"}},
	{[10]string{"bdcag", "acgebfd", "eac", "ae", "bade", "abcdgf", "bdacge", "bcgea", "aedgcf", "efcbg"}, [4]string{"cfdagb", "degacb", "eac", "edab"}},
	{[10]string{"ga", "ebcadgf", "gdaf", "bga", "fgecb", "dfabcg", "adcfb", "gadecb", "fecdab", "gbacf"}, [4]string{"fgda", "gdaf", "agb", "gfda"}},
	{[10]string{"gbcdaf", "fabdceg", "gbe", "ebfgad", "bfdga", "eg", "cefabg", "fdge", "adegb", "cedba"}, [4]string{"ge", "fcaedbg", "ebg", "aecbdfg"}},
	{[10]string{"ab", "bdeag", "gbdefc", "cgafdb", "fgedb", "agb", "befa", "bagfdec", "ecgda", "agfedb"}, [4]string{"ab", "aefb", "abg", "ab"}},
	{[10]string{"fgceab", "fcgdb", "fdcebg", "cbf", "gafbd", "cegfd", "bced", "cb", "cadfge", "cbdgafe"}, [4]string{"ecbagf", "cbfdeg", "cb", "bc"}},
	{[10]string{"fgb", "eafdgcb", "cefg", "ebdgac", "fcbged", "fbade", "fg", "gbfed", "gbecd", "fcabdg"}, [4]string{"fg", "eadfcgb", "fgcbda", "gedcb"}},
	{[10]string{"ecgfda", "faebc", "cb", "fadeb", "adecbg", "edfgabc", "bfgc", "cab", "gecaf", "ecbfag"}, [4]string{"fcgb", "bca", "gbcf", "cgbf"}},
	{[10]string{"gfaed", "cg", "cgab", "abdcf", "fabcgd", "gcfda", "fgc", "fdgacbe", "cbdaef", "fcebdg"}, [4]string{"fgc", "cfdageb", "cbfgad", "degfa"}},
	{[10]string{"fbacd", "gedfac", "fadegbc", "gbf", "degb", "ecgdf", "gb", "bdecfg", "gcdbf", "efbcag"}, [4]string{"cgedfb", "gb", "bcdfa", "cfdagbe"}},
	{[10]string{"dceg", "deacbg", "acfbeg", "egdafbc", "ebfda", "dc", "bcd", "cgabe", "gbdafc", "dbace"}, [4]string{"cged", "fabegc", "febad", "ebgcfa"}},
	{[10]string{"gbdcae", "cfgade", "dc", "dgbface", "fdage", "cbgfa", "adc", "gdafc", "dcfe", "gbdaef"}, [4]string{"fedc", "adgcf", "fcde", "dgacbfe"}},
	{[10]string{"afdgbe", "cgfabe", "dcgeab", "aef", "feacb", "fbgcade", "bgace", "cgef", "bcafd", "fe"}, [4]string{"ebacf", "abfdegc", "dbfac", "facgeb"}},
	{[10]string{"fbeg", "dgfecba", "cfgabd", "eafcb", "gb", "edacg", "gecfab", "aebcdf", "ebcag", "gcb"}, [4]string{"cfadeb", "gecda", "bg", "gb"}},
	{[10]string{"bfcaeg", "gd", "fbagc", "fgad", "fgcaedb", "fdgbac", "cdaeb", "dgfecb", "cagdb", "bdg"}, [4]string{"cbfgda", "dg", "afcedgb", "egfacb"}},
	{[10]string{"gd", "dcafgb", "dacfbe", "edbgacf", "gfbdc", "fbceg", "dbg", "dfag", "bdfca", "begdac"}, [4]string{"aecfdb", "eagdcfb", "fdcbgea", "fgdbc"}},
	{[10]string{"gf", "dfgc", "ebfgac", "bfdace", "gabfd", "adgfbc", "acbfd", "ecfagbd", "bgf", "aedgb"}, [4]string{"dfgc", "fdcg", "gfb", "gbf"}},
	{[10]string{"gbdae", "ebadf", "dbcfag", "acfdb", "dfec", "feacdbg", "geafcb", "dcbafe", "fe", "fae"}, [4]string{"efa", "bdcfa", "eaf", "debacgf"}},
	{[10]string{"deafgc", "fbgedca", "gfd", "dcga", "dcgbef", "aecfg", "dg", "agfebc", "adgfe", "dfabe"}, [4]string{"dgf", "badgcef", "fgace", "cabefgd"}},
	{[10]string{"badcg", "bagdce", "bdafgce", "efgbc", "df", "cbgfda", "cfdbg", "edgfac", "fgd", "adbf"}, [4]string{"dfcbg", "df", "cgbdae", "bdfgc"}},
	{[10]string{"dgabe", "beagc", "gbd", "gfcbdae", "gd", "bdegfa", "bgfedc", "fbcdae", "bafde", "fagd"}, [4]string{"bcdfge", "bgd", "bcage", "dfga"}},
	{[10]string{"adgf", "adgfce", "aecfg", "cdega", "dbgaecf", "abcgfe", "egd", "egcdfb", "dg", "deabc"}, [4]string{"cdgae", "dafg", "dfag", "afdg"}},
	{[10]string{"fbadec", "afcbg", "begac", "be", "ebcgda", "dceag", "edgb", "efdcga", "ebdacgf", "cbe"}, [4]string{"adfcebg", "ecafbd", "debgac", "edcafg"}},
	{[10]string{"cbd", "dcga", "cagbf", "dcafb", "cd", "cdebfg", "dgafcb", "baefd", "cbfaeg", "cfdgeab"}, [4]string{"fcdab", "geacbf", "gfbdce", "dgcafb"}},
	{[10]string{"becd", "abfdc", "cgbaf", "efcda", "bd", "ebdcaf", "gdfaceb", "bdf", "bdgeaf", "cedfga"}, [4]string{"ebdc", "fbd", "bdf", "fbgdae"}},
	{[10]string{"gefacd", "adcg", "fedbc", "cdgfe", "egd", "dg", "dbfega", "cgbfeda", "aecfg", "bcefga"}, [4]string{"defbcga", "efcga", "afgced", "gdecf"}},
	{[10]string{"gfacbd", "fbcgde", "cefba", "efdcga", "bfgcead", "egcdf", "gfb", "gcfbe", "bg", "edgb"}, [4]string{"gb", "fdegc", "gfb", "fgb"}},
	{[10]string{"ag", "agfdcb", "egfad", "fag", "adefb", "fdgec", "bdfgcae", "acfegd", "gebfdc", "aecg"}, [4]string{"fga", "ag", "egafd", "fcgebd"}},
	{[10]string{"cadgbf", "cd", "gefbac", "cdfba", "fbead", "cedagb", "gdcf", "adc", "gacbf", "fgebcda"}, [4]string{"gcbafd", "bdcaf", "gfcd", "abdfe"}},
	{[10]string{"fgebda", "fd", "afceb", "fcdeb", "fcad", "cgbed", "gabfcde", "dcabfe", "def", "befagc"}, [4]string{"eacfb", "efd", "efagcb", "def"}},
	{[10]string{"beaf", "gfcaebd", "fdeag", "dae", "acdgeb", "gcdfa", "dgabef", "gbedf", "ea", "dcbgef"}, [4]string{"ae", "aed", "ae", "ade"}},
	{[10]string{"gb", "dgeca", "gecab", "gcdfaeb", "ceafb", "bgc", "fagb", "adfceb", "baegcf", "gebfdc"}, [4]string{"fgba", "bg", "fecgab", "bg"}},
	{[10]string{"daf", "gbfda", "afcedg", "geadfb", "fbcdg", "fgaebc", "ad", "dbagefc", "dbae", "aebgf"}, [4]string{"aebd", "eadcbgf", "dfa", "beda"}},
	{[10]string{"dcbaef", "cabfd", "efgcba", "cbdgfa", "aedfg", "abecfdg", "ace", "cbde", "dfcae", "ce"}, [4]string{"cbefag", "ebcd", "cdabgf", "cbfda"}},
	{[10]string{"bcgefa", "bgfcead", "cd", "fcbadg", "bdcga", "dac", "gbaed", "gafdce", "cfbd", "bcagf"}, [4]string{"cbfag", "acd", "cfbd", "bdfc"}},
	{[10]string{"dgfcb", "aebgd", "gadefc", "bgfdec", "cbafgd", "cgbad", "egbdafc", "ac", "acbf", "cga"}, [4]string{"ca", "ac", "cabf", "abcf"}},
	{[10]string{"cgaed", "gafbdec", "fgaedc", "ba", "baegcd", "edfbg", "agcb", "bcdeaf", "ebdag", "bea"}, [4]string{"gfceda", "ecabfgd", "defgacb", "aeb"}},
	{[10]string{"cbfag", "cgdaf", "dcefab", "bac", "ba", "bdfceg", "cfgeb", "eagb", "cbfgea", "abcfdge"}, [4]string{"ba", "bac", "gabe", "cfbeda"}},
	{[10]string{"efcgd", "ecbafdg", "fgabed", "edb", "ebcg", "adcfb", "be", "efcdgb", "dfcgae", "dbefc"}, [4]string{"dcgef", "ebd", "becg", "gdfaebc"}},
	{[10]string{"fgedabc", "dafeb", "eg", "ged", "aegfbd", "afbecd", "fbgdc", "edfgb", "egaf", "dacegb"}, [4]string{"gaecdb", "cbegad", "fega", "fega"}},
	{[10]string{"bgfcdae", "fbcda", "df", "eafd", "gbcaed", "bdeca", "gfbac", "gfdbce", "bfadce", "bdf"}, [4]string{"afcbdeg", "gabfcde", "deaf", "cbeda"}},
	{[10]string{"fcgebda", "fcga", "bgfcd", "ebdfcg", "degba", "baf", "bafgd", "dfagbc", "bfcade", "af"}, [4]string{"afb", "agfc", "af", "cfegabd"}},
	{[10]string{"cbfg", "bfd", "eabfdg", "facbdg", "edfcgab", "bf", "fcabd", "aedcf", "ebgdac", "badcg"}, [4]string{"eafdc", "bdf", "bf", "fbcg"}},
	{[10]string{"fbdgce", "eagbfd", "dbacfeg", "dagb", "geabf", "cdeaf", "ebd", "bd", "fgbcae", "afebd"}, [4]string{"defac", "bed", "edb", "db"}},
	{[10]string{"aegcd", "bad", "gdaebcf", "cabfg", "db", "ecbd", "gafdce", "gcabd", "gbcade", "faebdg"}, [4]string{"dab", "dba", "dba", "cgfba"}},
	{[10]string{"ecdba", "bedacfg", "cdgeba", "dagcb", "bdge", "gdbfac", "fcadeg", "ced", "de", "acfeb"}, [4]string{"ed", "badfgc", "beacf", "bcgdafe"}},
	{[10]string{"fg", "abdeg", "fgebdca", "agdceb", "gfbd", "gdfae", "fga", "cdefa", "efagbc", "febgad"}, [4]string{"fg", "bgaefc", "gf", "bgafcde"}},
	{[10]string{"beadc", "gcbfed", "fdc", "bacf", "fc", "cbadefg", "feadc", "dgfea", "dcbefa", "dcgeba"}, [4]string{"cdf", "gefda", "dfc", "fbca"}},
	{[10]string{"bd", "bgcedf", "decafg", "gafcdb", "acgdf", "agbce", "bafd", "gaefbcd", "cdb", "acbdg"}, [4]string{"fcdga", "afbd", "cbfdeg", "cafdgb"}},
	{[10]string{"dabe", "bgcda", "egcdb", "fcdage", "dcebga", "bcfge", "dafgebc", "gdafcb", "ed", "ecd"}, [4]string{"ebad", "deba", "de", "gcbde"}},
	{[10]string{"ga", "gadcb", "ebcfdg", "egbcd", "agebdc", "eagb", "afcedg", "dbfca", "dfceagb", "dga"}, [4]string{"dabgc", "ga", "agd", "adbcg"}},
	{[10]string{"cbfagde", "ecdgf", "eabgfc", "fbd", "egbdf", "fageb", "edba", "dacbgf", "bdaegf", "bd"}, [4]string{"beda", "deab", "abgfe", "cegfd"}},
	{[10]string{"gcdfa", "gfea", "ef", "bdacfeg", "aedbc", "efc", "gfebdc", "adgbcf", "fecad", "aefcdg"}, [4]string{"feadc", "fe", "efc", "fgcad"}},
	{[10]string{"gdcae", "fga", "af", "acegfb", "fbda", "gfead", "edfgbc", "agdfecb", "dbgafe", "fdbge"}, [4]string{"fabd", "gfaed", "efgbd", "dagce"}},
	{[10]string{"afdebg", "febc", "cdbae", "abe", "eb", "dbfcea", "dabcfeg", "aecgdf", "dabgc", "dafce"}, [4]string{"cefb", "cedgaf", "eab", "acefd"}},
	{[10]string{"fgc", "dcabg", "gbdefac", "dfeg", "ecfdgb", "bfgcd", "gebcaf", "beafdc", "fg", "edbfc"}, [4]string{"cdfeb", "fcagbe", "fgbdace", "febdc"}},
	{[10]string{"aegfbd", "gacdeb", "fdeacgb", "dea", "dgcefb", "faeg", "bdfac", "dgefb", "ae", "debaf"}, [4]string{"debfg", "dea", "ea", "ead"}},
	{[10]string{"acgedb", "cfgae", "egfcbda", "fdeac", "eacbfd", "fd", "bdagfe", "dabec", "def", "fdcb"}, [4]string{"fed", "fd", "fcbd", "fde"}},
	{[10]string{"dba", "bgcfde", "da", "cdaf", "acfbgde", "fbgdc", "fgadeb", "dcbag", "agbec", "acgfdb"}, [4]string{"gacbdef", "adfgbe", "bgcad", "fcdabge"}},
	{[10]string{"acdfe", "cfgdab", "cdgea", "afgcebd", "dcgbae", "daf", "gfea", "fa", "egdfac", "bedfc"}, [4]string{"fa", "fad", "fa", "cedga"}},
	{[10]string{"caedg", "bafgde", "egcba", "dabegc", "agd", "fecdbga", "bcfeag", "fecdg", "ad", "abdc"}, [4]string{"gad", "cbeag", "dag", "cabd"}},
	{[10]string{"adfebg", "fedcagb", "fagdc", "gda", "eadc", "degcf", "gdafec", "ad", "cdbfge", "facbg"}, [4]string{"da", "aedc", "cbfgdea", "ad"}},
	{[10]string{"ecfdba", "dg", "dgafec", "gcabdfe", "fbadg", "fadcb", "gbdfca", "gfd", "bgefa", "dgbc"}, [4]string{"cgbd", "dbcaf", "cgdb", "dbgcfa"}},
	{[10]string{"ebacg", "eb", "begd", "dcebaf", "cdgea", "gafdce", "ecb", "cagfb", "bfadcge", "egcdba"}, [4]string{"bgdaecf", "bedacg", "dbeg", "cbdgae"}},
	{[10]string{"ebdagcf", "cgbea", "fe", "agefbc", "abgdf", "agfeb", "egafdc", "bcef", "feg", "deabcg"}, [4]string{"cegabd", "ecgafbd", "ef", "fe"}},
	{[10]string{"cd", "cgebdfa", "fgeabd", "cbd", "cabed", "gfbdec", "adcf", "abgce", "eafbd", "fcebad"}, [4]string{"efabcdg", "bcefgd", "baedc", "abfde"}},
	{[10]string{"acfg", "gbacfd", "edgfabc", "bdcaf", "dca", "caebdg", "dafbe", "gbcdf", "ac", "dfbgec"}, [4]string{"bgcfed", "cda", "dbgfc", "acd"}},
	{[10]string{"agdbefc", "abdc", "ecd", "gfcae", "aedbf", "gbcdef", "ebafdg", "ecadbf", "efadc", "cd"}, [4]string{"afdeb", "cde", "dbafce", "bfadge"}},
	{[10]string{"begad", "de", "fged", "gfbea", "dbe", "gbdac", "ebfgdac", "efgabd", "cadfbe", "fegabc"}, [4]string{"de", "dgfe", "cgfaeb", "gadbc"}},
	{[10]string{"gefdc", "gdfb", "gdbcaef", "deb", "ebfcd", "gcebad", "bd", "dacefg", "gdcbef", "fabce"}, [4]string{"gfedabc", "deb", "gcfade", "fdgb"}},
	{[10]string{"bdec", "dagcf", "cefbg", "bfgace", "bgedfc", "edacfbg", "fgdec", "fde", "adgefb", "de"}, [4]string{"dbfaegc", "cedb", "edfabcg", "fcdeg"}},
	{[10]string{"eaf", "gabed", "cabegdf", "fe", "ecbgaf", "fadeg", "gadfc", "fcdega", "cfde", "cbgafd"}, [4]string{"adgef", "fced", "cgdbafe", "fae"}},
	{[10]string{"fcbaedg", "gadce", "abcg", "ceb", "bgedf", "adfbec", "decagf", "bc", "dbceg", "eabdcg"}, [4]string{"ceb", "bgafecd", "cabg", "acefbd"}},
	{[10]string{"dagb", "cfgde", "acgfeb", "fbgac", "fcgdb", "bd", "dbf", "cbfagd", "fcbeda", "bcefgad"}, [4]string{"db", "fegdbca", "dbceafg", "aedbfc"}},
	{[10]string{"edagb", "da", "dab", "fbdge", "cgdfbe", "acegb", "adfe", "abdgfce", "dfgeab", "bfgcad"}, [4]string{"abd", "dab", "gfbead", "eagbc"}},
	{[10]string{"bdfega", "dcbgf", "efcgabd", "ce", "fec", "fcegd", "agce", "fegad", "daebfc", "efdcga"}, [4]string{"ec", "ecf", "dgacef", "bfecad"}},
	{[10]string{"cdabgef", "bgdfe", "aecgb", "gcfa", "fc", "acgebf", "dcaebf", "dcbgae", "cfgeb", "efc"}, [4]string{"fc", "fcga", "bfecga", "faecbdg"}},
	{[10]string{"dfeac", "gdfcea", "fcg", "acdebf", "cdge", "cgefa", "bdcgaf", "cg", "fegbdac", "bfage"}, [4]string{"cfg", "cfg", "gfeca", "fegac"}},
	{[10]string{"bfad", "bcf", "bedcga", "cadeb", "bf", "bceafd", "bcegfd", "dbaegfc", "ecfga", "cfbae"}, [4]string{"efacdb", "afbd", "bcf", "fb"}},
	{[10]string{"aegfcd", "bcdeg", "abcefg", "gcfea", "bcgafed", "bga", "afbe", "bagec", "ab", "fbdgac"}, [4]string{"gceba", "gacbe", "gbacef", "gab"}},
	{[10]string{"dca", "ebda", "bdfaceg", "acgdb", "gbcdea", "begcd", "efbdcg", "efacgd", "bfagc", "ad"}, [4]string{"adcbeg", "ad", "cbfga", "cfbgeda"}},
	{[10]string{"fadebc", "gcfbea", "becgf", "gf", "fgc", "abfg", "ecgbd", "dfecabg", "fcedag", "eafcb"}, [4]string{"ebfgc", "gaedfc", "bgcef", "gedabfc"}},
	{[10]string{"fadg", "eafgb", "gbead", "bdaec", "ecagbf", "dg", "gdb", "agfebd", "aedcgfb", "cfgbde"}, [4]string{"bfecdag", "bgd", "badce", "gd"}},
	{[10]string{"egcd", "efadb", "gacbfd", "bcaeg", "ecadgb", "adc", "bacegf", "adecb", "dc", "agedcfb"}, [4]string{"fcgabe", "degc", "acd", "cd"}},
	{[10]string{"bedfcag", "gc", "fagdec", "gfabce", "bfgec", "acgb", "befac", "gfc", "bgdef", "aecfbd"}, [4]string{"cg", "ebfdg", "cfg", "gc"}},
	{[10]string{"aegbdf", "gfabced", "fadebc", "gfde", "ed", "bcaeg", "fbgda", "eadbg", "fgbacd", "bde"}, [4]string{"dbcefag", "bed", "afcbged", "aegdb"}},
	{[10]string{"agd", "fagec", "cfbgda", "bfceagd", "fgeacb", "cdae", "gdbef", "dagef", "da", "efcdga"}, [4]string{"edbgf", "gda", "gedaf", "cdgfaeb"}},
	{[10]string{"fgaec", "eb", "fcgbd", "agfdeb", "cedb", "fabcdeg", "egb", "febcg", "dfbgec", "cdfgab"}, [4]string{"fgbcde", "dbeafg", "egcfdba", "gcbfd"}},
	{[10]string{"gfe", "fdbae", "ge", "bfgdc", "dgfaec", "cgeb", "dcfgbe", "dabfgc", "fdegb", "eabgcdf"}, [4]string{"gebc", "gceb", "defbcag", "afbcedg"}},
	{[10]string{"dfce", "cgbed", "fbcagd", "bcdgf", "ged", "bdfgae", "ed", "bdfceg", "agebc", "fabdegc"}, [4]string{"fcde", "gde", "gecdbf", "de"}},
	{[10]string{"dba", "dcafbg", "gcdefb", "fbged", "dgeab", "bfegad", "gaecd", "ba", "aebf", "egafdcb"}, [4]string{"beadg", "cfdbage", "bgafde", "acdge"}},
	{[10]string{"bcgdef", "cbdgeaf", "ae", "fbdace", "afbde", "gcdabe", "gdbfa", "edbcf", "fcea", "bea"}, [4]string{"dgafb", "adbfce", "ea", "fbdecga"}},
	{[10]string{"afdg", "fgebc", "fdeacbg", "fd", "gadcb", "fdcgb", "gfacbd", "gbeacd", "acdfeb", "dfc"}, [4]string{"fd", "bdcafeg", "caefbd", "dafg"}},
	{[10]string{"dfebag", "defa", "bdgfa", "beafg", "agbcfd", "eab", "cbagefd", "cfgbe", "ea", "acbgde"}, [4]string{"abdcge", "abe", "ae", "fecbg"}},
	{[10]string{"ef", "bgfdeac", "bdcefg", "cfdba", "ecf", "dgecba", "gfae", "abgefc", "ecbga", "acefb"}, [4]string{"debcag", "afge", "efag", "fe"}},
	{[10]string{"cg", "gbfc", "gabcd", "bgade", "dgc", "gbecafd", "egdfac", "dfbac", "abdcfe", "gabfdc"}, [4]string{"gcbad", "degfabc", "cgafbd", "adbgcf"}},
	{[10]string{"edbcgf", "acgbf", "fbaedc", "cbfae", "gfc", "agcdb", "dagecfb", "fg", "fcgbea", "eagf"}, [4]string{"fg", "adcbef", "fbeac", "fadcgeb"}},
	{[10]string{"deacg", "daeb", "dagbce", "ecdbgf", "edc", "ed", "bgaefc", "gecba", "dagcf", "abdgecf"}, [4]string{"cdaeg", "abcge", "dec", "ecd"}},
	{[10]string{"fcgadb", "caged", "bedc", "cdabeg", "dfaecbg", "befag", "db", "dgb", "eagbd", "gacdef"}, [4]string{"fecagdb", "begad", "dgb", "db"}},
	{[10]string{"fbcdga", "fgcad", "fecdg", "gcedab", "bfca", "adc", "ac", "bfdga", "dbfega", "eafbcdg"}, [4]string{"adc", "cbgade", "fgcadb", "ceabgd"}},
	{[10]string{"cefdbg", "gde", "bgfeac", "ed", "begcd", "cbefg", "bfde", "aedcgf", "cbafegd", "gacbd"}, [4]string{"fegbdac", "gcbefd", "cbgef", "bgedc"}},
	{[10]string{"bfdagc", "afdc", "bagcf", "gbfdc", "begcd", "begdaf", "beafgc", "dbefgac", "gfd", "fd"}, [4]string{"fcda", "df", "dfac", "dcbfg"}},
	{[10]string{"gcbfe", "bfeacg", "cedbf", "ge", "bfcga", "gfe", "gfecda", "ecbgdfa", "bgfadc", "bage"}, [4]string{"feg", "eagb", "gef", "abfgc"}},
	{[10]string{"bgfec", "bac", "abecdfg", "ab", "gdcea", "ebgcfa", "fagb", "aebcg", "caebfd", "gfcbed"}, [4]string{"ab", "ab", "cba", "acgbe"}},
	{[10]string{"gdbcf", "cbedg", "ce", "ceg", "abcgde", "adbeg", "efgcab", "bdfegac", "adec", "adfbeg"}, [4]string{"gaecdfb", "bedga", "cge", "cead"}},
	{[10]string{"eafcb", "cfdbe", "aeb", "acfeg", "ab", "dcgbfe", "adfb", "fgacedb", "ecabgd", "fdbcae"}, [4]string{"fbcde", "aedfcbg", "bea", "edcbf"}},
	{[10]string{"ebdca", "egc", "cgabe", "fbge", "cgbfda", "ge", "bacfg", "geacbf", "afdgec", "gebacfd"}, [4]string{"badce", "cge", "bgfcea", "ge"}},
	{[10]string{"gdfecb", "bdge", "gdf", "dgbcaf", "aefbgc", "bfecg", "fgdec", "dg", "ecdbfga", "deacf"}, [4]string{"dgf", "fedcagb", "gd", "gcedf"}},
	{[10]string{"agdeb", "cfdbag", "gcabd", "ed", "dacegbf", "gebfa", "eagbdc", "acde", "efcgbd", "edb"}, [4]string{"ed", "eagdb", "de", "de"}},
	{[10]string{"feag", "cfgbed", "feacbdg", "cfgade", "efcdg", "fcgdab", "afd", "dfeac", "af", "abedc"}, [4]string{"cbagfde", "dbagfc", "faeg", "dgabcf"}},
	{[10]string{"cdg", "efgbcd", "decgfa", "dcegb", "dgbf", "gd", "ebdcfga", "adebc", "gbcfe", "cefagb"}, [4]string{"bfecg", "dgbf", "gdbf", "dg"}},
	{[10]string{"gab", "dfceag", "gdcb", "bg", "gbead", "agdefbc", "gcabef", "acged", "efdba", "agdceb"}, [4]string{"gdcb", "bgdc", "edbfa", "bag"}},
	{[10]string{"bgafde", "afdgc", "gcafed", "efgdc", "ceda", "gdcebfa", "gdbacf", "de", "gbcef", "edg"}, [4]string{"edg", "gde", "eadfcbg", "cade"}},
	{[10]string{"gdb", "egcfadb", "gdef", "abcfd", "dg", "ebfag", "gadbef", "abdgf", "aefbgc", "cbgaed"}, [4]string{"fedgba", "dbegfac", "bfeagd", "gefcdba"}},
	{[10]string{"gad", "bdcag", "becdg", "gceabd", "dceafbg", "fgdeca", "fcgab", "ad", "baed", "dfcbge"}, [4]string{"cfaegdb", "gda", "ad", "gaedcfb"}},
	{[10]string{"dafbeg", "eadgbfc", "fgd", "bagcf", "gedb", "abdfce", "gafdb", "fcdgea", "gd", "efbda"}, [4]string{"gebd", "fdg", "gaefcd", "deagcf"}},
	{[10]string{"cdbe", "adefb", "ecfgbad", "bfcgad", "bacdf", "agcefd", "efd", "cdeafb", "aegfb", "de"}, [4]string{"gacebdf", "bdce", "fbcagd", "dcefab"}},
	{[10]string{"cb", "acdgbe", "bce", "cdegabf", "bgdefa", "aegdb", "cdbg", "acbde", "fcebga", "fdace"}, [4]string{"ebc", "abfdecg", "ceb", "gbacdfe"}},
	{[10]string{"cg", "fcedg", "edfbag", "ceag", "fdbec", "gdc", "faedcg", "bdcefag", "gbafcd", "dfeag"}, [4]string{"geca", "dgfacb", "cgd", "cg"}},
	{[10]string{"bcafde", "defac", "dcage", "dgfbce", "gafc", "cgaedf", "cgd", "cg", "fgcedab", "gdeab"}, [4]string{"cdg", "ebadg", "aedgcf", "fagc"}},
	{[10]string{"bdec", "dagbf", "ecagfb", "decbfa", "aefdc", "aeb", "dcegfa", "adefcbg", "ebadf", "be"}, [4]string{"cdbe", "gfadb", "eadfc", "dcbe"}},
	{[10]string{"bcgeda", "fc", "badfe", "gbfdaec", "fcge", "cfd", "bedfc", "cedbg", "adfgbc", "dcgebf"}, [4]string{"cf", "ebacgfd", "fgce", "fdbceag"}},
	{[10]string{"gd", "fgbdea", "dbg", "badce", "gdcab", "fcagedb", "fbacg", "beacfg", "cgfd", "dcafbg"}, [4]string{"fcdg", "gdcbfa", "dbg", "gdfc"}},
	{[10]string{"cag", "cdfa", "cgade", "edgcbf", "fcebga", "begfcad", "gfdcae", "egabd", "defgc", "ac"}, [4]string{"ca", "agc", "fgbace", "gac"}},
	{[10]string{"ecabf", "geacfb", "eabdcf", "dec", "cd", "bgecfda", "gfdbce", "edcba", "cdaf", "adgeb"}, [4]string{"ced", "aedfgbc", "dc", "gbeda"}},
	{[10]string{"fbdac", "ea", "cadefg", "fae", "beag", "facbe", "bfgeacd", "cfdegb", "bgfec", "gfecba"}, [4]string{"agbe", "ea", "bafce", "egbdcf"}},
	{[10]string{"bdc", "dbega", "abcdg", "gbcfa", "dc", "gcde", "cgbaefd", "cdfbae", "eagdbf", "ebgadc"}, [4]string{"adefcb", "agcfb", "cdb", "gced"}},
	{[10]string{"ceb", "ce", "debfca", "abgdfe", "bdaeg", "cabegd", "caeg", "dbgec", "dcbfg", "bdecgfa"}, [4]string{"ce", "abegd", "gdfcb", "fcbdgae"}},
	{[10]string{"cafedgb", "dgf", "bcfd", "cgdfbe", "acedfg", "gbfed", "agedb", "aecgfb", "bfgec", "fd"}, [4]string{"fdgbec", "feacdg", "df", "bgcfe"}},
	{[10]string{"ef", "abgfc", "cagfeb", "gdebfc", "fbae", "bcadfg", "defgabc", "feg", "gefac", "gceda"}, [4]string{"degac", "gfbac", "gabfc", "gfcea"}},
	{[10]string{"bfdeca", "bd", "eabfgc", "degbcfa", "cgdb", "gefbc", "dbf", "ecfbgd", "gaedf", "edgfb"}, [4]string{"dbfge", "bdf", "egbfca", "db"}},
	{[10]string{"fac", "cbdfg", "eacfbg", "gaecfdb", "bacfd", "bfeda", "acgd", "ac", "gbfecd", "cgfadb"}, [4]string{"dagc", "ac", "dcga", "bedfa"}},
	{[10]string{"egfbca", "aegcf", "gdcfab", "fdc", "adce", "fcedg", "cd", "cfgead", "fdebg", "cedagbf"}, [4]string{"dcgfe", "fdgbcea", "aedc", "cfebga"}},
	{[10]string{"da", "bfged", "afedcb", "gcebdf", "gefda", "fedbga", "edgbfca", "gdba", "eagfc", "eda"}, [4]string{"ead", "dcfebg", "defcgb", "efgbcd"}},
	{[10]string{"adefgb", "ecgf", "gcedb", "fedacb", "bec", "fbecdg", "cdfaebg", "gedfb", "gcbad", "ce"}, [4]string{"dcbfeag", "cfbged", "efgbad", "bcdfage"}},
	{[10]string{"dfec", "gcd", "agdbe", "bdafgc", "afcbge", "bgecfad", "bgcde", "bgfce", "bgedfc", "cd"}, [4]string{"fbecag", "beagdfc", "cgd", "cedf"}},
	{[10]string{"edbc", "agcfbe", "adefg", "cd", "defca", "gfabdc", "cad", "dgabefc", "dfebac", "bceaf"}, [4]string{"dbce", "ebdc", "febcdga", "cdgbefa"}},
	{[10]string{"ecdb", "eb", "acgebdf", "ebf", "dcfebg", "bgdefa", "fcgbd", "gcdbaf", "feacg", "cfebg"}, [4]string{"bef", "ecfbg", "egdcbfa", "cdgfb"}},
	{[10]string{"ebfcgd", "aefcb", "bfadge", "efd", "ed", "adfcgb", "fbgdc", "bfedagc", "fcebd", "dcge"}, [4]string{"gced", "fbcgd", "edgc", "ed"}},
	{[10]string{"beagf", "cefbag", "ebafc", "gfce", "gdfeab", "fcaedbg", "ec", "dfcba", "ecb", "gbadec"}, [4]string{"ce", "acfdb", "dbgcafe", "ec"}},
	{[10]string{"abegfc", "fbcged", "bgfea", "fecabdg", "gef", "ge", "abdgf", "gaec", "afecb", "cfbdea"}, [4]string{"cbfaedg", "bface", "gfe", "efabgc"}},
	{[10]string{"bdg", "bacd", "edagb", "fbegac", "fgaed", "bdgcae", "cgbea", "dfbegc", "ecgafdb", "bd"}, [4]string{"gdeaf", "cfdgbe", "gbcaed", "bd"}},
	{[10]string{"fbce", "fdegca", "egafb", "aef", "egdcfba", "eabgc", "fe", "ecdabg", "agfbd", "ecgbaf"}, [4]string{"fadgb", "cbage", "bcgdea", "abfgec"}},
	{[10]string{"dbafge", "fedgcab", "efgdbc", "bc", "fbac", "fbdea", "gceda", "bfceda", "bdcea", "ebc"}, [4]string{"eadcb", "fcdeab", "cafb", "bdecgf"}},
	{[10]string{"cbgae", "fec", "aefgcb", "acfeg", "fe", "efgcbd", "bgceafd", "fbea", "cagbed", "dcfga"}, [4]string{"edbagfc", "acefg", "gbdface", "ecf"}},
	{[10]string{"cbfdga", "bae", "ebdf", "fbadge", "afdgb", "gfeac", "acbdge", "bdacfeg", "eb", "agfeb"}, [4]string{"edbf", "bea", "efdcbag", "aeb"}},
	{[10]string{"efdcgab", "bedgaf", "aecbgf", "cgfe", "gfaeb", "cbgea", "ecbfad", "ec", "gdcab", "ecb"}, [4]string{"cgfe", "eabfg", "eabgf", "ec"}},
	{[10]string{"bdefa", "dfebgca", "bfe", "eadfcb", "dcbage", "bfgad", "debca", "cgebaf", "ef", "cdfe"}, [4]string{"gacdbe", "eagcfbd", "afcgebd", "dfgbace"}},
	{[10]string{"gca", "ebdcfa", "fdcgea", "acedb", "bdcgf", "eabg", "gbacde", "bcgdeaf", "ga", "gacdb"}, [4]string{"ebag", "abge", "gac", "aecdb"}},
	{[10]string{"abcgf", "afgcbe", "baeg", "gbc", "ecdbgf", "gafce", "facbd", "cfgade", "bg", "egabdfc"}, [4]string{"aegb", "cgb", "bcdgef", "abcgf"}},
	{[10]string{"ecagbf", "bdfgca", "ecgfd", "gbe", "fecgb", "dafgecb", "bacfg", "bgadec", "efba", "eb"}, [4]string{"geb", "cadfebg", "beaf", "be"}},
	{[10]string{"agcdb", "bgcaf", "bf", "fadebgc", "daecbf", "gbdf", "fbc", "ebgcda", "afegc", "bgcafd"}, [4]string{"facgdeb", "dbfg", "fb", "bf"}},
	{[10]string{"fdbeag", "cgd", "bdefcag", "afegd", "eacgdf", "dc", "dgefc", "ebcagd", "gfebc", "facd"}, [4]string{"geadf", "dc", "dc", "cdg"}},
	{[10]string{"adcgfe", "cgebf", "acdbf", "abcedf", "dgfbaec", "edba", "cabef", "ea", "afe", "dfagcb"}, [4]string{"bfgec", "bfeca", "ae", "deba"}},
	{[10]string{"adegb", "bgcd", "cbdfgea", "dc", "cedba", "cagefd", "dcbgae", "gebdaf", "faebc", "ced"}, [4]string{"egadb", "ced", "dcbg", "dc"}},
	{[10]string{"fdeg", "gd", "fdbaec", "bfacg", "dgcfa", "ecfdga", "fadcegb", "fdcea", "dceabg", "gda"}, [4]string{"efgd", "eacdfb", "bcfadeg", "dag"}},
	{[10]string{"af", "ebgfda", "ecagdb", "febgc", "afe", "dabce", "cfda", "febcda", "baecf", "becafdg"}, [4]string{"fa", "dbfega", "fa", "bfeagd"}},
	{[10]string{"debagfc", "cbdgfe", "dcb", "fdbeg", "cegb", "fbced", "fcead", "cbdfga", "bc", "fabdeg"}, [4]string{"dcb", "efadgb", "dgcfeab", "ebafdg"}},
	{[10]string{"cgfdae", "bdef", "becfg", "gcbefa", "ed", "abcdg", "gedcb", "cgbdef", "bdcgefa", "ecd"}, [4]string{"cde", "dec", "bfed", "bfcgdea"}},
	{[10]string{"dcfegab", "becdg", "gcaedb", "egd", "bcdaef", "gcfdb", "eg", "eafcdg", "cdeab", "ageb"}, [4]string{"ged", "cegafbd", "egd", "eg"}},
	{[10]string{"ba", "eafcd", "fcbegda", "acfgbd", "bedcgf", "ebgcfa", "fabdc", "dcgfb", "bagd", "baf"}, [4]string{"ba", "cafed", "efdbcag", "efdac"}},
	{[10]string{"aebcdgf", "cd", "efbgac", "bfcgd", "cbfga", "gbefd", "bfdcga", "bcad", "geafcd", "dgc"}, [4]string{"fabcged", "fgaceb", "cbda", "dc"}},
	{[10]string{"cgbaedf", "fbceg", "eafbcd", "degcf", "gfdcae", "gdcfa", "dbagcf", "dgae", "edf", "ed"}, [4]string{"fcgeb", "efd", "ed", "gdea"}},
	{[10]string{"ac", "agfdcbe", "gcea", "dgecf", "fgdca", "efgcad", "dfbeac", "fca", "fdgba", "gcebfd"}, [4]string{"agbdf", "ca", "ac", "gbdfa"}},
	{[10]string{"gbedaf", "dbgfc", "cbg", "abdfgc", "geadbc", "fecbd", "ebcgafd", "cafg", "gc", "dafbg"}, [4]string{"cdbfg", "gcfbad", "gcbdea", "fadcegb"}},
	{[10]string{"bagfec", "efg", "bfcdga", "egac", "abfecdg", "afdbge", "dbefc", "ge", "cgbfe", "acfbg"}, [4]string{"fbdeag", "eg", "eg", "cbdfe"}},
	{[10]string{"edbgcf", "dgbac", "bdceag", "dbagf", "cgd", "ecagb", "bceagf", "cd", "adce", "cgbafed"}, [4]string{"aecd", "egdbca", "gacbe", "gcd"}},
	{[10]string{"bgfdc", "gdfbe", "afdbgc", "gc", "fbdac", "bacg", "gbafdec", "cdg", "cbdefa", "defagc"}, [4]string{"egdfb", "gcabdef", "cdbfa", "dgacfb"}},
	{[10]string{"adgcfb", "gcd", "dfecgb", "cgebfa", "bgcaf", "adgcf", "cdbafeg", "dg", "ecafd", "adbg"}, [4]string{"bcgaf", "fadcgeb", "gd", "cgbaedf"}},
	{[10]string{"bafgd", "cfaebg", "cefa", "cfbgead", "bcf", "bgcea", "defgcb", "bgacf", "ebcgda", "cf"}, [4]string{"fcbgae", "adfbg", "cdfgeb", "adgbf"}},
	{[10]string{"becfa", "bcgdf", "gbceadf", "gebcfa", "defa", "adbcge", "de", "dec", "cebdf", "bdefac"}, [4]string{"faed", "fdcagbe", "edc", "bfacge"}},
	{[10]string{"ebfgd", "gabd", "gd", "gfbcaed", "abgfec", "faegb", "dge", "aefdbg", "agdfec", "cfebd"}, [4]string{"faecdg", "cgaedf", "baefgd", "dcfeb"}},
	{[10]string{"ecabd", "ged", "aebgfc", "ecagd", "dg", "fdcegb", "afgce", "dgaf", "dcaegf", "cedbgaf"}, [4]string{"edcga", "gfda", "gd", "dge"}},
	{[10]string{"dfg", "fg", "eafdb", "aegdc", "faged", "gfba", "debafg", "ebgfcad", "gfbdce", "acedbf"}, [4]string{"fedagb", "edbfac", "gf", "fgd"}},
	{[10]string{"dbgca", "fedca", "bfgcad", "adgcf", "agdecb", "dgf", "cefdgba", "bgaedf", "fcgb", "gf"}, [4]string{"acefdgb", "beafdg", "gfd", "cfbg"}},
	{[10]string{"edgacf", "adegcb", "fd", "dbacgfe", "fabd", "dfbce", "egcfb", "edfbac", "cbeda", "fde"}, [4]string{"dafb", "bafd", "acefbd", "fd"}},
	{[10]string{"bcefa", "gbfd", "eadcgb", "bgfdae", "gedcfa", "abefd", "efd", "df", "gedba", "fbdgeac"}, [4]string{"bfdae", "fgeacdb", "bacfe", "gdceba"}},
	{[10]string{"bdcefg", "aefgcd", "gadcb", "gfb", "decbfag", "dfecg", "bfdgc", "defb", "cefagb", "fb"}, [4]string{"defb", "fbcged", "dbfegac", "egfdc"}},
	{[10]string{"fe", "cedgab", "cbfe", "dacbfge", "cedba", "fed", "gafcd", "dfaec", "edbfac", "efdgab"}, [4]string{"abgedf", "efcb", "bedac", "gefbad"}},
	{[10]string{"afcgbe", "cf", "fbc", "cegadb", "bfcdea", "fcgabed", "aebdc", "adfgb", "cedf", "cfadb"}, [4]string{"dafbc", "bfeagc", "bcf", "deacbgf"}},
	{[10]string{"ebcafdg", "ecfbg", "bae", "dgcabf", "bcaedg", "dfbeca", "ae", "ebcag", "gdacb", "dage"}, [4]string{"cadbeg", "gbacd", "bfdegca", "eba"}},
	{[10]string{"cadbe", "efgbc", "feacbg", "dgfecb", "bgd", "dg", "bgfcad", "ebdgc", "dgef", "gbdefca"}, [4]string{"aefcgb", "gd", "bfagced", "dgb"}},
	{[10]string{"edgcfa", "dgbecfa", "bgedac", "dfbaec", "acgdf", "dg", "abfgc", "egdf", "adcef", "gdc"}, [4]string{"cgd", "cbgfa", "eacgdb", "gcafd"}},
	{[10]string{"cad", "cfbead", "abceg", "edfa", "fadbcg", "becad", "edgfcb", "bgfaedc", "cdbfe", "da"}, [4]string{"egcab", "afde", "dcabfe", "ecfabd"}},
	{[10]string{"fgbdc", "be", "egb", "bcdefag", "ceba", "gedbac", "egcda", "dcfgae", "debcg", "gbdaef"}, [4]string{"ceba", "edgac", "eadcfg", "bcdaeg"}},
	{[10]string{"cg", "adcfgbe", "edbcga", "gbead", "cbdfa", "cga", "bgedaf", "faecdg", "dbacg", "cebg"}, [4]string{"gcdabfe", "fdacge", "cga", "bcgad"}},
	{[10]string{"gacf", "gf", "bgeafd", "gef", "gefbc", "efbagc", "bdceag", "fdebc", "acfbedg", "gbaec"}, [4]string{"bcgaedf", "egf", "gcefb", "cafdegb"}},
	{[10]string{"fgbadce", "cdbag", "agcfd", "faec", "acfegd", "fa", "fag", "cfged", "fdabeg", "dbfcge"}, [4]string{"agf", "bagefd", "edgfc", "gdbacef"}},
	{[10]string{"agcdb", "gcadfb", "dgfcb", "gfdbec", "adebg", "adc", "ac", "cfab", "beagdcf", "daecfg"}, [4]string{"ac", "ca", "bedga", "ca"}},
	{[10]string{"dfa", "abdcfg", "agbf", "cfabd", "dacgb", "dcebf", "eadgcb", "fa", "gbdafce", "ecfgad"}, [4]string{"af", "fa", "bcdaf", "bgaf"}},
	{[10]string{"ebdfc", "cgbaed", "eg", "edabfc", "gefb", "adfgc", "gdfec", "deg", "ebdfgc", "debgcaf"}, [4]string{"abfcde", "bfge", "gde", "dge"}},
	{[10]string{"dabfegc", "dfegb", "cbgf", "cf", "fdebgc", "dcbaef", "gcfed", "adecg", "fbdgea", "fce"}, [4]string{"cf", "gbfc", "fgcb", "begcadf"}},
}

func (r *display_results) get_signal(number int) string {
	// fmt.Println("------")
	// fmt.Println(number)
	// fmt.Println(r.signal_map[number])
	return r.input.signal_pattern[r.signal_map[number]]
}

func remove_dupes(a string, b string) string {
	for _, v := range b {
		a = strings.ReplaceAll(a, string(v), "")
	}
	return a
}

func (r *display_results) first_parse() {
	for i, v := range r.input.signal_pattern {
		switch len(v) {
		case 2:
			r.signal_map[1] = i
		case 3:
			r.signal_map[7] = i
		case 4:
			r.signal_map[4] = i
		case 7:
			r.signal_map[8] = i
		}
	}

	for i, v := range r.input.display_pattern {
		switch len(v) {
		case 2:
			r.display_value[i] = 1
		case 3:
			r.display_value[i] = 7
		case 4:
			r.display_value[i] = 4
		case 7:
			r.display_value[i] = 8
		}
	}
}

func (r *display_results) second_parse() {
	one_val := r.get_signal(1)
	four_val := r.get_signal(4)
	for i, v := range r.input.signal_pattern {
		if len(v) == 5 {
			// number 3 will have 3 segments left if both segments in 1 are removed
			v1 := remove_dupes(v, one_val)
			if len(v1) == 3 {
				r.signal_map[3] = i
				// fmt.Println("Found 3")
				continue
			}
			// else
			v2 := remove_dupes(v, four_val)
			if len(v2) == 3 {
				r.signal_map[2] = i
				// fmt.Println("Found 2")
				continue
			} else {
				r.signal_map[5] = i
				// fmt.Println("Found 5")
				continue
			}
		}
	}

	for i, v := range r.input.display_pattern {
		if len(v) == 5 {
			// number 3 will have 3 segments left if both segments in 1 are removed
			v1 := remove_dupes(v, one_val)
			if len(v1) == 3 {
				r.display_value[i] = 3
				continue
			}
			// else
			v2 := remove_dupes(v, four_val)
			if len(v2) == 3 {
				r.display_value[i] = 2
				continue
			} else {
				r.display_value[i] = 5
				continue
			}
		}
	}
}

func (r *display_results) third_parse() {
	four_val := r.get_signal(4)
	five_val := r.get_signal(5)
	seven_val := r.get_signal(7)
	for i, v := range r.input.signal_pattern {
		if len(v) == 6 {
			v1 := remove_dupes(v, seven_val)
			v1 = remove_dupes(v1, four_val)
			if len(v1) == 1 {
				r.signal_map[9] = i
				continue
			}

			v2 := remove_dupes(v, five_val)
			if len(v2) == 2 {
				r.signal_map[0] = i
				continue
			} else {
				r.signal_map[6] = i
				continue
			}

		}
	}

	for i, v := range r.input.display_pattern {
		if len(v) == 6 {
			v1 := remove_dupes(v, seven_val)
			v1 = remove_dupes(v1, four_val)
			if len(v1) == 1 {
				r.display_value[i] = 9
				continue
			}

			v2 := remove_dupes(v, five_val)
			if len(v2) == 2 {
				r.display_value[i] = 0
				continue
			} else {
				r.display_value[i] = 6
				continue
			}

		}
	}
}

func (r *display_results) get_value() uint {
	val := uint(0)
	val += (uint(r.display_value[0]) * uint(1000))
	val += (uint(r.display_value[1]) * uint(100))
	val += (uint(r.display_value[2]) * uint(10))
	val += (uint(r.display_value[3]) * uint(1))
	return val
}

func (d display) create_result() display_results {
	r := display_results{
		d,
		[10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		[4]int{-1, -1, -1, -1},
		[7]string{},
	}

	return r
}

func count_Found(displays []display_results) (sum int) {

	for _, v := range displays {
		for _, r := range v.display_value {
			if r > -1 {
				sum++
			}
		}
	}

	return sum
}

func Day08A() int {

	data := day08_data

	displays := []display_results{}
	for _, v := range data {
		d := v.create_result()
		d.first_parse()
		displays = append(displays, d)
	}

	return count_Found(displays)
}

func Day08B() int {
	data := day08_data
	sum := uint(0)
	for _, v := range data {
		d := v.create_result()
		d.first_parse()
		d.second_parse()
		d.third_parse()
		sum += d.get_value()
	}

	return int(sum)
}
