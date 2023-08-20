package alg


import (
	"context"
	"testing"
)

func TestConvertSentence2ASCII(t *testing.T) {
	cases := []struct{
		sentence string
		asciiStc string
	} {
		{"这是一段中文", "######"},
		{"这是，一段带中文「标点」的中文；", "################"},
		{"中文 is good", "## is good"},
		{"早上，weather is good", "###weather is good"},
		{"weather 是真的 good 啊！", "weather ### good ##"},
		{"Today's 天\nvery 好", "Today's #\nvery #"},
		{"-你好！ \n-how are you", "-### \n-how are you"},
		{"-你好！ \n -how are you", "-### \n -how are you"},
	}

	ctx := context.Background()

	for i, c := range cases {
		asciiStc, replaces := ConvertSentence2ASCII(ctx, c.sentence)
		t.Logf("case: %s, len=%d", c.sentence, len(c.sentence))
		t.Logf("convert: %s, len=%d, %v", asciiStc, len(asciiStc), replaces)

		if asciiStc != c.asciiStc {
			t.Fatalf("case-%d", i)
		}
	}
}


func TestDoCutSentenceLessThenMin(t *testing.T) {
	cases := []struct{
		sentence string
		hl []int32

		cuttedSentence string
		cuttedHL []int32
		keepHead, keepTail bool
	}{
		{//there
			"Seems there is no other way.",  []int32{6, 11},
			"Seems there is no other way.", []int32{6, 11}, true, true,
		},
		{//other way
			"Seems there is no other way.",  []int32{18, 27},
			"Seems there is no other way.", []int32{18, 27}, true, true,
		},
		{//my
			"What you said was exactly what was weighing on my mind before.", []int32{47, 49},
			"What you said was exactly what was weighing on my mind before.", []int32{47, 49}, true, true,
		},
		{//weighing
			"What you said was exactly what was weighing on my mind before.", []int32{18, 25},
			"What you said was exactly what was weighing on my mind before.", []int32{18, 25}, true, true,
		},
		{//there
			"Seems there is no 其他 way.",  []int32{6, 11},
			"Seems there is no 其他 way.", []int32{6, 11}, true, true,
		},
		{//there
			"看起来 there is no 其他 way.",  []int32{4, 9},
			"看起来 there is no 其他 way.", []int32{4, 9}, true, true,
		},
		{//其他
			"看起来 there is no 其他 way.",  []int32{16, 18},
			"看起来 there is no 其他 way.", []int32{16, 18}, true, true,
		},
		{//there is
			"看起来 there is no 其他 way.",  []int32{4, 12},
			"看起来 there is no 其他 way.", []int32{4, 12}, true, true,
		},
	}

	ctx := context.Background()

	for i, c := range cases {
		cuttedStc, cuttedHL, keepHead, keepTail := doCutSentence(ctx, c.sentence, c.hl, 150)
		t.Logf("case-%d: %s, [%d, %d)", i, c.sentence, c.cuttedHL[0], c.cuttedHL[1])
		t.Logf("cutted: %s, [%d, %d), (%t, %t)", cuttedStc, cuttedHL[0], cuttedHL[1], keepHead, keepTail)

		if cuttedStc != c.cuttedSentence {
			t.Fatalf("case-%d", i)
		}
		if cuttedHL[0] != c.cuttedHL[0] || cuttedHL[1] != c.cuttedHL[1] {
			t.Fatalf("case-%d", i)
		}
		if keepHead != c.keepHead || keepTail != c.keepTail {
			t.Fatalf("case-%d", i)
		}
		if runeStc, runeCuttedStc := []rune(c.sentence), []rune(cuttedStc); string(runeStc[c.hl[0]:c.hl[1]]) != string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]) {
			t.Fatalf("case-%d", i)
		} else {
			t.Logf("<%s>, <%s>", string(runeStc[c.hl[0]:c.hl[1]]), string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]))
		}
	}

}


