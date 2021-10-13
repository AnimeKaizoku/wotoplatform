#
# This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
# Copyright (c) 2021 AmanoTeam.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, version 3.
#
# This program is distributed in the hope that it will be useful, but
# WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.
#


buildApp() 
{
	# clear the screen (the terminal)
	clear

	echo -e "building it, please wait a bit..."

	go build -o wp-server
}

runApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are done building it,\n->now running the server...\n-------------------"

	./wp-server
}

testApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are running all test files (*_test.go);\nplease wait a bit"

	go test -v ./...
}

if [ "$1" == "test" ];
then
	testApp;
	exit 0
fi;

operations=0

if [ -z "$1" ] || [ "$1" == "true" ] || [ "$1" == "1" ];
then
	buildApp;
	operations=$((i+1))
fi;

if [ -z "$2" ] || [ "$2" == "true" ] || [ "$2" == "1" ];
then
	runApp;
	operations=$((i+1))
fi;

if [ $operations == 0 ]
then
	echo "You have done nothing!"
fi;