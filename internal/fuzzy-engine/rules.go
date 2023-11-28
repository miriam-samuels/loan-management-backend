package fis

//	note: so you don't get confused
//	these are the rules generated to judge user creditworthiness
//	{CreditScore, Collateral, Income, CriminalRecord, EmploymentTerm, Creditworthines}
// {"bad", "bad", "lower", "bad", "short", "bad"},
//  if credicscore is bad and collatral is bad and incomelevel is lower and criminalrecord is bad and employmentterm is short
// then creditworthiness is bad
// was going to use numeral before .. but did not want to confuse the reader more... pele dear
//  could not think of a beeter pathway .. if you do fing one ... tell me and just create a pr with the solution

var Rules []FISRules = []FISRules{
	{"bad", "bad", "lower", "bad", "short", "bad"},
	{"average", "bad", "lower", "bad", "short", "bad"},
	{"good", "bad", "lower", "bad", "short", "bad"},
	{"bad", "average", "lower", "bad", "short", "bad"},
	{"average", "average", "lower", "bad", "short", "bad"},
	{"good", "average", "lower", "bad", "short", "bad"},
	{"bad", "good", "lower", "bad", "short", "bad"},
	{"average", "good", "lower", "bad", "short", "bad"},
	{"good", "good", "lower", "bad", "short", "bad"},
	{"bad", "bad", "middle", "bad", "short", "bad"},
	{"average", "bad", "middle", "bad", "short", "bad"},
	{"good", "bad", "middle", "bad", "short", "bad"},
	{"bad", "average", "middle", "bad", "short", "bad"},
	{"average", "average", "middle", "bad", "short", "bad"},
	{"good", "average", "middle", "bad", "short", "bad"},
	{"bad", "good", "middle", "bad", "short", "bad"},
	{"average", "good", "middle", "bad", "short", "bad"},
	{"good", "good", "middle", "bad", "short", "bad"},
	{"bad", "bad", "upper", "bad", "short", "bad"},
	{"average", "bad", "upper", "bad", "short", "bad"},
	{"good", "bad", "upper", "bad", "short", "bad"},
	{"bad", "average", "upper", "bad", "short", "bad"},
	{"average", "average", "upper", "bad", "short", "bad"},
	{"good", "average", "upper", "bad", "short", "bad"},
	{"bad", "good", "upper", "bad", "short", "bad"},
	{"average", "good", "upper", "bad", "short", "bad"},
	{"good", "good", "upper", "bad", "short", "bad"},
	{"bad", "bad", "lower", "fair", "short", "bad"},
	{"average", "bad", "lower", "fair", "short", "bad"},
	{"good", "bad", "lower", "fair", "short", "bad"},
	{"bad", "average", "lower", "fair", "short", "bad"},
	{"average", "average", "lower", "fair", "short", "bad"},
	{"good", "average", "lower", "fair", "short", "bad"},
	{"bad", "good", "lower", "fair", "short", "bad"},
	{"average", "good", "lower", "fair", "short", "bad"},
	{"good", "good", "lower", "fair", "short", "bad"},
	{"bad", "bad", "middle", "fair", "short", "bad"},
	{"average", "bad", "middle", "fair", "short", "average"},
	{"good", "bad", "middle", "fair", "short", "average"},
	{"bad", "average", "middle", "fair", "short", "bad"},
	{"average", "average", "middle", "fair", "short", "average"},
	{"good", "average", "middle", "fair", "short", "bad"}, // stop point
	{"bad", "good", "middle", "fair", "short", "bad"},
	{"average", "good", "middle", "fair", "short", "bad"},
	{"good", "good", "middle", "fair", "short", "bad"},
	{"bad", "bad", "upper", "fair", "short", "bad"},
	{"average", "bad", "upper", "fair", "short", "bad"},
	{"good", "bad", "upper", "fair", "short", "bad"},
	{"bad", "average", "upper", "fair", "short", "bad"},
	{"average", "average", "upper", "fair", "short", "bad"},
	{"good", "average", "upper", "fair", "short", "bad"},
	{"bad", "good", "upper", "fair", "short", "bad"},
	{"average", "good", "upper", "fair", "short", "bad"},
	{"good", "good", "upper", "fair", "short", "bad"},
	{"bad", "bad", "lower", "good", "short", "bad"},
	{"average", "bad", "lower", "good", "short", "bad"},
	{"good", "bad", "lower", "good", "short", "bad"},
	{"bad", "average", "lower", "good", "short", "bad"},
	{"average", "average", "lower", "good", "short", "bad"},
	{"good", "average", "lower", "good", "short", "bad"},
	{"bad", "good", "lower", "good", "short", "bad"},
	{"average", "good", "lower", "good", "short", "bad"},
	{"good", "good", "lower", "good", "short", "bad"},
	{"bad", "bad", "middle", "good", "short", "bad"},
	{"average", "bad", "middle", "good", "short", "bad"},
	{"good", "bad", "middle", "good", "short", "bad"},
	{"bad", "average", "middle", "good", "short", "bad"},
	{"average", "average", "middle", "good", "short", "bad"},
	{"good", "average", "middle", "good", "short", "bad"},
	{"bad", "good", "middle", "good", "short", "bad"},
	{"average", "good", "middle", "good", "short", "bad"},
	{"good", "good", "middle", "good", "short", "bad"},
	{"bad", "bad", "upper", "good", "short", "bad"},
	{"average", "bad", "upper", "good", "short", "bad"},
	{"good", "bad", "upper", "good", "short", "bad"},
	{"bad", "average", "upper", "good", "short", "bad"},
	{"average", "average", "upper", "good", "short", "bad"},
	{"good", "average", "upper", "good", "short", "bad"},
	{"bad", "good", "upper", "good", "short", "bad"},
	{"average", "good", "upper", "good", "short", "bad"},
	{"good", "good", "upper", "good", "short", "bad"},
	{"bad", "bad", "lower", "bad", "medium", "bad"},
	{"average", "bad", "lower", "bad", "medium", "bad"},
	{"good", "bad", "lower", "bad", "medium", "bad"},
	{"bad", "average", "lower", "bad", "medium", "bad"},
	{"average", "average", "lower", "bad", "medium", "bad"},
	{"good", "average", "lower", "bad", "medium", "bad"},
	{"bad", "good", "lower", "bad", "medium", "bad"},
	{"average", "good", "lower", "bad", "medium", "bad"},
	{"good", "good", "lower", "bad", "medium", "bad"},
	{"bad", "bad", "middle", "bad", "medium", "bad"},
	{"average", "bad", "middle", "bad", "medium", "bad"},
	{"good", "bad", "middle", "bad", "medium", "bad"},
	{"bad", "average", "middle", "bad", "medium", "bad"},
	{"average", "average", "middle", "bad", "medium", "bad"},
	{"good", "average", "middle", "bad", "medium", "bad"},
	{"bad", "good", "middle", "bad", "medium", "bad"},
	{"average", "good", "middle", "bad", "medium", "bad"},
	{"good", "good", "middle", "bad", "medium", "bad"},
	{"bad", "bad", "upper", "bad", "medium", "bad"},
	{"average", "bad", "upper", "bad", "medium", "bad"},
	{"good", "bad", "upper", "bad", "medium", "bad"},
	{"bad", "average", "upper", "bad", "medium", "bad"},
	{"average", "average", "upper", "bad", "medium", "bad"},
	{"good", "average", "upper", "bad", "medium", "bad"},
	{"bad", "good", "upper", "bad", "medium", "bad"},
	{"average", "good", "upper", "bad", "medium", "bad"},
	{"good", "good", "upper", "bad", "medium", "bad"},
	{"bad", "bad", "lower", "fair", "medium", "bad"},
	{"average", "bad", "lower", "fair", "medium", "bad"},
	{"good", "bad", "lower", "fair", "medium", "bad"},
	{"bad", "average", "lower", "fair", "medium", "bad"},
	{"average", "average", "lower", "fair", "medium", "bad"},
	{"good", "average", "lower", "fair", "medium", "bad"},
	{"bad", "good", "lower", "fair", "medium", "bad"},
	{"average", "good", "lower", "fair", "medium", "bad"},
	{"good", "good", "lower", "fair", "medium", "bad"},
	{"bad", "bad", "middle", "fair", "medium", "bad"},
	{"average", "bad", "middle", "fair", "medium", "bad"},
	{"good", "bad", "middle", "fair", "medium", "bad"},
	{"bad", "average", "middle", "fair", "medium", "bad"},
	{"average", "average", "middle", "fair", "medium", "bad"},
	{"good", "average", "middle", "fair", "medium", "bad"},
	{"bad", "good", "middle", "fair", "medium", "bad"},
	{"average", "good", "middle", "fair", "medium", "bad"},
	{"good", "good", "middle", "fair", "medium", "bad"},
	{"bad", "bad", "upper", "fair", "medium", "bad"},
	{"average", "bad", "upper", "fair", "medium", "bad"},
	{"good", "bad", "upper", "fair", "medium", "bad"},
	{"bad", "average", "upper", "fair", "medium", "bad"},
	{"average", "average", "upper", "fair", "medium", "bad"},
	{"good", "average", "upper", "fair", "medium", "bad"},
	{"bad", "good", "upper", "fair", "medium", "bad"},
	{"average", "good", "upper", "fair", "medium", "bad"},
	{"good", "good", "upper", "fair", "medium", "bad"},
	{"bad", "bad", "lower", "good", "medium", "bad"},
	{"average", "bad", "lower", "good", "medium", "bad"},
	{"good", "bad", "lower", "good", "medium", "bad"},
	{"bad", "average", "lower", "good", "medium", "bad"},
	{"average", "average", "lower", "good", "medium", "bad"},
	{"good", "average", "lower", "good", "medium", "bad"},
	{"bad", "good", "lower", "good", "medium", "bad"},
	{"average", "good", "lower", "good", "medium", "bad"},
	{"good", "good", "lower", "good", "medium", "bad"},
	{"bad", "bad", "middle", "good", "medium", "bad"},
	{"average", "bad", "middle", "good", "medium", "bad"},
	{"good", "bad", "middle", "good", "medium", "bad"},
	{"bad", "average", "middle", "good", "medium", "bad"},
	{"average", "average", "middle", "good", "medium", "bad"},
	{"good", "average", "middle", "good", "medium", "bad"},
	{"bad", "good", "middle", "good", "medium", "bad"},
	{"average", "good", "middle", "good", "medium", "bad"},
	{"good", "good", "middle", "good", "medium", "bad"},
	{"bad", "bad", "upper", "good", "medium", "bad"},
	{"average", "bad", "upper", "good", "medium", "bad"},
	{"good", "bad", "upper", "good", "medium", "bad"},
	{"bad", "average", "upper", "good", "medium", "bad"},
	{"average", "average", "upper", "good", "medium", "bad"},
	{"good", "average", "upper", "good", "medium", "bad"},
	{"bad", "good", "upper", "good", "medium", "bad"},
	{"average", "good", "upper", "good", "medium", "bad"},
	{"good", "good", "upper", "good", "medium", "bad"},
	{"bad", "bad", "lower", "bad", "long", "bad"},
	{"average", "bad", "lower", "bad", "long", "bad"},
	{"good", "bad", "lower", "bad", "long", "bad"},
	{"bad", "average", "lower", "bad", "long", "bad"},
	{"average", "average", "lower", "bad", "long", "bad"},
	{"good", "average", "lower", "bad", "long", "bad"},
	{"bad", "good", "lower", "bad", "long", "bad"},
	{"average", "good", "lower", "bad", "long", "bad"},
	{"good", "good", "lower", "bad", "long", "bad"},
	{"bad", "bad", "middle", "bad", "long", "bad"},
	{"average", "bad", "middle", "bad", "long", "bad"},
	{"good", "bad", "middle", "bad", "long", "bad"},
	{"bad", "average", "middle", "bad", "long", "bad"},
	{"average", "average", "middle", "bad", "long", "bad"},
	{"good", "average", "middle", "bad", "long", "bad"},
	{"bad", "good", "middle", "bad", "long", "bad"},
	{"average", "good", "middle", "bad", "long", "bad"},
	{"good", "good", "middle", "bad", "long", "bad"},
	{"bad", "bad", "upper", "bad", "long", "bad"},
	{"average", "bad", "upper", "bad", "long", "bad"},
	{"good", "bad", "upper", "bad", "long", "bad"},
	{"bad", "average", "upper", "bad", "long", "bad"},
	{"average", "average", "upper", "bad", "long", "bad"},
	{"good", "average", "upper", "bad", "long", "bad"},
	{"bad", "good", "upper", "bad", "long", "bad"},
	{"average", "good", "upper", "bad", "long", "bad"},
	{"good", "good", "upper", "bad", "long", "bad"},
	{"bad", "bad", "lower", "fair", "long", "bad"},
	{"average", "bad", "lower", "fair", "long", "bad"},
	{"good", "bad", "lower", "fair", "long", "bad"},
	{"bad", "average", "lower", "fair", "long", "bad"},
	{"average", "average", "lower", "fair", "long", "bad"},
	{"good", "average", "lower", "fair", "long", "bad"},
	{"bad", "good", "lower", "fair", "long", "bad"},
	{"average", "good", "lower", "fair", "long", "bad"},
	{"good", "good", "lower", "fair", "long", "bad"},
	{"bad", "bad", "middle", "fair", "long", "bad"},
	{"average", "bad", "middle", "fair", "long", "bad"},
	{"good", "bad", "middle", "fair", "long", "bad"},
	{"bad", "average", "middle", "fair", "long", "bad"},
	{"average", "average", "middle", "fair", "long", "bad"},
	{"good", "average", "middle", "fair", "long", "bad"},
	{"bad", "good", "middle", "fair", "long", "bad"},
	{"average", "good", "middle", "fair", "long", "bad"},
	{"good", "good", "middle", "fair", "long", "bad"},
	{"bad", "bad", "upper", "fair", "long", "bad"},
	{"average", "bad", "upper", "fair", "long", "bad"},
	{"good", "bad", "upper", "fair", "long", "good"}, // stop point
	{"bad", "average", "upper", "fair", "long", "bad"},
	{"average", "average", "upper", "fair", "long", "average"},
	{"good", "average", "upper", "fair", "long", "good"},
	{"bad", "good", "upper", "fair", "long", "average"},
	{"average", "good", "upper", "fair", "long", "good"},
	{"good", "good", "upper", "fair", "long", "good"},
	{"bad", "bad", "lower", "good", "long", "bad"},
	{"average", "bad", "lower", "good", "long", "bad"},
	{"good", "bad", "lower", "good", "long", "bad"},
	{"bad", "average", "lower", "good", "long", "bad"},
	{"average", "average", "lower", "good", "long", "bad"},
	{"good", "average", "lower", "good", "long", "bad"},
	{"bad", "good", "lower", "good", "long", "bad"},
	{"average", "good", "lower", "good", "long", "bad"},
	{"good", "good", "lower", "good", "long", "bad"},
	{"bad", "bad", "middle", "good", "long", "bad"},
	{"average", "bad", "middle", "good", "long", "average"},
	{"good", "bad", "middle", "good", "long", "average"},
	{"bad", "average", "middle", "good", "long", "bad"},
	{"average", "average", "middle", "good", "long", "average"},
	{"good", "average", "middle", "good", "long", "good"},
	{"bad", "good", "middle", "good", "long", "bad"},
	{"average", "good", "middle", "good", "long", "average"},
	{"good", "good", "middle", "good", "long", "good"},
	{"bad", "bad", "upper", "good", "long", "average"},
	{"average", "bad", "upper", "good", "long", "good"},
	{"good", "bad", "upper", "good", "long", "good"},
	{"bad", "average", "upper", "good", "long", "average"},
	{"average", "average", "upper", "good", "long", "good"},
	{"good", "average", "upper", "good", "long", "good"},
	{"bad", "good", "upper", "good", "long", "average"},
	{"average", "good", "upper", "good", "long", "good"},
	{"good", "good", "upper", "good", "long", "good"},
}