func TestDoCutSentenceMoreThenMin(t *testing.T) {
	cases := []struct{
		sentence string
		hl []int32

		cuttedSentence string
		cuttedHL []int32
		keepHead, keepTail bool
	}{
		{// facilitates
			"Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{84, 95},
			"events, does your team have any system that facilitates in-team communication and help?", []int32{44, 55}, false, true,
		},
		{// company-wide
			"Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{17, 29},
			"Other than those company-wide tools and events, does your team have any system that", []int32{17, 29}, true, false,
		},
		{// team
			"Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{58, 62},
			"company-wide tools and events, does your team have any system that facilitates in-team", []int32{41, 45}, false, false,
		},
		{// person
			"We have a mentor-mentee program within our team, where a member who's been with the team for some time would act as a mentor, the go-to person, for a newcomer for the six-month probation period.",  []int32{136, 142},
			"would act as a mentor, the go-to person, for a newcomer for the six-month probation period.", []int32{33, 39}, false, true,
		},
		{// team
			"We have a mentor-mentee program within our team, where a member who's been with the team for some time would act as a mentor, the go-to person, for a newcomer for the six-month probation period.",  []int32{43, 47},
			"We have a mentor-mentee program within our team, where a member who's been with the", []int32{43, 47}, true, false,
		},
		{// within our team
			"We have a mentor-mentee program within our team, where a member who's been with the team for some time would act as a mentor, the go-to person, for a newcomer for the six-month probation period.",  []int32{32, 47},
			"We have a mentor-mentee program within our team, where a member who's been with the", []int32{32, 47}, true, false,
		},
		{// any system
			"Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{68, 78},
			"and events, does your team have any system that facilitates in-team communication", []int32{32, 42}, false, true,
		},
	}
	ctx := context.Background()
	for i, c := range cases {
		cuttedStc, cuttedHL, keepHead, keepTail := doCutSentence(ctx, c.sentence, c.hl, 80)
		t.Logf("case-%d: %s, [%d, %d)", i, c.sentence, c.cuttedHL[0], c.cuttedHL[1])
		t.Logf("cutted: %s, [%d, %d), (%t, %t)", cuttedStc, cuttedHL[0], cuttedHL[1], keepHead, keepTail)

		if cuttedStc != c.cuttedSentence {
			t.Fatalf("case-%d", i)
		}
		if cuttedHL[0] != c.cuttedHL[0] || cuttedHL[1] != c.cuttedHL[1] {
			t.Fatalf("case-%d", i)
		}
		if keepHead != c.keepHead || keepTail != c.keepTail {
			t.Fatalf("case-%d", i)
		}
		if runeStc, runeCuttedStc := []rune(c.sentence), []rune(cuttedStc); string(runeStc[c.hl[0]:c.hl[1]]) != string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]) {
			t.Fatalf("case-%d", i)
		} else {
			t.Logf("<%s>, <%s>", string(runeStc[c.hl[0]:c.hl[1]]), string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]))
		}
	}
}

func TestDoCutSentenceMoreThenMin2(t *testing.T) {
	cases := []struct{
		sentence string
		hl []int32

		cuttedSentence string
		cuttedHL []int32
		keepHead, keepTail bool
	}{
		{// elite
			"It was mentioned repeatedly in the feedback from the users in third- and fourth-tier cities that all they could find on our platform is videos featuring big city skylines and elite lifestyles.",  []int32{175, 180},
			"feedback from the users in third- and fourth-tier cities that all they could find on our platform is videos featuring big city skylines and elite lifestyles.", []int32{140, 145}, false, true,
		},
	}
	ctx := context.Background()
	for i, c := range cases {
		cuttedStc, cuttedHL, keepHead, keepTail := doCutSentence(ctx, c.sentence, c.hl, 150)
		t.Logf("case-%d: %s, [%d, %d)", i, c.sentence, c.cuttedHL[0], c.cuttedHL[1])
		t.Logf("cutted: %s, [%d, %d), (%t, %t)", cuttedStc, cuttedHL[0], cuttedHL[1], keepHead, keepTail)

		if cuttedStc != c.cuttedSentence {
			t.Fatalf("case-%d", i)
		}
		if cuttedHL[0] != c.cuttedHL[0] || cuttedHL[1] != c.cuttedHL[1] {
			t.Fatalf("case-%d", i)
		}
		if keepHead != c.keepHead || keepTail != c.keepTail {
			t.Fatalf("case-%d", i)
		}
		if runeStc, runeCuttedStc := []rune(c.sentence), []rune(cuttedStc); string(runeStc[c.hl[0]:c.hl[1]]) != string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]) {
			t.Fatalf("case-%d", i)
		} else {
			t.Logf("<%s>, <%s>", string(runeStc[c.hl[0]:c.hl[1]]), string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]))
		}

	}


}



