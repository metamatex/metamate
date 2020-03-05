package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Verb struct {
	Processed  bool
	Infinitive string
	Present3rd string
	Past3rd    string
	Additional string
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	f, err := os.Open("list.txt")
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	verbs := []Verb{}
	for {
		var s string
		s, err = r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
				goto done
			}

			return
		}

		s = strings.Replace(s, "\n", "", -1)

		parts := strings.Split(s, " ")

		verb := Verb{}
		verb.Infinitive = parts[1]
		if len(parts) > 2 {
			verb.Additional = concat(parts[2:])
		}

		verbs = append(verbs, verb)
	}
done:

	//spew.Dump(verbs)

	//i := 1
	verbs = process(verbs)

	//spew.Dump(verbs)

	verbs = removeEmpty(verbs)

	formatInfintive(verbs)
	//verbs = getRuleBreaker(verbs)

	//print(verbs)
	//spew.Dump(len(verbs))

	return
}

func process(verbs []Verb) ([]Verb) {
	for i, _ := range verbs {
		if contains(verbs[i], IRREGULAR) {
			verbs[i] = processIrregular(verbs[i])

			continue
		}

		if contains(verbs[i], RULE_BREAKER) {
			verbs[i] = processRuleBreaker(verbs[i])

			continue
		}

		if contains(verbs[i], DROP_THE_FINAL_E) {
			verbs[i] = processDropTheFinalE(verbs[i])

			continue
		}

		if contains(verbs[i], DOUBLE_THE_FINAL_CONSONANT) {
			verbs[i] = processDoubleTheFinalConsonant(verbs[i])

			continue
		}

		if contains(verbs[i], CHANGE_THE_Y_TO_I) {
			verbs[i] = processChangeToYToI(verbs[i])

			continue
		}

		if contains(verbs[i], ADD_ES) {
			verbs[i] = processAddEs(verbs[i])

			continue
		}

		if contains(verbs[i], REGULAR) {
			verbs[i] = processRegular(verbs[i])

			continue
		}
	}

	return verbs
}

func processRegular(v Verb) (Verb) {
	v.Present3rd = v.Infinitive + "s"
	v.Past3rd = v.Infinitive + "ed"

	v.Processed = true

	return v
}

func processDoubleTheFinalConsonant(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	if strings.HasSuffix(v.Infinitive, "l") {
		v.Present3rd = v.Infinitive + "s"
		v.Past3rd = v.Infinitive + "ed"
	} else {
		v.Present3rd = v.Infinitive + "s"
		v.Past3rd = v.Infinitive + v.Infinitive[len(v.Infinitive)-1:] + "ed"
	}

	v.Processed = true

	return v
}

func processAddEs(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	v.Present3rd = v.Infinitive + "es"
	v.Past3rd = v.Infinitive + "ed"

	v.Processed = true

	return v
}

func processChangeToYToI(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	v.Present3rd = v.Infinitive[:len(v.Infinitive)-1] + "ies"
	v.Past3rd = v.Infinitive[:len(v.Infinitive)-1] + "ied"

	v.Processed = true

	return v
}

func processDropTheFinalE(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	v.Present3rd = v.Infinitive + "s"
	v.Past3rd = v.Infinitive + "d"

	v.Processed = true

	return v
}

func processIrregular(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	v.Processed = true

	return v
}

func processRuleBreaker(v Verb) (Verb) {
	if v.Processed == true {
		panic(fmt.Sprintf("verb %v already processed", v.Infinitive))
	}

	v.Processed = true

	return v
}

const (
	REGULAR                    = ""
	IRREGULAR                  = "+"
	RULE_BREAKER               = "RB"
	DROP_THE_FINAL_E           = "1"
	DOUBLE_THE_FINAL_CONSONANT = "2"
	CHANGE_THE_Y_TO_I          = "3"
	ADD_ES                     = "4"
)

func contains(verb Verb, ss ...string) bool {
	for _, s := range ss {
		if !strings.Contains(verb.Additional, s) {
			return false
		}
	}

	return true
}

func print(verbs []Verb) {
	for _, v := range verbs {
		println(v.Infinitive)
	}
}

func concat(ss []string) (s string) {
	for _, s0 := range ss {
		s = s + " " + s0
	}

	return
}

func formatPresent(verbs []Verb) {
	fmt.Printf("package present\n\ntype Present string\n\nconst (\n")
	for _, v := range verbs {
		fmt.Printf("	%s Present = \"%s\"\n", strings.Title(v.Present3rd), v.Present3rd)
	}
	fmt.Printf(")")
}

func formatPast(verbs []Verb) {
	fmt.Printf("package past\n\ntype Past string\n\nconst (\n")
	for _, v := range verbs {
		fmt.Printf("	%s Past = \"%s\"\n", strings.Title(v.Past3rd), v.Past3rd)
	}
	fmt.Printf(")")
}

func formatInfintive(verbs []Verb) {
	fmt.Printf("package past\n\ntype Past string\n\nconst (\n")
	for _, v := range verbs {
		fmt.Printf("	%s = \"%s\"\n", strings.Title(v.Infinitive), v.Infinitive)
	}
	fmt.Printf(")")
}

func removeEmpty(verbs []Verb) (verbs0 []Verb) {
	for _, v := range verbs {
		if v.Present3rd == "" {
			continue
		}

		verbs0 = append(verbs0, v)
	}

	return
}
