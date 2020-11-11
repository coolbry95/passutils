#!/usr/bin/env perl

# to generate the wordlist
# hashcat --stdout epixoip_top10k -r ruleall.rule | head > wordlist tail > wordlist

# to generate rules
# rules is from spread sheet
# single rules
# $d=1; for i in `cat rules`; do echo $i > rule.(($d +1)); done

# to get a list of functions
# grep 'func' apply.go

use strict;
use warnings;

my $hashcat = "/home/coolbry95/crack/hashcat-0.49/hashcat-cli64.bin";
my $wordlist = "/home/coolbry95/wrappertesting/testlist";

my $rulesingle = "/home/coolbry95/wrappertesting/rulesingle";
my $wordsingle = "/home/coolbry95/wrappertesting/wordsingle";

open(my $singlefh, '>', $rulesingle) or die "cannot open rulesingle: $!";
open(my $wordlistfh, '>', $wordsingle) or die "cannot open wordsingle: $!";
open(my $wordlistsfh, '<', $wordlist) or die "cannot open wordsingle: $!";

opendir(my $DIR, "/home/coolbry95/wrappertesting/rules/");
my @rules = readdir($DIR);
closedir($DIR);

my $rule;
my $plain;

select((select($singlefh), $| = 1)[0]);
select((select($wordlistfh), $| = 1)[0]);

my $temp = 0;
my $s = 0;
# loop through each rule file
foreach my $a (@rules) {
	chomp $a;
	next if $a eq ".";
	next if $a eq "..";

	#print $a . "\n";
	# open up the single rule file
	open(my $rulefh, '<', "/home/coolbry95/wrappertesting/rules/" . $a);
	# print the start of the variable
	#print "var ruletest = []struct {" . "\n";
	# start the loop of the rule file
	while ($rule = <$rulefh>) {
		# print the rule we are on now to the file
		print $singlefh $rule;

		# put this here to have the rule on top of it
		# for easier matching
		print '//' . $rule;
		print "var ruletest = []struct {" . "\n";
		print "string\nstring\n} {\n";
		
		# start the loop of the wordlist
		while ($plain = <$wordlistsfh>) {
			# print the current plain to the wordlist
			print $wordlistfh $plain;
			#print $plain;
			
			$s++;
			#sleep if $s == 8;
			#exit if $s == 7;
			#print $plain if $s == 10;
			#print $rule if $s == 10;
			my $asdf = "/home/coolbry95/wrappertesting/rules/" . $a;
			# get what hashcat thinks the mangled is
			my $out = `$hashcat --stdout $wordsingle -r $asdf`;

			#print "$out";
			#print $plain;
			#print $rule;
			#print length($plain) . "\n";
			# print the next part of the variable
			chomp $out;
			chomp $plain;
			print "\t{" . '"'. $plain .'"'. ", " .'"'. $out .'"'. "}," . "\n";

			#print $plain;
			#my $the = truncate($wordlistfh, length($plain));
			my $the = truncate($wordlistfh, 0);
			print "fucking asdfasdf" if ! defined $the;
			#print "fucking shit fuck" if !$the;
			seek($wordlistfh, 0, 0);
		}
		seek($wordlistsfh, 0, 0);
		#truncate($singlefh, length($rule));
	}
	# close the rule file we are on now
	close($rulefh);

	#$temp++;
	#last if $temp == 1;
	# print the last part of the variable
	print '}' . "\n";
}
close($wordlistfh);
close($wordlistsfh);
close($singlefh);