func TestDoCutSentenceMoreSpace(t *testing.T) {
	cases := []struct{
		sentence string
		hl []int32

		cuttedSentence string
		cuttedHL []int32
		keepHead, keepTail bool
	}{
		{// company-wide，句首有空
			" Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{17+1, 29+1},
			" Other than those company-wide tools and events, does your team have any system that", []int32{18, 30}, true, false,
		},
		{// company-wide，句首多个空格
			"  Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{17+2, 29+2},
			"  Other than those company-wide tools and events, does your team have any system", []int32{19, 31}, true, false,
		},
		{// company-wide，句首多个空格
			"   Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{17+3, 29+3},
			"   Other than those company-wide tools and events, does your team have any system", []int32{20, 32}, true, false,
		},
		{// facilitates
			" Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{84+1, 95+1},
			"events, does your team have any system that facilitates in-team communication and help?", []int32{44, 55}, false, true,
		},
		{// facilitates
			"   Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{84+3, 95+3},
			"events, does your team have any system that facilitates in-team communication and help?", []int32{44, 55}, false, true,
		},
		{// facilitates， any__system
			"Other than those company-wide tools and events, does your team have any  system that facilitates in-team communication and help?",  []int32{84+1, 95+1},
			"does your team have any  system that facilitates in-team communication and help?", []int32{37, 48}, false, true,
		},
		{// facilitates， any___system
			"Other than those company-wide tools and events, does your team have any   system that facilitates in-team communication and help?",  []int32{84+2, 95+2},
			"does your team have any   system that facilitates in-team communication and help?", []int32{38, 49}, false, true,
		},
		{// facilitates， any____system
			"Other than those company-wide tools and events, does your team have any   system that facilitates in-team communication and help?",  []int32{84+3, 95+3},
			"does your team have any   system that facilitates in-team communication and help?", []int32{39, 50}, false, true,
		},
		{// team， that__facilitates
			"Other than those company-wide tools and events, does your team have any system that  facilitates in-team communication and help?",  []int32{58, 62},
			"those company-wide tools and events, does your team have any system that  facilitates", []int32{47, 51}, false, false,
		},
		{// team， that___facilitates
			"Other than those company-wide tools and events, does your team have any system that   facilitates in-team communication and help?",  []int32{58, 62},
			"those company-wide tools and events, does your team have any system that   facilitates", []int32{47, 51}, false, false,
		},
	}

	ctx := context.Background()
	for i, c := range cases {
		cuttedStc, cuttedHL, keepHead, keepTail := doCutSentence(ctx, c.sentence, c.hl, 80)
		t.Logf("case-%d: %s, [%d, %d)", i, c.sentence, c.cuttedHL[0], c.cuttedHL[1])
		t.Logf("cutted: %s, [%d, %d), (%t, %t)", cuttedStc, cuttedHL[0], cuttedHL[1], keepHead, keepTail)

		if cuttedStc != c.cuttedSentence {
			t.Fatalf("case-%d", i)
		}
		if cuttedHL[0] != c.cuttedHL[0] || cuttedHL[1] != c.cuttedHL[1] {
			t.Fatalf("case-%d", i)
		}
		if keepHead != c.keepHead || keepTail != c.keepTail {
			t.Fatalf("case-%d", i)
		}
		if runeStc, runeCuttedStc := []rune(c.sentence), []rune(cuttedStc); string(runeStc[c.hl[0]:c.hl[1]]) != string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]) {
			t.Fatalf("case-%d", i)
		} else {
			t.Logf("<%s>, <%s>", string(runeStc[c.hl[0]:c.hl[1]]), string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]))
		}
	}
}



