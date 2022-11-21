#!/bin/sh

green="$(tput setaf 2)"
red="$(tput setaf 1)"
normal="$(tput sgr0)"

run_single()
{
	local inputpepo outdir testdir testout failed

	inputpepo="$1"
	outdir="$2"
	testdir="$3"

	testout="$outdir/$inputpepo.out"

	failed="0"

	if ! [ -f "$testdir"/input ]; then
		echo "No input file present, will create an empty one" >&2
		touch "$testdir/input"
		failed="1"
	fi

	"$qdpep8" -t \
		-o "$outdir/output"  \
		-i "$testdir/input" \
		"$inputpepo" >"$outdir/trace"

	if ! [ -f "$testdir/expected_output" ]; then
		echo "missing expected_output, will create with output from current test" >&2
		cp "$outdir/output" "$testdir/expected_output"
		failed="1"
	fi

	if ! [ -f "$testdir/expected_trace" ]; then
		echo "missing expected_trace, will create trace from current test" >&2
		cp "$outdir/trace" "$testdir/expected_trace"
		failed="1"
	fi

	diff -au "$testdir/expected_output" "$outdir/output" >"$outdir/diffout"

	if [ "$?" -ne 0 ]; then
		echo "output mismatch" >&2
		cat "$outdir/diffout" >&2
		failed="1"
	fi

	diff -au "$testdir/expected_trace" "$outdir/trace" >"$outdir/difftrace"

	if [ "$?" -ne 0 ]; then
		echo "trace mismatch" >&2
		cat "$outdir/difftrace" >&2
		failed="1"
	fi

	if [ "$failed" -eq 0 ]; then
		echo "${green}OK - $pepos $testdir ${normal}"
	else
		echo "${red}FAIL - $pepos $testdir ${normal}"
	fi
}

run_multi()
(
	local pepo_path testdir

	pepo_path="$(realpath "$1")"
	testdir="$2"

	for dir in subtests/*; do
		run_single "$pepo_path" "$testdir" "$dir"
	done
)

run_tests()
(
	local testdir pepos

	cd "$1"

	testdir="$(mktemp -d)"

	pepos="$(ls *pepo)"
	if [ "$(echo "$pepos" | wc -l)" -ne 1 ]; then
		echo "Malconfigured test, aborting"
		return 1
	fi

	if [ -d subtests ]; then
		run_multi "$pepos" "$testdir"
	else
		run_single "$pepos" "$testdir" "."
	fi

	rm -rf "$testdir"
)

(
	cd ..
	make
)

qdpep8="$(realpath "../bin/qdpep8_cli")"

tests="$(ls tests)"

if [ "$#" -eq 1 ]; then
	filter_tests="$1"
	tests="$(echo "$tests" | grep "$1")"
fi

for i in $tests; do
	run_tests tests/"$i"
done
