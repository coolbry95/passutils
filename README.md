# passutils
Utilities to aid in password cracking

###maskgen
Maskgen generates hashcat masks. Supply an input file and it outputs to stdout.

###potcut
Potcut takes the pot file and original hashlist and outputs the plains to stdout. This works on salted hashlists as well as non salted hashlists.

###ruleprocessor
Ruleprocessor takes a wordlist and a rule file and applies the rules to the wordlist. It is not the fastest ever but will handle UTF-8 rules properly. There is no need to split a two byte character into two rules. eg. $\x03 $\x04 could just be a single $\x07. NOTE that the rules were just made up and I do not know what they actually append.

Ruleprocessor also has a package rules which contains all of the needed functions to parse a rule file and apply all of the rules to a word.