func TestDoCutSentenceWithNonASCII(t *testing.T) {
	cases := []struct{
		sentence string
		hl []int32

		cuttedSentence string
		cuttedHL []int32
		keepHead, keepTail bool
	}{
		{// company-wide
			"“Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{18, 30},
			"“Other than those company-wide tools and events, does your team have any system that", []int32{18, 30}, true, false,
		},
		{// company-wide
			"句首Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{19, 31},
			"句首Other than those company-wide tools and events, does your team have any system", []int32{19, 31}, true, false,
		},
		{// company-wide
			"句首 Other than those company-wide tools and events, does your team have any system that facilitates in-team communication and help?",  []int32{20, 32},
			"句首 Other than those company-wide tools and events, does your team have any system", []int32{20, 32}, true, false,
		},
		{// facilitates
			"Other than those company-wide tools and events, 这一段 does your team have any system that facilitates in-team communication and help?",  []int32{84+4, 95+4},
			"这一段 does your team have any system that facilitates in-team communication and help?", []int32{40, 51}, false, true,
		},
		{// facilitates
			"句首Other than those company-wide 句中 tools and events, 这一段 does your team have any system that facilitates in-team communication and help?",  []int32{93, 104},
			"这一段 does your team have any system that facilitates in-team communication and help?", []int32{40, 51}, false, true,
		},

		{// company-wide
			"- Other than those company-wide tools and events.\n- does your team have any system that facilitates in-team communication and help?",  []int32{19, 31},
			"- Other than those company-wide tools and events.\n- does your team have any system", []int32{19, 31}, true, false,
		},
		{// company-wide
			"- “Other than those company-wide tools and events.”\n- “does your team have any system that facilitates in-team communication and help?”",  []int32{20, 32},
			"- “Other than those company-wide tools and events.”\n- “does your team have any system", []int32{20, 32}, true, false,
		},
		{// facilitates
			"- Other than those company-wide tools and events, \n- does your team have any system that facilitates in-team communication and help?",  []int32{89, 100},
			"\n- does your team have any system that facilitates in-team communication and help?", []int32{39, 50}, false, true,
		},
		{// facilitates
			"- “Other than those company-wide tools and events,” \n- “does your team have any system that facilitates in-team communication and help?”",  []int32{92, 103},
			"“does your team have any system that facilitates in-team communication and help?”", []int32{37, 48}, false, true,
		},
	}

	ctx := context.Background()
	for i, c := range cases {
		cuttedStc, cuttedHL, keepHead, keepTail := doCutSentence(ctx, c.sentence, c.hl, 80)
		t.Logf("case-%d: %s, [%d, %d)", i, c.sentence, c.cuttedHL[0], c.cuttedHL[1])
		t.Logf("cutted: %s, [%d, %d), (%t, %t)", cuttedStc, cuttedHL[0], cuttedHL[1], keepHead, keepTail)

		if cuttedStc != c.cuttedSentence {
			t.Fatalf("case-%d", i)
		}
		if cuttedHL[0] != c.cuttedHL[0] || cuttedHL[1] != c.cuttedHL[1] {
			t.Fatalf("case-%d", i)
		}
		if keepHead != c.keepHead || keepTail != c.keepTail {
			t.Fatalf("case-%d", i)
		}
		if runeStc, runeCuttedStc := []rune(c.sentence), []rune(cuttedStc); string(runeStc[c.hl[0]:c.hl[1]]) != string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]) {
			t.Fatalf("case-%d", i)
		} else {
			t.Logf("<%s>, <%s>", string(runeStc[c.hl[0]:c.hl[1]]), string(runeCuttedStc[cuttedHL[0]:cuttedHL[1]]))
		}
	}

}

