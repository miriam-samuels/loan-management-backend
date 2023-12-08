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
	{"good", "average", "middle", "fair", "short", "average"},
	{"bad", "good", "middle", "fair", "short", "bad"},
	{"average", "good", "middle", "fair", "short", "average"},
	{"good", "good", "middle", "fair", "short", "average"},
	{"bad", "bad", "upper", "fair", "short", "bad"},
	{"average", "bad", "upper", "fair", "short", "average"},
	{"good", "bad", "upper", "fair", "short", "average"},
	{"bad", "average", "upper", "fair", "short", "bad"},
	{"average", "average", "upper", "fair", "short", "average"},
	{"good", "average", "upper", "fair", "short", "average"},
	{"bad", "good", "upper", "fair", "short", "average"},
	{"average", "good", "upper", "fair", "short", "good"},
	{"good", "good", "upper", "fair", "short", "good"},
	{"bad", "bad", "lower", "good", "short", "bad"},
	{"average", "bad", "lower", "good", "short", "bad"},
	{"good", "bad", "lower", "good", "short", "bad"},
	{"bad", "average", "lower", "good", "short", "bad"},
	{"average", "average", "lower", "good", "short", "average"},
	{"good", "average", "lower", "good", "short", "average"},
	{"bad", "good", "lower", "good", "short", "bad"},
	{"average", "good", "lower", "good", "short", "average"},
	{"good", "good", "lower", "good", "short", "good"},
	{"bad", "bad", "middle", "good", "short", "bad"},
	{"average", "bad", "middle", "good", "short", "bad"},
	{"good", "bad", "middle", "good", "short", "average"},
	{"bad", "average", "middle", "good", "short", "bad"},
	{"average", "average", "middle", "good", "short", "average"},
	{"good", "average", "middle", "good", "short", "good"},
	{"bad", "good", "middle", "good", "short", "bad"},
	{"average", "good", "middle", "good", "short", "good"},
	{"good", "good", "middle", "good", "short", "good"},
	{"bad", "bad", "upper", "good", "short", "bad"},
	{"average", "bad", "upper", "good", "short", "average"},
	{"good", "bad", "upper", "good", "short", "good"},
	{"bad", "average", "upper", "good", "short", "bad"},
	{"average", "average", "upper", "good", "short", "average"},
	{"good", "average", "upper", "good", "short", "good"},
	{"bad", "good", "upper", "good", "short", "bad"},
	{"average", "good", "upper", "good", "short", "good"},
	{"good", "good", "upper", "good", "short", "good"},
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
	{"good", "average", "middle", "bad", "medium", "average"},
	{"bad", "good", "middle", "bad", "medium", "bad"},
	{"average", "good", "middle", "bad", "medium", "average"},
	{"good", "good", "middle", "bad", "medium", "average"},
	{"bad", "bad", "upper", "bad", "medium", "bad"},
	{"average", "bad", "upper", "bad", "medium", "bad"},
	{"good", "bad", "upper", "bad", "medium", "bad"},
	{"bad", "average", "upper", "bad", "medium", "bad"},
	{"average", "average", "upper", "bad", "medium", "average"},
	{"good", "average", "upper", "bad", "medium", "average"},
	{"bad", "good", "upper", "bad", "medium", "average"},
	{"average", "good", "upper", "bad", "medium", "good"},
	{"good", "good", "upper", "bad", "medium", "good"},
	{"bad", "bad", "lower", "fair", "medium", "bad"},
	{"average", "bad", "lower", "fair", "medium", "bad"},
	{"good", "bad", "lower", "fair", "medium", "bad"},
	{"bad", "average", "lower", "fair", "medium", "bad"},
	{"average", "average", "lower", "fair", "medium", "bad"},
	{"good", "average", "lower", "fair", "medium", "bad"},
	{"bad", "good", "lower", "fair", "medium", "bad"},
	{"average", "good", "lower", "fair", "medium", "average"},
	{"good", "good", "lower", "fair", "medium", "average"},
	{"bad", "bad", "middle", "fair", "medium", "bad"},
	{"average", "bad", "middle", "fair", "medium", "bad"},
	{"good", "bad", "middle", "fair", "medium", "bad"},
	{"bad", "average", "middle", "fair", "medium", "bad"},
	{"average", "average", "middle", "fair", "medium", "average"},
	{"good", "average", "middle", "fair", "medium", "average"},
	{"bad", "good", "middle", "fair", "medium", "bad"},
	{"average", "good", "middle", "fair", "medium", "good"},
	{"good", "good", "middle", "fair", "medium", "good"},
	{"bad", "bad", "upper", "fair", "medium", "bad"},
	{"average", "bad", "upper", "fair", "medium", "average"},
	{"good", "bad", "upper", "fair", "medium", "average"},
	{"bad", "average", "upper", "fair", "medium", "average"},
	{"average", "average", "upper", "fair", "medium", "average"},
	{"good", "average", "upper", "fair", "medium", "good"},
	{"bad", "good", "upper", "fair", "medium", "average"},
	{"average", "good", "upper", "fair", "medium", "good"},
	{"good", "good", "upper", "fair", "medium", "good"},
	{"bad", "bad", "lower", "good", "medium", "bad"},
	{"average", "bad", "lower", "good", "medium", "bad"},
	{"good", "bad", "lower", "good", "medium", "bad"},
	{"bad", "average", "lower", "good", "medium", "bad"},
	{"average", "average", "lower", "good", "medium", "average"},
	{"good", "average", "lower", "good", "medium", "average"},
	{"bad", "good", "lower", "good", "medium", "bad"},
	{"average", "good", "lower", "good", "medium", "good"},
	{"good", "good", "lower", "good", "medium", "good"},
	{"bad", "bad", "middle", "good", "medium", "bad"},
	{"average", "bad", "middle", "good", "medium", "average"},
	{"good", "bad", "middle", "good", "medium", "good"},
	{"bad", "average", "middle", "good", "medium", "average"},
	{"average", "average", "middle", "good", "medium", "average"},
	{"good", "average", "middle", "good", "medium", "good"},
	{"bad", "good", "middle", "good", "medium", "average"},
	{"average", "good", "middle", "good", "medium", "good"},
	{"good", "good", "middle", "good", "medium", "good"},
	{"bad", "bad", "upper", "good", "medium", "bad"},
	{"average", "bad", "upper", "good", "medium", "average"},
	{"good", "bad", "upper", "good", "medium", "good"},
	{"bad", "average", "upper", "good", "medium", "average"},
	{"average", "average", "upper", "good", "medium", "good"},
	{"good", "average", "upper", "good", "medium", "good"},
	{"bad", "good", "upper", "good", "medium", "average"},
	{"average", "good", "upper", "good", "medium", "good"},
	{"good", "good", "upper", "good", "medium", "good"},
	{"bad", "bad", "lower", "bad", "long", "bad"},
	{"average", "bad", "lower", "bad", "long", "bad"},
	{"good", "bad", "lower", "bad", "long", "bad"},
	{"bad", "average", "lower", "bad", "long", "bad"},
	{"average", "average", "lower", "bad", "long", "bad"},
	{"good", "average", "lower", "bad", "long", "bad"},
	{"bad", "good", "lower", "bad", "long", "average"},
	{"average", "good", "lower", "bad", "long", "bad"},
	{"good", "good", "lower", "bad", "long", "average"},
	{"bad", "bad", "middle", "bad", "long", "bad"},
	{"average", "bad", "middle", "bad", "long", "bad"},
	{"good", "bad", "middle", "bad", "long", "bad"},
	{"bad", "average", "middle", "bad", "long", "bad"},
	{"average", "average", "middle", "bad", "long", "average"},
	{"good", "average", "middle", "bad", "long", "average"},
	{"bad", "good", "middle", "bad", "long", "bad"},
	{"average", "good", "middle", "bad", "long", "average"},
	{"good", "good", "middle", "bad", "long", "good"},
	{"bad", "bad", "upper", "bad", "long", "bad"},
	{"average", "bad", "upper", "bad", "long", "average"},
	{"good", "bad", "upper", "bad", "long", "average"},
	{"bad", "average", "upper", "bad", "long", "bad"},
	{"average", "average", "upper", "bad", "long", "average"},
	{"good", "average", "upper", "bad", "long", "good"},
	{"bad", "good", "upper", "bad", "long", "bad"},
	{"average", "good", "upper", "bad", "long", "good"},
	{"good", "good", "upper", "bad", "long", "good"},
	{"bad", "bad", "lower", "fair", "long", "bad"},
	{"average", "bad", "lower", "fair", "long", "bad"},
	{"good", "bad", "lower", "fair", "long", "bad"},
	{"bad", "average", "lower", "fair", "long", "bad"},
	{"average", "average", "lower", "fair", "long", "bad"},
	{"good", "average", "lower", "fair", "long", "average"},
	{"bad", "good", "lower", "fair", "long", "bad"},
	{"average", "good", "lower", "fair", "long", "average"},
	{"good", "good", "lower", "fair", "long", "average"},
	{"bad", "bad", "middle", "fair", "long", "bad"},
	{"average", "bad", "middle", "fair", "long", "average"},
	{"good", "bad", "middle", "fair", "long", "average"},
	{"bad", "average", "middle", "fair", "long", "bad"},
	{"average", "average", "middle", "fair", "long", "average"},
	{"good", "average", "middle", "fair", "long", "good"},
	{"bad", "good", "middle", "fair", "long", "bad"},
	{"average", "good", "middle", "fair", "long", "average"},
	{"good", "good", "middle", "fair", "long", "good"},
	{"bad", "bad", "upper", "fair", "long", "bad"},
	{"average", "bad", "upper", "fair", "long", "bad"},
	{"good", "bad", "upper", "fair", "long", "good"},
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
